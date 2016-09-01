package cmd

import (
	"flag"
	"strings"

	"github.com/Sirupsen/logrus"
)

type BuildCmd struct {
}

func (c *BuildCmd) Run(args []string) int {
	logrus.Info("Argh! Come on! It's time to weight the anchor")

	var path string
	cmdFlags := flag.NewFlagSet("event", flag.ContinueOnError)
	cmdFlags.StringVar(&path, "list", "./list.txt", "list")
	if err := cmdFlags.Parse(args); err != nil {
		logrus.WithField("error", err).Warn("Problem to parse arguments")
	}

	return 0
}

func (c *BuildCmd) Help() string {
	helpText := `
Generate feed from a lists
Options:
	-list=./list.txt			List's path
`
	return strings.TrimSpace(helpText)
}

func (r *BuildCmd) Synopsis() string {
	return "Start argh"
}
