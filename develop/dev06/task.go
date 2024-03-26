package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Недостаточно аргументов. Ожидается минимум 3 аргумента.")
		os.Exit(1)
	}

	arg1 := os.Args[1]
	arg2 := os.Args[2]
	arg3 := os.Args[3]

	fmt.Println("Аргумент 1:", arg1)
	fmt.Println("Аргумент 2:", arg2)
	fmt.Println("Аргумент 3:", arg3)
}

type flags struct {
	fields string
}
