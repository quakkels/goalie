package settings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var environments = map[string]string{
	"production":    "settings/prod.json",
	"preproduction": "settings/preproduction.json",
	"tests":         "settings/tests.json",
}

// Settings has the api's runtime settings
type Settings struct {
	PrivateKeyPath     string
	PublicKeyPath      string
	JWTExpirationDelta int
}

var settings = Settings{}

// Init the api's settings using the API_ENV environment variable
func Init() {
	env := os.Getenv("API_ENV")
	if env == "" {
		fmt.Println("Warning: Setting environment to preproduction due to lack of API_ENV environment variable.")
		env = "preproduction"
	}
	loadSettingsByEnv(env)
}

func loadSettingsByEnv(env string) {
	content, err := ioutil.ReadFile(environments[env])
	if err != nil {
		fmt.Println("Error: Cannot read configuration file. ", err)
	}
	err = json.Unmarshal(content, &settings)
	if err != nil {
		fmt.Println("Error while parsing the configuration file. ", err)
	}
}
