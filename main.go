package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Taanviir/blog-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

	s := &state{
		cfg: &cfg,
	}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	if len(os.Args) < 2 {
		log.Fatalf("Usage: ./blog-aggregator <command> <args>")
	}

	cmd := command{
		name: os.Args[1],
		args: os.Args[2:],
	}

	err = cmds.run(s, cmd)
    if err != nil {
        log.Fatal(err)
    }
}
