package config

import (
	"encoding/json"
	//"github.com/davecgh/go-spew/spew"
	"github.com/skmetaly/pbblog/framework/database"
	"github.com/skmetaly/pbblog/framework/session"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//ConfigInterface Parser must implement ParseJSON
type ConfigInterface interface {
	ParseJSON([]byte) error
}

//Config contains the application settings
type Config struct {
	DatabaseConfig database.DatabaseConfig
	SessionConfig  session.SessionConfig
}

//Return the json content of files
func (c *Config) getConfigJSON(configFolder string, configFile string) ([]byte, error) {
	var err error
	var input = io.ReadCloser(os.Stdin)
	configPath, err := filepath.Abs(configFolder + "/" + configFile + ".json")

	if err != nil {
		log.Fatalln("Could not parse %q: %v", configPath, err)
	}

	if input, err = os.Open(configPath); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	return jsonBytes, err
}

//Load loads all known configs from files
func (c *Config) Load(configFolder string) {
	var err error
	configFile := "database"

	jsonBytes, _ := c.getConfigJSON(configFolder, configFile)
	err = json.Unmarshal(jsonBytes, &c.DatabaseConfig)

	// Parse the config database
	if err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}

	configFile = "session"

	jsonBytes, _ = c.getConfigJSON(configFolder, configFile)
	err = json.Unmarshal(jsonBytes, &c.SessionConfig)

	// Parse the session config
	if err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}

//NewConfig returns a new config instance
func NewConfig() Config {
	c := Config{}
	c.Load("config")

	return c
}
