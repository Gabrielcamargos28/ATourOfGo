package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

/*
A função sqrt retorn
a a raiz quadrada de x .
Por padrão, se x for um valor negativo,
sqrt retornará um NaN indefinido.
*/
func main(){
	fmt.Println(sqrt(2), sqrt(-4))
}
