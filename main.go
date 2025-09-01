package main

import (
	"fmt"
	"os"

	"github.com/gfrei/gator-cli/internal/config"
)

type state struct {
	config *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	commandFunc, ok := c.commandMap[cmd.name]

	if !ok {
		return fmt.Errorf("command not found %v", cmd.name)
	}

	err := commandFunc(s, cmd)
	if err != nil {
		return err
	}

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, ok := c.commandMap[name]; ok {
		return fmt.Errorf("command already registered")
	}
	c.commandMap[name] = f

	return nil
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
	config, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_state := state{
		config: &config,
	}

	_commands := commands{
		commandMap: make(map[string]func(*state, command) error, 0),
	}

	_commands.register("login", handlerLogin)

	if len(os.Args) < 2 {
		fmt.Println("No arguments, exiting...")
		os.Exit(1)
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	cmd := command{
		name: cmdName,
		args: args,
	}

	err = _commands.run(&_state, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
