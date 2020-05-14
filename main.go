package main

import (
	"strings"
	"time"

	"github.com/azeezolaniran2016/solr-server/api"
	"github.com/sendgrid/go-solr"
	"github.com/sirupsen/logrus"
)

func logrusLevel(logLevel string) logrus.Level {
	ll := strings.ToLower(logLevel)
	switch ll {
	case "info":
		{
			return logrus.InfoLevel
		}
	default:
		{
			return logrus.DebugLevel
		}
	}
}

func main() {
	logger := logrus.NewEntry(
		logrus.New(),
	)

	logger.Logger.SetFormatter(&logrus.JSONFormatter{})

	cfg, err := api.ParseConfig("")
	if err != nil {
		logger.WithError(err).Fatal("could not parse config")
	}

	// solrZk := solr.NewSolrZK(cfg.SolrZooKeepers, cfg.SolrZooKeeperRoot, cfg.SolrCollection)

	// // locator := solrZk.GetSolrLocator()
	// err = solrZk.Listen()
	// if err != nil {
	// 	logger.WithError(err).Fatal("failed to listen to zookeeper")
	// }
	// https, err := solrZk.UseHTTPS()
	// if err != nil {
	// 	panic(err)
	// }
	solrHTTP, err := solr.NewSolrHTTP(true, cfg.SolrCollection) //, solr.User("solr"), solr.Password("admin"), solr.MinRF(2), solr.QueryRouter(solr.NewRoundRobinRouter()))
	if err != nil {
		panic(err)
	}

	solrHTTPRetrier := solr.NewSolrHttpRetrier(solrHTTP, 5, 100*time.Millisecond)
	logger.Info("created solr client")
	logger.Logger.SetLevel(logrusLevel(cfg.LogLevel))
	logger.Debugf("config: %+v", cfg)

	deps := &api.Dependencies{
		Log:        logger,
		SolrClient: solrHTTPRetrier,
	}

	apiServer := api.New(cfg, deps)

	// Start server
	logger.Fatal(apiServer.Start())
}
