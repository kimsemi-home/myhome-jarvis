package daemon

import (
	"os"
	"strconv"
	"strings"
)

func DefaultConfig(root string, version string) Config {
	config := Config{
		Root:    root,
		Host:    envString("MYHOME_BIND_HOST", "127.0.0.1"),
		Port:    envInt("MYHOME_BIND_PORT", 3888),
		Execute: os.Getenv("MYHOME_EXECUTE") == "true",
		Version: version,
	}
	return withDefaultResourceBounds(config)
}

func withDefaultResourceBounds(config Config) Config {
	if config.ReadHeaderTimeout <= 0 {
		config.ReadHeaderTimeout = defaultReadHeaderTimeout
	}
	if config.ReadTimeout <= 0 {
		config.ReadTimeout = defaultReadTimeout
	}
	if config.WriteTimeout <= 0 {
		config.WriteTimeout = defaultWriteTimeout
	}
	if config.IdleTimeout <= 0 {
		config.IdleTimeout = defaultIdleTimeout
	}
	if config.MaxHeaderBytes <= 0 {
		config.MaxHeaderBytes = defaultMaxHeaderBytes
	}
	return config
}

func envString(name string, fallback string) string {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	return value
}

func envInt(name string, fallback int) int {
	value := strings.TrimSpace(os.Getenv(name))
	if value == "" {
		return fallback
	}
	parsed, err := strconv.Atoi(value)
	if err != nil {
		return fallback
	}
	return parsed
}

func isWildcardHost(host string) bool {
	return host == "0.0.0.0" || host == "::" || host == ""
}
