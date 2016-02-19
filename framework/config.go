package framework

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Parser must implement ParseJSON
type ConfigInterface interface {
	ParseJSON([]byte) error
}

// Config contains the application settings
type Config struct {
}

// ParseJSON unmarshals bytes to structs
func (c *Config) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// Load the JSON config file
func Load(configFile string, p Config) {
	var err error
	var input = io.ReadCloser(os.Stdin)
	if input, err = os.Open(configFile); err != nil {
		log.Fatalln(err)
	}

	// Read the config file
	jsonBytes, err := ioutil.ReadAll(input)
	input.Close()
	if err != nil {
		log.Fatalln(err)
	}

	// Parse the config
	if err := p.ParseJSON(jsonBytes); err != nil {
		log.Fatalln("Could not parse %q: %v", configFile, err)
	}
}
