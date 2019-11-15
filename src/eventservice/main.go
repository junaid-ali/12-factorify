package main

import (
	"flag"
	"fmt"
	"log"
	"eventservice/rest"
	"eventservice/lib/configuration"
	//"eventservice/lib/persistence"
	"eventservice/lib/persistence/dblayer"
)

func main() {
	confPath := flag.String("conf", `./configuration/config.json`, "flag to set the path to the configuration json file")
	flag.Parse()

	config, _ := configuration.ExtractConfiguration(*confPath)
	fmt.Println(config.DatabaseType)

	fmt.Println("Connecting to the database")
	dbHandler, err := dblayer.NewPersistenceLayer(config.DatabaseType, config.DBConnection)
	if err != nil {
		log.Fatal("Connection to db failed")
	}
	//start RESTful API
	log.Fatal(rest.ServeAPI(config.RestfulEndpoint, dbHandler))
}
