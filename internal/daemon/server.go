package daemon

import (
	"sync/atomic"
	"time"

	"github.com/kimsemi-home/myhome-jarvis/internal/commands"
)

const (
	defaultReadHeaderTimeout = 5 * time.Second
	defaultReadTimeout       = 15 * time.Second
	defaultWriteTimeout      = 30 * time.Second
	defaultIdleTimeout       = 60 * time.Second
	defaultMaxHeaderBytes    = 1 << 20
)

type Config struct {
	Root              string
	Host              string
	Port              int
	Execute           bool
	AllowLANBind      bool
	Version           string
	CommandPlatform   string
	CommandRunner     commands.Runner
	ReadHeaderTimeout time.Duration
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	IdleTimeout       time.Duration
	MaxHeaderBytes    int
}

type Server struct {
	config   Config
	started  time.Time
	requests atomic.Uint64
	events   *eventLog
}
