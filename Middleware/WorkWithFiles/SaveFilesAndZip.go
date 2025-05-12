package WorkWithFiles

import (
	"MPT-Schedule/Models"
	"archive/zip"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ZippingFiles(files []*excelize.File, rData *Models.RequestData, GroupFolder string) string {
	// Создаем временную директорию
	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Создаем структуру папок в временной директории
	scDir := filepath.Join(tempDir, "schedules")
	err = os.Mkdir(scDir, 0755)
	if err != nil {
		log.Fatal(err)
	}

	// Сохраняем Excel файлы
	for i := 0; i < len(files); i++ {
		filePath := filepath.Join(scDir, files[i].Path+".xlsx")
		if err := files[i].SaveAs(filePath); err != nil {
			log.Fatal(err)
		}
	}

	// Создаем файл с ошибками
	filePath := filepath.Join(tempDir, "errors.txt")
	file, _ := os.Create(filePath)
	defer file.Close()
	file.WriteString(strings.Join(rData.FilesErrors, "\n"))

	// Копируем содержимое GroupFolder в временную директорию
	if GroupFolder != "" {
		err = copyFolderContents(GroupFolder, filepath.Join(tempDir, filepath.Base(GroupFolder)))
		if err != nil {
			log.Printf("Warning: couldn't copy group folder: %v", err)
		}
	}

	return createZip(tempDir)
}

func copyFolderContents(src, dst string) error {
	return filepath.Walk(src, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}

		dstPath := filepath.Join(dst, relPath)

		if info.IsDir() {
			return os.MkdirAll(dstPath, info.Mode())
		}

		// Копируем файл
		srcFile, err := os.Open(path)
		if err != nil {
			return err
		}
		defer srcFile.Close()

		dstFile, err := os.Create(dstPath)
		if err != nil {
			return err
		}
		defer dstFile.Close()

		_, err = io.Copy(dstFile, srcFile)
		return err
	})
}

func createZip(sourceDir string) string {
	zipFileName := fmt.Sprintf("schedules_%d.zip", time.Now().Unix())
	zipFilePath := filepath.Join(os.TempDir(), zipFileName)

	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		panic(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil // пропускаем директории
		}

		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath
		header.Method = zip.Deflate // Используем сжатие

		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = io.Copy(writer, file)
		return err
	})
	if err != nil {
		panic(err)
	}
	return zipFilePath
}
