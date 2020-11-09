package main

import (
	"log"

	"github.com/go_skill_uper/pkg/apiserver"
	"github.com/go_skill_uper/pkg/storage"
)

//Нужно будет потом раскидать функции на разные модули

func main() {
	config := apiserver.NewConfig()
	s := storage.New(config.DBType, config.DBPath)

	if err := apiserver.Start(config, s); err != nil {
		log.Fatal(err)
	}
}
