package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatih/structs"
	"github.com/flock3/ip2country/config"
	"github.com/infinityworks/go-common/logger"
	"github.com/infinityworks/go-common/router"
	"github.com/sirupsen/logrus"
)

var (
	applicationCfg config.Config
	eLogger        *logrus.Logger
)

// Exporter Sets up all the runtime and metrics
type Exporter struct {
	config.Config
}

func init() {
	applicationCfg = config.Init()
	eLogger = logger.Start(&applicationCfg)
}

func main() {
	eLogger.WithFields(structs.Map(applicationCfg)).Info("Starting application")

	binding := fmt.Sprintf(":%s", applicationCfg.ListenPort())

	h := Controller{
		Config: applicationCfg,
	}

	r := router.NewRouter(eLogger, createRoutes(h))
	log.Fatal(http.ListenAndServe(binding, r))
}
