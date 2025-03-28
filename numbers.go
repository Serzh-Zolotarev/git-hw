package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	n := ""
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := strconv.Atoi(n); err == nil {
		fmt.Printf("Вы ввели число: %s\n", n)
	} else {
		fmt.Printf("Вы ввели текст: %s\n", n)
	}
}
