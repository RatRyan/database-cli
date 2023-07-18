package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RatRyan/dbapp/internal/console"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	for console.Running {
		fmt.Print("\nDatabase> ")

		rawInput, _ := reader.ReadString('\n')
		input := strings.Fields(rawInput)
		command := append([]string{""}, input...)

		if err := console.App.Run(command); err != nil {
			continue
		}
	}
}
