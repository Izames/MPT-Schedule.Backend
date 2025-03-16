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

func ZippingFiles(files []*excelize.File) string {
	//сжатие в зип
	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	scDir := filepath.Join(tempDir, "schedules")
	err = os.Mkdir(scDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(files); i++ {
		filePath := filepath.Join(scDir, files[i].Path+".xlsx")
		if err := files[i].SaveAs(filePath); err != nil {
			log.Fatal(err)
		}
	}
	filePath := filepath.Join(tempDir, "errors.txt") //Construct the full path
	file, _ := os.Create(filePath)
	defer file.Close()
	defer os.Remove(filePath)
	file.WriteString(strings.Join(Models.FilesErrors, "\n"))
	return createZip(tempDir)

}
func createZip(sourceDir string) string {
	zipFileName := fmt.Sprintf("schedules.zip", time.Now().Unix())
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
		header.Name = relPath // используем относительный путь

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
