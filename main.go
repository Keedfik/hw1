package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Укажите полный путь до файла вторым аргументом")
	}

	filePth := os.Args[1]

	var fileName, fileExt string
	// Напишите код, который выведет следующее
	// filename: <name>
	// extension: <extension>

	// Подсказка. Возможно вам понадобится функция strings.LastIndex
	// Для проверки своего решения используйте функции filepath.Base() filepath.Ext(
	// ) Они могут помочь для проверки решения

	PthArray := strings.Split(filePth, "/")
	NameArray := strings.Split(PthArray[len(PthArray)-1], ".")
	if len(NameArray) > 1 {
		fileExt = NameArray[len(NameArray)-1]
	}
	fileName = strings.Split(PthArray[len(PthArray)-1], fmt.Sprintf(`.%s`, fileExt))[0]

	fmt.Printf("filename: %s\n", fileName)
	fmt.Printf("extension: %s\n", fileExt)
}
