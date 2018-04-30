package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}

// Just for fun: build main.go and run it with main.go
// > ./main main.go
