package main

import (
	"fmt"
)

func main() {
  const pi = 3.14 int
  	var R, S, l, D int

	fmt.Print("Введите площадь окружности: ")
	fmt.Scanln(&S)
	
  R = 2 * S / pi
	
  D = math.Sqrt(R)
	fmt.Println("Диаметр окружности: " + D)	
  
  l = pi * D
	fmt.Println("Длина окружности: " + l)
}	
