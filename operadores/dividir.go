package operadores

import (
	"errors"
	"strconv"
	"strings"
)

func Dividir(operacion string) (int, error) {

	valores := strings.Split(operacion, "/")
	resultado := 0

	for i := 0; i < len(valores); i++ {
		num, err := strconv.Atoi(valores[i])

		if err != nil {
			///fmt.Println(error)
			//fmt.Println("Error tiene que ingresar un numero entero")
			//fmt.Println("o Solo debes realizar con un operador")
			return 0, err
		}
		if num == 0 {
			return 0, errors.New("0 no es divisible")
		} else {
			if resultado == 0 {
				resultado = num
			} else {
				resultado /= num
			}

		}

	}

	return resultado, nil
}
