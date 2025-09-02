package main

import (
	"database/sql"

	"github.com/gfrei/gator-cli/internal/config"
	"github.com/gfrei/gator-cli/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}

func initState() (state, error) {
	config, err := config.Read()
	if err != nil {
		return state{}, err
	}

	db, err := sql.Open("postgres", config.DbUrl)
	if err != nil {
		return state{}, err
	}

	dbQueries := database.New(db)

	_state := state{
		db:     dbQueries,
		config: &config,
	}

	return _state, nil

}
