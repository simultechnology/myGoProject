package main

import (
	f "fmt"
	"github.com/simultechnology/gosample"
	. "strings"
	"os"
	"log"
	"bufio"
)

func main() {
	f.Println(ToUpper(gosample.Message)) // hello world

	localMessage := "moon"
	f.Printf("Hi, %s\n", localMessage)

	n := 0
	for {
		n++
		if n > 20 {
			break
		}
		if n%2 == 0 {
			continue
		}
		f.Printf("%d,", n)
	}
	f.Println()
	// fileから内容を読み込む
	file, err := os.Open("src/data/sample.txt")

	if err != nil {
		log.Fatal(err)
	} else {
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			f.Println(scanner.Text())
		}
	}
}
