package main

import (
	"fmt"
	"os"

	"github.com/linushsianghung/cmd/server"
	"github.com/linushsianghung/internal/app/persistent"
	"github.com/linushsianghung/internal/app/persistent/mongodb"
	"github.com/linushsianghung/internal/pkg/configs"

	log "github.com/sirupsen/logrus"
)

func initLogrus() {
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

func main() {
	log.Infof("***** [INIT:AMAZINGSERVER] ***** Start to launch server 🤓 ...")
	initLogrus()
	configs.InitConfig()
	port := configs.GetConfigStr("service.rest.port")

	var operator persistent.NoSqlOperator
	switch configs.GetConfigStr("connection.nosql.type") {
	case "MongoDB":
		connStr := server.FetchConnectionString(configs.GetConfigBool("connection.nosql.local"))
		operator = mongodb.NewOperator(connStr)
	case "Cassandra":
		fmt.Println("Waiting for implementation")
		os.Exit(1)
	default:
		fmt.Println("Using local memory as temporary persistent")
	}

	s := server.NewHTTPServer(port, operator)
	server.InitServer(s)
}
