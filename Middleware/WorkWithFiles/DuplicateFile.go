package WorkWithFiles

import (
	"io"
	"os"
)

func DuplicateFile() error {
	src, err := os.Open("Template/SCTemplate.xlsx")
	if err != nil {
		return err
	}
	defer src.Close()

	// Создаем новый файл для записи (если он существует, он будет перезаписан)
	dst, err := os.Create("Schedule1.xlsx")
	if err != nil {
		return err
	}
	defer dst.Close()
	dst2, err := os.Create("Schedule2.xlsx")
	if err != nil {
		return err
	}
	defer dst2.Close()
	// Копируем содержимое из исходного файла в новый файл
	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	_, err = io.Copy(dst2, src)
	if err != nil {
		return err
	}
	// Явное закрытие файла назначения для сброса буфера и освобождения ресурсов
	err = dst.Close()
	if err != nil {
		return err
	}
	err = dst2.Close()
	if err != nil {
		return err
	}
	return nil
}
