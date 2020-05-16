package main

import (
	"strings"
	"time"

	"github.com/azeezolaniran2016/solr-server/api"
	"github.com/pkg/errors"
	"github.com/sendgrid/go-solr"
	"github.com/sirupsen/logrus"
)

// logrus level converts log levels (e.g info, debug) to logrus levels (which are integers)
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
	logger.Logger.SetLevel(logrusLevel(cfg.LogLevel))

	solrClient, err := createSolrClient(cfg)
	if err != nil {
		logger.WithError(err).Fatal()
	}
	logger.Info("created solr client")

	deps := &api.Dependencies{
		Log:        logger,
		SolrClient: solrClient,
	}

	apiServer := api.New(cfg, deps)

	// Start server
	logger.Fatal(apiServer.Start())
}

func createSolrClient(cfg *api.Config) (solr.SolrHTTP, error) {
	if cfg.SolrUseZK {
		solrZk := solr.NewSolrZK(cfg.SolrZooKeepers, cfg.SolrZooKeeperRoot, cfg.SolrCollection)

		// locator := solrZk.GetSolrLocator()
		err := solrZk.Listen()
		if err != nil {
			return nil, errors.Wrap(err, "failed to listen to zookeeper")
		}

		if cfg.SolrUseHTTPS {
			_, err := solrZk.UseHTTPS()
			if err != nil {
				return nil, errors.Wrap(err, "failed to set up solr ZK to use HTTPS")
			}
		}
	}

	solrHTTP, err := solr.NewSolrHTTP(cfg.SolrUseHTTPS, cfg.SolrCollection) //, solr.User("solr"), solr.Password("admin"), solr.MinRF(2), solr.QueryRouter(solr.NewRoundRobinRouter()))
	if err != nil {
		return nil, errors.Wrap(err, "failed to create new SolrHTTP client")
	}

	return solr.NewSolrHttpRetrier(solrHTTP, cfg.SolrHTTPRetries, time.Duration(cfg.SolrHTTPRetryBackoffMS)*time.Millisecond), nil
}
