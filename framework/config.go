package config

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Parser must implement ParseJSON
type ConfigInterface interface {
	ParseJSON([]byte) error
}

// configuration contains the application settings
type Config struct {
	/*	Database  database.Databases      `json:"Database"`
		Email     email.SMTPInfo          `json:"Email"`
		Recaptcha recaptcha.RecaptchaInfo `json:"Recaptcha"`
		Server    server.Server           `json:"Server"`
	*/
	Session session.Session `json:"Session"`
	/*	Template  view.Template           `json:"Template"`
		View      view.View               `json:"View"`
	*/
}

// ParseJSON unmarshals bytes to structs
func (c *configuration) ParseJSON(b []byte) error {
	return json.Unmarshal(b, &c)
}

// Load the JSON config file
func Load(configFile string, p Parser) {
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
