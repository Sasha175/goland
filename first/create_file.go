package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello Gold!"
	fmt.Println("Create file")
	file, err := os.Create("C:\\Users\\sasha\\YandexDisk\\Zenbook\\Программирование\\Golang\\first\\hello.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

	fmt.Println("Done.")
}
