package api

import "github.com/kelseyhightower/envconfig"

// Config is the configuration of the API
type Config struct {
	ServerAddress     string `envconfig:"SERVER_ADDRESS" default:":80"`
	ReadTimeoutMS     int    `envconfig:"READ_TIMEOUT_MS" default:"20000"`
	WriteTimeoutMS    int    `envconfig:"WRITE_TIMEOUT_MS" default:"20000"`
	LogLevel          string `envconfig:"LOG_LEVEL" default:"debug"`
	SolrServerURL     string `envconfig:"SOLR_SERVER_URL" required:"true"`
	SolrZooKeepers    string `envconfig:"SOLR_ZOO_KEEPERS" required:"true"`
	SolrZooKeeperRoot string `envconfig:"SOLR_ZOO_KEEPER_ROOT" required:"true"`
	SolrCollection    string `envconfig:"SOLR_COLLECTION" required:"true"`
}

// ParseConfig parses api config from the environment vars
func ParseConfig(prefix string) (*Config, error) {
	cfg := new(Config)
	return cfg, envconfig.Process(prefix, cfg)
}
