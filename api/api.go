package api

import (
	"net/http"
	"time"

	"github.com/sendgrid/go-solr"
	"github.com/sirupsen/logrus"
)

// Dependencies is the dependencies of the API
type Dependencies struct {
	Log        *logrus.Entry
	SolrClient solr.SolrHTTP
}

// API interface defines the method provided by api package
type API interface {
	Start() error
}

// api is a private struct which implements the API interface through the
type api struct {
	*Dependencies
	*Config
}

// New creates and returns an API
func New(cfg *Config, deps *Dependencies) API {
	return &api{
		Config:       cfg,
		Dependencies: deps,
	}
}

// Start starts a new API server
func (a *api) Start() error {
	server := &http.Server{
		Addr:         a.ServerAddress,
		ReadTimeout:  time.Duration(a.ReadTimeoutMS) * time.Millisecond,
		WriteTimeout: time.Duration(a.WriteTimeoutMS) * time.Millisecond,
		Handler:      a.routes(),
		// MaxHeaderBytes: api.MaxHeaderBytes,
	}

	errChan := make(chan error)
	go func() {
		a.Log.Infof("server started @ %s", a.ServerAddress)
		errChan <- server.ListenAndServe()
	}()

	return <-errChan

}
