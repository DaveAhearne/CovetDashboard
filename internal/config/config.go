package config

import (
	"flag"
	"os"
)

var AppConfig Config

type Config struct {
	ApplicationPort  string
	ApplicationHost  string
	DatabaseUsername string
	DatabasePassword string
	DatabaseHost     string
	DatabasePort     string
	DatabaseName     string
	ExternalAddress  string
	ExternalPort     string
}

func NewConfig() Config {
	var port string
	var address string

	flag.StringVar(&port, "port", "1234", "the port to run on")
	flag.StringVar(&address, "address", "localhost", "the address to run on")
	flag.Parse()

	return Config{
		ApplicationPort:  port,
		ApplicationHost:  address,
		DatabaseUsername: os.Getenv("db_username"),
		DatabasePassword: os.Getenv("db_password"),
		DatabaseHost:     os.Getenv("db_host"),
		DatabasePort:     os.Getenv("db_port"),
		DatabaseName:     os.Getenv("db_name"),
		ExternalAddress:  os.Getenv("external_address"),
		ExternalPort:     os.Getenv("external_port"),
	}
}
