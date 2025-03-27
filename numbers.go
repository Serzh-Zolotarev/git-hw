package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	n := ""
	fmt.Print("Введите целое число: ")
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := strconv.Atoi(n); err == nil {
		fmt.Printf("Вы ввели число: %s\n", n)
	} else {
		fmt.Printf("Вы ввели: %s\n", n)
	}

	fmt.Printf("Вы ввели число: %d\n", n)
}
