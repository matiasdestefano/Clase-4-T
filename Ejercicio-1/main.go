package main

import (
	"fmt"
	"os"
)

func main() {
	leerArchivo("./customers.txt")
	fmt.Println("Ejecucion finalizada.")
}

func leerArchivo(filename string) []byte {
	defer func() {
		err := recover()

		if err != nil {
			fmt.Println(err)
		}
	}()

	fileData, err := os.ReadFile(filename)
	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
	return fileData
}
