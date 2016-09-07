package cmd

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/SlyMarbo/rss"
	"github.com/gorilla/feeds"
)

type BuildCmd struct {
}

type ByPublishdate []*feeds.Item

func (a ByPublishdate) Len() int {
	return len(a)
}
func (a ByPublishdate) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (a ByPublishdate) Less(i, j int) bool {
	delta := a[i].Created.Sub(a[j].Created)
	return delta.Seconds() > 0
}

func (c *BuildCmd) Run(args []string) int {

	now := time.Now()

	feed := &feeds.Feed{
		Title:       "Docker Captain's feed",
		Description: "Updates from the docker captains!",
		Link:        &feeds.Link{Href: "http://argh.gianarb.it"},
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

	sort.Sort(ByPublishdate(feed.Items))

	path, err := filepath.Abs(args[0])

	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(path)

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
	cat feeds.txt | ./argh generate index.xml
`
	return strings.TrimSpace(helpText)
}

func (r *BuildCmd) Synopsis() string {
	return "Start argh"
}
