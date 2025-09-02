package main

import (
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	_state, err := initState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_commands := initCommands()

	_commands.register("login", handlerLogin)
	_commands.register("register", handlerRegister)
	_commands.register("reset", handlerReset)
	_commands.register("users", handlerUsers)
	_commands.register("agg", handlerAgg)

	err = processCommand(&_state, &_commands, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
