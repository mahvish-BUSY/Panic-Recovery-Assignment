package main

import (
	"fmt"
)

func divide(a, b int) int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered from Panic, Error : ", err)
		}
	}()
	fmt.Println("Inside divide func")
	res := a/b
	fmt.Println("divide func terminated successfully")
	return res
}
func main() {
	var a, b int
	fmt.Println("Inside main func")
	fmt.Println("Enter two numbers for division : ")
	fmt.Scanf("%d %d", &a, &b)
	res := divide(a, b)
	fmt.Println("Result : ",res)
	fmt.Println("main func terminated successfully")
}
