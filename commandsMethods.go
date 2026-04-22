package main

import "log"

func (c *commands) run(s *state, cmd command) error {
	function, ok := c.registeredCommands[cmd.Name]

	if !ok {
		log.Fatalf("Command not found")
	}

	return function(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}
