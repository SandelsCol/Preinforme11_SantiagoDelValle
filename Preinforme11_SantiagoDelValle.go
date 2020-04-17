package Go

import (
	"fmt"
)

func contraseña() {
	var x string
	fmt.Scanf("%s", &x)
	var arreglo [4][7]int

	arreglo[0][5] = 3
	arreglo[2][6] = 2
	arreglo[3][6] = 1

	final_arreglo := arreglo[0][5] * arreglo[2][6] * arreglo[3][6]

	fmt.Println(final_arreglo)
}
