package main

import (
	"fmt"
)

func main() {
	var a int

	fmt.Print("Введите число: ")
	fmt.Scanln(&a)
  
  fmt.Println("единицы числа:" + (a % 10))
  fmt.Println("десятки числа: " + ((a / 10) % 10))
  fmt.Println("сотни числа: " + ((a / 100) % 10))
}
  
