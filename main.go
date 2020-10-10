package main

import "fmt"

func main() {
	var n int
	var entrada int
	var suma int
	//recibe entrero para definir tamaños del slice
	fmt.Scan(&n)
	s := make([]int, 0, n)
	//recibe los enteros y los almacena en el slice
	for i := 0; i < n; i++ {
		fmt.Scan(&entrada)
		s = append(s, entrada)
	}
	//realiza la suma
	for _, v := range s {
		suma += v
	}
	//muestra la suma
	fmt.Println(suma)
}
