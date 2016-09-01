package main

import (
	"os"

	"github.com/gianarb/argh/cmd"
	"github.com/mitchellh/cli"
)

func main() {
	c := cli.NewCLI("argh", "0.0.1")
	c.Args = os.Args[1:]

	c.Commands = map[string]cli.CommandFactory{
		"generate": func() (cli.Command, error) {
			return &cmd.BuildCmd{}, nil
		},
	}

	exitStatus, _ := c.Run()

	os.Exit(exitStatus)
}
