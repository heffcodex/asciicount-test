package main

import (
	"fmt"
	"os"

	"asciicount-test/cmd"
	"asciicount-test/consts"
)

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println(consts.Usage)
		return
	}

	var command cmd.Command

	switch args[1] {
	case consts.CmdGenerate:
		command = new(cmd.CommandGenerate)
	case consts.CmdCount:
		command = new(cmd.CommandCount)
	default:
		fmt.Println(consts.Usage)
		return
	}

	if err := command.Execute(); err != nil {
		fmt.Printf("An error occurred: %s\n", err.Error())
	}
}
