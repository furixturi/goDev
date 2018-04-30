package main

import (
	"fmt"
	"os"
)

func main() {
	//[/var/folders/6p/_z0jx7v974v_wsql5n2t79jw0000gn/T/go-build553020332/b001/exe/main wow wow wow]
	// first is the temporary address of the compiled main.go
	// after that are the words the user typed
	fmt.Println(os.Args)
}
