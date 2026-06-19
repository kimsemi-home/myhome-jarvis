package daemon

import (
	"errors"
	"fmt"
	"time"
)

func New(config Config) (*Server, error) {
	config = withDefaultResourceBounds(config)
	if config.Root == "" {
		return nil, errors.New("root is required")
	}
	if config.Host == "" {
		config.Host = "127.0.0.1"
	}
	if config.Port <= 0 || config.Port > 65535 {
		return nil, fmt.Errorf("invalid port %d", config.Port)
	}
	if isWildcardHost(config.Host) && !config.AllowLANBind {
		return nil, errors.New("wildcard or LAN bind requires explicit allow-lan flag")
	}
	return &Server{
		config:  config,
		started: time.Now().UTC(),
		events:  newEventLog(maxRequestEvents),
	}, nil
}
