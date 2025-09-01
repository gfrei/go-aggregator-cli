package main

import (
	"fmt"
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
		return fmt.Errorf("unknown command %q", cmd.name)
	}

	return commandFunc(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, ok := c.commandMap[name]; ok {
		return fmt.Errorf("command %q already registered", name)
	}
	c.commandMap[name] = f

	return nil
}

func initCommands() commands {
	return commands{
		commandMap: make(map[string]func(*state, command) error),
	}
}

func processCommand(s *state, cmds *commands, args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("no command provided")
	}

	cmd := command{
		name: args[0],
		args: args[1:],
	}

	return cmds.run(s, cmd)
}
