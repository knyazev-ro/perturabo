package main

import (
	"os"
	"perturabo/commands"
)

func main() {
	commands.Handle(os.Args)
}
