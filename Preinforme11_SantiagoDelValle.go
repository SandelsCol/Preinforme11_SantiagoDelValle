package Go
import (
    "bufio"
    "fmt"
    "os"
)
func contraseña() {
	var n int
	fmt.Scanf("%s",&n)
	var z int
	fmt.Scanf("%s",&z)
	var x int
	fmt.Scanf("%s",&x)
	var k int
	fmt.Scanf("%s",&k)
	var p int
	fmt.Scanf("%s",&p)
	var o int
	fmt.Scanf("%s",&o)

	var arreglo [4][7] int
	Parte1:=arreglo[0][5] := 8
	parte2:=arreglo[2][6] := 36521
	parte3:=arreglo[3][6] := 32389

	final_arreglo := Parte1 * parte2 * parte3

	if n==final_arreglo and z==final_arreglo and x==final_arreglo and k==final_arreglo and p==final_arreglo and o==final_arreglo{
		fmt.Println("Su contraseña es",final_arreglo)
	}else{
		fmt.Println("Haga bien los pasos por favor")
	}
}
