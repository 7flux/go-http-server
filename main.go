package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    f, err := os.Open("messages.txt")

    if err != nil {
        log.Fatal(err)
    }

    defer f.Close()

    getLineChannel(f)
}

func getLineChannel(f io.ReadCloser) <-chan string {
    out := make(chan string, 1)

    go func() {
        defer close(out)
        defer f.Close()

        str := ""
        for {
            data := make([]byte, 8)
            n, err := f.Read(data)
            if err != nil {
                break
            }

            data = data[:n] // to save space?

        }
    }()

    fmt.Printf("read: %s\n", out)

    return out
}
