package cmd

import (
	"example.com/m/config"
	"example.com/m/network"
	"example.com/m/repository"
	"example.com/m/service"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	service    *service.Service
	repository *repository.Repository
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}
	c.repository = repository.NewRepository()
	c.service = service.NewService(c.repository)
	c.network = network.NewNetwork(c.service)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
