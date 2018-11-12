package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/BurntSushi/toml"
)

/*
Config holds the configuration for the server
*/
type Config struct {
	Port   int
	AppDir string
}

/*
NewConfig constructs a configuraiton option with set defaults
*/
func NewConfig() Config {
	return Config{8080, "app"}
}

func readConfig(configFile string) (Config, error) {
	config := Config{8080, "app"}
	_, err := toml.DecodeFile(configFile, &config)
	return config, err
}

func main() {
	configFile := "config.toml"
	if len(os.Args) == 2 {
		configFile = os.Args[1]
	}
	config, err := readConfig(configFile)
	if err != nil {
		log.Fatalf("Problem reading config file at %v: %v", configFile, err)
	}
	fs := http.FileServer(http.Dir(config.AppDir))
	http.Handle("/app/", http.StripPrefix("/app/", fs))
	addr := fmt.Sprintf(":%v", config.Port)
	log.Printf("Starting server on %v", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
