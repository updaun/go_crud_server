package cmd

import (
	"fmt"

	"example.com/m/config"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}
	fmt.Println(c.config.Server.Port)
	return c
}
