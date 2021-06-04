package main

import (
	"fmt"
	"os"

	"github.com/haunt98/strtor/internal/normalize"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(`go run main.go "Your string goes here"`)
		return
	}

	fmt.Println(normalize.Normalize(args[1]))
}
