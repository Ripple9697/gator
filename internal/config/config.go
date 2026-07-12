// Package config handles application configuration. stop harrasing me linter?
package config

import (
	"fmt"
	"log"
	"os"
)

type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func Read() Config {
	configPath := "~/.gatorconfig.json"
	content, err := os.ReadFile(configPath)
	if err != nil {
		log.Fatal(err
	}
	fmt.Println(string(content))
}
