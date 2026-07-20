package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/Ripple9697/gator/internal/config"
	"github.com/Ripple9697/gator/internal/database"
	"github.com/google/uuid"
)

type state struct {
	db  *database.Queries
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
	handler, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	return handler(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login expects a single argument, the username.")
	}
	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("sucessfully switched user")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register expects a single argument, the username.")
	}

	params := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      cmd.args[0],
	}

	respUser, err := s.db.CreateUser(context.Background(), params)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}
	fmt.Println("sucessfully added user")
	fmt.Println(respUser)
	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.Delete(context.Background())
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println("sucessfully Deleted")
	return nil
}

func handlerUsers(s *state, cmd command) error {
	resp, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, v := range resp {
		if v.Name == s.cfg.CurrentUserName {
			fmt.Printf("* %s (current)\n", v.Name)
		} else {
			fmt.Printf("* %s\n", v.Name)
		}
	}

	return nil
}

func handlerAgg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("couldn't fetch feed: %w", err)
	}
	fmt.Printf("%+v\n", feed)
	return err
}
