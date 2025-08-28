package main

import (
	"fmt"

	"github.com/gfrei/gator-cli/internal/config"
)

func main() {
	config.SetUser("gfrei")

	cfg, err := config.Read()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(cfg)
	}
}
