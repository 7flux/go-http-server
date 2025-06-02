package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(f)
}