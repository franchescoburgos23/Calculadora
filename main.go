package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/franchescoburgos23/Calculadora/operadores"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("La calculadora")
	scanner.Scan()
	operacion := scanner.Text()

	if strings.Contains(operacion, "+") {
		resultado := operadores.Sumar(operacion)
		fmt.Println(resultado)
	} else if strings.Contains(operacion, "-") {
		resultado := operadores.Restar(operacion)
		fmt.Println(resultado)
	} else if strings.Contains(operacion, "*") {
		resultado := operadores.Multiplicar(operacion)
		fmt.Println(resultado)

	} else if strings.Contains(operacion, "/") {
		resultado, err := operadores.Dividir(operacion)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(resultado)

	} else {
		fmt.Println("El operador esta mal ingresado")
	}

}
