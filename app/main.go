package app

import (
	"github.com/hcsaustrup/gopt/helpers"
)

type Configuration struct {
	Path       string
	Repository *helpers.Repository
}

func (c *Configuration) GetRepository() (*helpers.Repository, error) {
	if c.Repository != nil {
		return c.Repository, nil
	}
	repository, err := helpers.NewRepository(c.Path)
	if err == nil {
		c.Repository = repository
	}
	return repository, err
}

var Config = &Configuration{}
