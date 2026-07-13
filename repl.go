package main

import (
	"fmt"

	"github.com/Ripple9697/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	handler,ok := c.handlers[cmd.name]
	if !ok {
	return fmt.Errorf("command not found")
	}
	return handler(s,cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}



func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login expects a single argument, the username.")
	}
	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("sucessfully set username")
	return nil

}
