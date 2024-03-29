package main

import (
	"fmt"
	"os"

	"github.com/haunt98/strtor/normalize"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println(`go run ./cmd/strtor/main.go "Your string goes here"`)
		return
	}

	fmt.Println(normalize.Normalize(args[1]))
}
