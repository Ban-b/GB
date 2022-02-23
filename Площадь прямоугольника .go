package main

import (
	"fmt"
)

func main() {
	var a, b int

	fmt.Print("Введите первое число: ")
	fmt.Scanln(&a)

	fmt.Print("Введите второе число: ")
	fmt.Scanln(&b)
  
  var c = a * b
  fmt.Println(c)
}
