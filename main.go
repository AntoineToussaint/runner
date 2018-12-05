package main

import (
	"os"

	"github.com/hygge.io/hygge.io-core/runner"
)

func main() {
	switch os.Args[1] {
	case "run":
		runner.Run()
	case "child":
		runner.Child()
	default:
		panic("help")
	}
}
