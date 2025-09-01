package main

import (
	"fmt"
	"os"
)

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
	_state, err := initState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_commands := initCommands()

	_commands.register("login", handlerLogin)

	err = processCommand(_state, _commands)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
