package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

type Cliente struct {
	Legajo          int
	NombreYApellido string
	DNI             int
	Telefono        string
	Domicilio       string
}

const filename = "./customers.txt"

func main() {
	cliente := Cliente{}
	cliente = generarLegajoCliente(cliente)

	defer func() {
		errorArchivo := verificarSiClienteExiste(cliente)
		if errorArchivo != nil {
			fmt.Println(errorArchivo.Error())
		}
	}()

	defer func() {
		_, errorCliente := validarDatosCliente(cliente)
		if errorCliente != nil {
			fmt.Println(errorCliente.Error())
		}
	}()

	fmt.Println("Fin de la ejecucion")
}

func generarLegajoCliente(cliente Cliente) Cliente {
	cliente.Legajo = rand.Int()
	if cliente.Legajo == 0 {
		panic("el legajo del cliente es invalido")
	}
	return cliente
}

func verificarSiClienteExiste(c Cliente) error {
	defer recuperarErrorArchivo()
	_, err := os.ReadFile(filename)
	if err != nil {
		panic("el archivo indicado no fue encontrado o esta dañado")
	}
	return err
}

func validarDatosCliente(c Cliente) (bool, error) {

	clienteOK := false

	if c.Legajo == 0 {
		defer mostrarErroresEjecucion()
		return clienteOK, errors.New("el legajo del cliente no existe")
	}
	if c.DNI == 0 {
		defer mostrarErroresEjecucion()
		return clienteOK, errors.New("el DNI del cliente no existe")
	}
	if c.Domicilio == "" {
		defer mostrarErroresEjecucion()
		return clienteOK, errors.New("el domicilio del cliente no existe")
	}
	if c.NombreYApellido == "" {
		defer mostrarErroresEjecucion()
		return clienteOK, errors.New("el nombre y apellido del cliente no existe")
	}
	if c.Telefono == "" {
		defer mostrarErroresEjecucion()
		return clienteOK, errors.New("el telefono del cliente no existe")
	}

	clienteOK = true
	return clienteOK, nil
}

func recuperarErrorArchivo() {
	err := recover()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("No han quedado archivos abiertos")
}

func mostrarErroresEjecucion() {
	fmt.Println("Se detectaron varios errores de ejecucion")
}
