package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/go_skill_uper/pkg/apiserver"
	"github.com/go_skill_uper/pkg/storage"
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

	fmt.Print("Server was started on ", *bindAddr)

	//start the rest apiserver
	if err := apiserver.Start(config, storageUnit); err != nil {
		log.Fatal(err)
	}
}
