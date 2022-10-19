package main

import (
	"fmt"
	"os"

	"github.com/seailly/mi/repl"
)

func main() {
	fmt.Printf("Mi v0.0.1 ")
	repl.Start(os.Stdin, os.Stdout)
}
