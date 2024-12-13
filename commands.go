package main

import "errors"

type command struct {
	Name string
	Args []string
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if _, ok := c.handlers[cmd.Name]; !ok {
		return errors.New("command is not registered")
	}

	return c.handlers[cmd.Name](s, cmd)
}
