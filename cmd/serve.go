package cmd

import (
	"flag"
	"net/http"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/gianarb/argh/front"
	"github.com/gorilla/mux"
)

type ServeCmd struct {
}

func (c *ServeCmd) Run(args []string) int {
	logrus.Info("Argh! Come on! It's time to weight the anchor")

	var port string
	cmdFlags := flag.NewFlagSet("event", flag.ContinueOnError)
	cmdFlags.StringVar(&port, "port", ":8000", "port")
	if err := cmdFlags.Parse(args); err != nil {
		logrus.WithField("error", err).Warn("Problem to parse arguments")
	}

	logrus.Infof("HTTP Server runs on port %s", port)
	r := mux.NewRouter()
	r.HandleFunc("/feed.xml", front.FeedHandler()).Methods("GET")
	http.ListenAndServe(port, r)
	return 0
}

func (c *ServeCmd) Help() string {
	helpText := `
Usage: start HTTP server.
Options:
	-port=:8000			Servert port
`
	return strings.TrimSpace(helpText)
}

func (r *ServeCmd) Synopsis() string {
	return "Start argh"
}
