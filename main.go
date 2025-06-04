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

	defer f.Close()

	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
	
		if err != nil {
			break
		}

		fmt.Printf("read: %s\n", string(data[:n]))
	}
}