package main

import "fmt"

func main(){
	sum := 0
	for i:= 0; i < 10; i++{
		//0+1+2+3+4+5+6+7+8+9
		sum += i
	}

	fmt.Println(sum)
}