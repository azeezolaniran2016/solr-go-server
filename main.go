package main

import (
	"github.com/azeezolaniran2016/solr-server/api"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.NewEntry(
		logrus.New(),
	)

	cfg, err := api.ParseConfig("")
	if err != nil {
		logger.WithError(err).Fatal("could not parse config")
	}

	deps := &api.Dependencies{
		Log: logger,
	}

	apiServer := api.New(cfg, deps)

	// Start server
	logger.Fatal(apiServer.Start())
}
