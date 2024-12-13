package main

import (
	"errors"
	"fmt"
)

type command struct {
	name string
	args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("username is missing")
	}

	username := cmd.args[0]
	err := s.cfg.SetUser(username)
	if err != nil {
		return err
	}

	fmt.Printf("User %s has been set\n", username)

	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	if _, ok := c.handlers[name]; !ok {
		c.handlers[name] = f
	}
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.handlers[cmd.name]; !ok {
		return errors.New("command is not registered")
	}

	return c.handlers[cmd.name](s, cmd)
}
