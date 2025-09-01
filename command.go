package main

import (
	"fmt"
	"os"
)

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

func initCommands() commands {
	return commands{
		commandMap: make(map[string]func(*state, command) error, 0),
	}
}

func processCommand(_state state, _commands commands) error {
	if len(os.Args) < 2 {
		return fmt.Errorf("no commands")
	}

	cmdName := os.Args[1]
	args := os.Args[2:]

	cmd := command{
		name: cmdName,
		args: args,
	}

	err := _commands.run(&_state, cmd)
	if err != nil {
		return err
	}

	return nil
}
