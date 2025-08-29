package main

import (
	"fmt"

	"github.com/gfrei/gator-cli/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login error: no username")
	}

	user := cmd.args[0]

	err := s.config.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Println("Login as:", user)

	return nil
}

func main() {
}
