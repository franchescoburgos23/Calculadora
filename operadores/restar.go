package operadores

import (
	"fmt"
	"strconv"
	"strings"
)

func Restar(operacion string) int {

	valores := strings.Split(operacion, "-")
	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, error := strconv.Atoi(valores[i])

		if error != nil {
			fmt.Println(error)
			fmt.Println("Error tiene que ingresar un numero entero")
			fmt.Println("o Solo debes realizar con un operador")
		} else {
			if resultado == 0 {
				resultado = num
			} else {
				resultado -= num
			}

		}

	}

	return resultado
}
