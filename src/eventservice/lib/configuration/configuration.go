package configuration

import (
	"encoding/json"
	"fmt"
	"os"
	"eventservice/lib/persistence/dblayer"
)

var (
	DBTypeDefault		= dblayer.DBTYPE("mongodb")
	DBConnectionDefault	= "mongodb://127.0.0.1"
	RestfulEPDefault	= "localhost:8181"
)

type ServiceConfig struct {
	DatabaseType	dblayer.DBTYPE	`json:"databasetype"`
	DBConnection	string		`json:"dbconnection"`
	RestfulEndpoint	string		`json:"restfulapi_endpoint"`
}

func ExtractConfiguration(filename string) (ServiceConfig, error) {
	conf := ServiceConfig{
		DBTypeDefault,
		DBConnectionDefault,
		RestfulEPDefault,
	}

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Configuration file not found. Continuing with default values.")
		return conf, err
	}

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		fmt.Println("Error decoding configuration file. Continuing with default values.")
		return conf, err
	}

	return conf, nil
}
