package cmd

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/SlyMarbo/rss"
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

	f, err := os.Open(path)
	scanner := bufio.NewScanner(f)
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		return 1
	}
	if err = scanner.Err(); err != nil {
		logrus.WithField("error", err).Warn("Impossible read this file.")
	}
	for scanner.Scan() {
		feed, err := rss.Fetch(scanner.Text())
		if err != nil {
			logrus.WithField("error", err).Warnf("%s impossible to read. I jump it please verify", scanner.Text())
			continue
		}
		fmt.Println(feed.Title)
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
