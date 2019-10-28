package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// ModeProd defines production mode
const ModeProd = "prod"

// Config defines configuration parameters
type Config struct {
	DbDSN    string `envconfig:"DB_DSN" default:"root:root@tcp(db_container:3306)/backend?charset=utf8mb4"`
	HTTPport string `envconfig:"HTTP_PORT" default:"8080"`
	LocAddr  string `envconfig:"LOCATION_ADDRESS" default:":5001"`
	Mode     string `envconfig:"MODE" default:"dev"`
}

// LoadConfigData loads environment parameters
func LoadConfigData() Config {
	var cnf Config
	if err := envconfig.Process("geoapi", &cnf); err != nil {
		panic(fmt.Sprintf("Failed reading environment variables: %s", err))
	}
	return cnf
}
