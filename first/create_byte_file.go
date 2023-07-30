package main

import (
	"fmt"
	"os"
)

func main() {
	// Открываем файл для записи
	file, err := os.Create("output.bin")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Записываем строку в файл как бинарные данные
	str := "Hello Bold!"
	bytesWritten, err := file.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Записано %d байт в файл.", bytesWritten)
}
