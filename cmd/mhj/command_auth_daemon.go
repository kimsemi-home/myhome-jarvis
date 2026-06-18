package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/kimsemi-home/myhome-jarvis/internal/auth"
	"github.com/kimsemi-home/myhome-jarvis/internal/daemon"
	"github.com/kimsemi-home/myhome-jarvis/internal/supervisor"
)

func runAuth(root string, args []string) error {
	if len(args) == 1 && args[0] == "status" {
		return writeJSON(auth.Status(root))
	}
	if len(args) == 2 && args[0] == "token" && args[1] == "create" {
		result, err := auth.Create(root, false)
		if err != nil {
			return err
		}
		return writeJSON(result)
	}
	if len(args) == 2 && args[0] == "token" && args[1] == "rotate" {
		result, err := auth.Create(root, true)
		if err != nil {
			return err
		}
		return writeJSON(result)
	}
	return errors.New("usage: mhj auth <status|token create|token rotate>")
}

func runDaemon(root string, args []string) error {
	if len(args) == 1 && args[0] == "status" {
		return writeJSON(supervisor.Status(root, nil))
	}
	config := daemon.DefaultConfig(root, version)
	flags := flag.NewFlagSet("daemon", flag.ContinueOnError)
	flags.SetOutput(os.Stderr)
	flags.StringVar(&config.Host, "host", config.Host, "bind host")
	flags.IntVar(&config.Port, "port", config.Port, "bind port")
	flags.BoolVar(&config.AllowLANBind, "allow-lan", false, "allow non-localhost bind")
	flags.BoolVar(&config.Execute, "execute", config.Execute, "allow explicit execute requests")
	if err := flags.Parse(args); err != nil {
		return err
	}
	server, err := daemon.New(config)
	if err != nil {
		return err
	}
	fmt.Fprintf(os.Stderr, "myhome-jarvis daemon listening on %s:%d\n", config.Host, config.Port)
	return server.ListenAndServe()
}
