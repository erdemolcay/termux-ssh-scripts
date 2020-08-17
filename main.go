package main

import (
	"fmt"
	"os"
	"termux-ssh-scripts/app"
	"termux-ssh-scripts/config"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(
			"usage: termux-ssh-scripts <command> [<args>]\n" +
				"commands: install, update\n" +
				"args: api-token, zone-id")
		os.Exit(1)
	}

	sc := os.Args[1]

	if sc != "install" && sc != "update" {
		fmt.Printf("Invalid command: %q\n", sc)
		os.Exit(1)
	}

	c := config.New(sc)

	a := app.New(c)

	switch sc {
	case "install":
		a.Install()
		os.Exit(0)
	case "update":
		a.Run()
		os.Exit(0)
	}
}
