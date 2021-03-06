package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config struct holds the application's configuration
type Config struct {
	HerokuAPIKey                   string
	HerokuAPIEndpoint              string   `json:"herokuAPIEndpoint"`
	HerokuTeam                     string   `json:"herokuTeam"`
	ContextName                    string   `json:"contextName"`
	Envs                           []string `json:"envs"`
	Buildpacks                     []string `json:"buildpacks"`
	LogentriesAPIKey               string
	LogentriesAPIEndpoint          string   `json:"logentriesAPIEndpoint"`
	LogentriesLogsetKey            string
	LogentriesHerokuLogStructureID string   `json:"logentriesHerokuLogStructureID"`
	RollbarAPIEndpoint             string   `json:"rollbarAPIEndpoint"`
	RollbarAccountAccessToken      string
}

// Reader is used by other modules to read the application's configuration
var Reader Config

// InitConfig simply parses the `.config.json` file into a Config struct
// and loads the struct into the Reader global variable
func InitConfig() {
	var config Config
	file, err := os.Open("config/.config.json")
	defer file.Close()

	if err != nil {
		log.Println("Error reading file ", err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		log.Printf("Error decoding JSON: %v\n", err)
	}

	Reader = config
	Reader.HerokuAPIKey = os.Getenv("HEROKU_API_KEY")
	Reader.LogentriesAPIKey = os.Getenv("LOGENTRIES_API_KEY")
	Reader.LogentriesLogsetKey = os.Getenv("LOGENTRIES_LOGSET_KEY")
	Reader.RollbarAccountAccessToken = os.Getenv("ROLLBAR_ACCOUNT_ACCESS_TOKEN")
}
