package config

import (
	"encoding/json"
	"os"
)

func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	return c.write()
}

func (c *Config) write() error {
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(fullPath, data, 0600)
}
