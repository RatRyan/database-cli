package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/RatRyan/dbapp/internal/console"
)

var reader = bufio.NewReader(os.Stdin)

func main() {
	go console.SetupMap()
	cls := exec.Command("cmd", "/c", "cls")
	cls.Stdout = os.Stdout
	cls.Run()

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