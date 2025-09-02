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

func setUser(s *state, username string) error {
	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user %q not registered", username)
	}

	err = s.config.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("logged in as: %q\n", username)

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

	fmt.Printf("registered user %q\n", user.Name)

	return setUser(s, user.Name)
}

func handlerReset(s *state, cmd command) error {
	count, err := s.db.CountUsers(context.Background())
	if err != nil {
		return err
	}

	err = s.db.DeleteUsers(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("deleted %v users\n", count)

	return nil
}

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user.Name != s.config.CurrentUserName {
			fmt.Printf("* %v\n", user.Name)
		} else {
			fmt.Printf("* %v (current)\n", user.Name)
		}
	}

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
	_commands.register("register", handlerRegister)
	_commands.register("reset", handlerReset)
	_commands.register("users", handlerUsers)

	err = processCommand(&_state, &_commands, os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
