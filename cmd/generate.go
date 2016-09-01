package cmd

import (
	"bufio"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/SlyMarbo/rss"
	"github.com/gorilla/feeds"
)

type BuildCmd struct {
}

func (c *BuildCmd) Run(args []string) int {

	now := time.Now()

	feed := &feeds.Feed{
		Title:       "Docker Captain's feed",
		Link:        &feeds.Link{Href: "http://captains.today"},
		Description: "Updates from the docker captains!",
		Created:     now,
	}

	feed.Items = []*feeds.Item{}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "\n" {
			break
		}
		captainFeed, err := rss.Fetch(scanner.Text())
		if err != nil {
			logrus.WithField("error", err).Warnf("%s impossible to read. I jump it please verify", scanner.Text())
			continue
		}

		for _, item := range captainFeed.Items {
			feed.Items = append(feed.Items, &feeds.Item{
				Title:       item.Title,
				Link:        &feeds.Link{Href: item.Link},
				Description: item.Summary,
				Author:      &feeds.Author{Name: captainFeed.Nickname},
				Created:     item.Date,
			})
		}

	}

	f, err := os.Create("/tmp/captains.xml")

	if err != nil {
		log.Fatal(err)
	}

	w := bufio.NewWriter(f)
	err = feed.WriteAtom(w)

	if err != nil {
		log.Fatal(err)
	}
	return 0
}

func (c *BuildCmd) Help() string {
	helpText := `
Generate feed from a lists
Options:
	-list=./feeds.txt			Text file containing a list of feeds
`
	return strings.TrimSpace(helpText)
}

func (r *BuildCmd) Synopsis() string {
	return "Start argh"
}
