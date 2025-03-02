package WorkWithFiles

import (
	"archive/zip"
	"fmt"
	"github.com/xuri/excelize/v2"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"time"
)

func Zipping(file *excelize.File) string {
	// Создаем временную директорию
	tempDir, err := os.MkdirTemp("", "temp")
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Создаем папки для семестров
	folder1 := filepath.Join(tempDir, "Первый семестр")
	folder2 := filepath.Join(tempDir, "Второй семестр")
	if err := os.MkdirAll(folder1, 0755); err != nil {
		log.Fatal(err)
	}
	if err := os.MkdirAll(folder2, 0755); err != nil {
		log.Fatal(err)
	}

	// Сохраняем файл в папку "Первый семестр"
	filePath1 := filepath.Join(folder1, "Расписание.xlsx")
	if err := file.SaveAs(filePath1); err != nil {
		log.Fatal(err)
	}

	// Сохраняем файл в папку "Второй семестр"
	filePath2 := filepath.Join(folder2, "Расписание.xlsx")
	if err := file.SaveAs(filePath2); err != nil {
		log.Fatal(err)
	}

	// Создаем один зип-архив, содержащий обе папки
	zipFilePath := createZip(tempDir)

	return zipFilePath
}

func createZip(sourceDir string) string {
	// Генерируем имя зип-файла с временной меткой
	zipFileName := fmt.Sprintf("schedules.zip", time.Now().Unix())
	zipFilePath := filepath.Join(os.TempDir(), zipFileName)

	// Создаем зип-файл
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	err = filepath.Walk(sourceDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Получаем относительный путь
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		// Создаем заголовок для файла в архиве
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = relPath // Используем относительный путь

		// Если это директория, добавляем её в архив
		if info.IsDir() {
			header.Name += "/" // Добавляем слэш для обозначения директории
			_, err = zipWriter.CreateHeader(header)
			return err
		}

		// Создаем файл в архиве
		writer, err := zipWriter.CreateHeader(header)
		if err != nil {
			return err
		}

		// Открываем исходный файл
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Копируем содержимое файла в архив
		_, err = io.Copy(writer, file)
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	return zipFilePath
}
