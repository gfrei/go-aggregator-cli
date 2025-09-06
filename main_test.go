package main

import (
	"testing"

	"github.com/gfrei/gator-cli/internal/config"
)

func TestHandlerLogin(t *testing.T) {
	config := config.New()

	state := state{
		config: &config,
	}

	cmd := command{
		name: "login",
		args: []string{"gfrei"},
	}

	err := handlerLogin(&state, cmd)

	if err != nil {
		t.Errorf("")
		return
	} else {
		t.Log("Ok!")
	}
}

//TODO: test inputs
