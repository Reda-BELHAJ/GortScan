package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func main() {
	argsWithProg := os.Args

	color.Cyan("[+]")
	fmt.Println(argsWithProg)

	// Initialising
}
