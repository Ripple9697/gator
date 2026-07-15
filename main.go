package main

import (
	"fmt"
	"os"

	"github.com/Ripple9697/gator/internal/config"
)

func main() {
	fmt.Println("started")
	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Read config again: %+v\n", cfg)

	stg := state{cfg: &cfg}
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmds := commands{handlers: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)

	cmd := command{
		name: arguments[1],
		args: arguments[2:],
	}

	err = cmds.run(&stg, cmd)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
