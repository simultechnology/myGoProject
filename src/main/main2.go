package main

import (
	"errors"
	"log"
	"fmt"
)

func div(i, j int) (int, error) {
	if j == 0 {
		return 0, errors.New("devided by zero")
	}
	return i/j, nil
}

func main() {
	n, err := div(10, 0)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(n)
}

