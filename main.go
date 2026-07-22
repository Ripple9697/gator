package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Ripple9697/gator/internal/config"
	"github.com/Ripple9697/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("started")

	cfg, err := config.Read()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	dbQueries := database.New(db)

	stg := state{db: dbQueries, cfg: &cfg}
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	cmds := commands{handlers: make(map[string]func(*state, command) error)}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)

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
