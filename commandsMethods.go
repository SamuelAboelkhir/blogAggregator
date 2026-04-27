package main

import "errors"

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.registeredCommands[cmd.Name]

	if !ok {
		return errors.New("Command not found")
	}

	return function(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
