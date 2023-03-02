package main

import (
	"flag"
	config "github.com/gusleein/goconfig"
	log "github.com/gusleein/golog"
	"github.com/inspectorvitya/collector-data/api"
	"github.com/inspectorvitya/collector-data/db"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Init(true, "console")
	env := flag.String("env", "config", "")
	flag.Parse()
	config.Init(*env)
	db.Init()

	go api.Start()

	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-osSignals:
		api.Stop()
		db.Stop()
	}

}
