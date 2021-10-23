package main

import (
	"fmt"
	"os"
)

func main() {
	parseArgs(os.Args[1:])
}

func parseArgs(args []string) {
	for i := range args {
		if args[i] == "-h" || args[i] == "--help" {
			showHelp()
		}
	}
}

func showHelp() {
	fmt.Println("Help message")
}
