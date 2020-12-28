package main

import (
	"flag"
	"log"

	"github.com/go_skill_uper/pkg/apiserver"
	"github.com/go_skill_uper/pkg/storage"
	"github.com/go_skill_uper/pkg/web"
)

func main() {
	// parse values from start arguments
	bindAddr := flag.String("port", ":8080", "Bind address :8080")
	logLevel := flag.String("log", "debug", "Log level :debug")
	dbFlag := flag.String("db", "./test.db", "Database path ./test.db")
	flag.Parse()

	config := apiserver.NewConfig(*bindAddr, *logLevel, *dbFlag)

	// create storage object
	storageUnit := storage.New(config.DBType, config.DBPath)

	log.Println("Server was started on ", *bindAddr)

	//start the rest apiserver
	go apiserver.Start(config, storageUnit)

	web.Start()

}
