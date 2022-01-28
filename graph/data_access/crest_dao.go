package data_access

import (
	"github.com/cryanbrow/eve-graphql-go/graph/configuration"
	log "github.com/sirupsen/logrus"
)

var (
	baseUriESI string
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	baseUriESI = configuration.AppConfig.Esi.Default.Url
}
