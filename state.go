package main

import (
	"github.com/gfrei/gator-cli/internal/config"
)

type state struct {
	config *config.Config
}

func initState() (state, error) {
	config, err := config.Read()
	if err != nil {
		return state{}, err
	}

	_state := state{
		config: &config,
	}

	return _state, nil

}
