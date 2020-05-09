package api

import "github.com/kelseyhightower/envconfig"

// Config is the configuration of the API
type Config struct {
	SolrServerURL  string `envconfig:"SOLR_SERVER_URL" required:"true"`
	ServerAddress  string `envconfig:"SERVER_ADDRESS" default:":4040"`
	ReadTimeoutMS  int    `envconfig:"READ_TIMEOUT_MS" default:"20000"`
	WriteTimeoutMS int    `envconfig:"WRITE_TIMEOUT_MS" default:"20000"`
}

// ParseConfig parses api config from the environment vars
func ParseConfig(prefix string) (*Config, error) {
	cfg := new(Config)
	return cfg, envconfig.Process(prefix, cfg)
}
