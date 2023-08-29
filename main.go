/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"os"
	"path/filepath"

	"github.com/hcsaustrup/gopt/app"
	"github.com/hcsaustrup/gopt/cmd"
	"github.com/sirupsen/logrus"
)

func main() {

	path := os.Getenv("GOPT_PACKAGES")
	if path == "" {
		path = filepath.Join(os.Getenv("HOME"), "opt", "packages")
		logrus.Warnf("GOPT_PACKAGES is empty - using %s", path)
	}

	app.Config.Path = path

	cmd.Execute()
}
