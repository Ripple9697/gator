package main

import (
	"fmt"
	"log"
	"os"
	"github.com/Ripple9697/gator/internal/config"
)

func main() {
	fmt.Println("started")
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config again: %+v\n", cfg)


	stg := state{cfg: &cfg}
	arguments :=  os.Args 
	if len(arguments) < 2 {
		log.Fatalf("Required command name")
	}

	cmds := commands{handlers: make(map[string]func(*state, command) error)}

	cmds.register("login",handlerLogin)

	cmd := command{arguments[1],arguments[2:]}

	err = cmds.run(&stg,cmd)
	if err != nil {
		log.Fatalf("cannot run? %+v\n",err)
	}
}
