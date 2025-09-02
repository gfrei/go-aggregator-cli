package main

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gfrei/gator-cli/internal/database"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("login error: no username")
	}

	user := cmd.args[0]

	return setUser(s, user)
}

func setUser(s *state, user string) error {
	err := s.config.SetUser(user)
	if err != nil {
		return err
	}

	fmt.Println("Login as:", user)

	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("register error: no username")
	}

	fullName := (strings.Join(cmd.args, " "))

	user, err := s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      fullName,
	})

	if err != nil {
		return err
	}

	fmt.Printf("register user %q\n", user.Name)

	return setUser(s, user.Name)
}

func main() {
	_state, err := initState()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	_commands := initCommands()

	_commands.register("login", handlerLogin)
	_commands.register("register", handlerRegister)

	err = processCommand(&_state, &_commands, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
