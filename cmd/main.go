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

func clearScreen() {
	cls := exec.Command("cmd", "/c", "cls")
	cls.Stdout = os.Stdout
	cls.Run()
}

func main() {
	clearScreen()
	for console.Running {
		fmt.Print("\n > ")
		rawInput, _ := reader.ReadString('\n')
		input := strings.Fields(rawInput)
		command := append([]string{""}, input...)
		console.App.Run(command)
	}
}
