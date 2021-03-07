package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Config struct holds the application's configuration
type Config struct {
	HerokuAPIKey                   string   `json:herokuAPIKey`
	HerokuAPIEndpoint              string   `json:herokuAPIEndpoint`
	HerokuTeam                     string   `json:herokuTeam`
	ContextName                    string   `json:contextName`
	Envs                           []string `json:envs`
	Buildpacks                     []string `json:buildpacks`
	LogentriesAPIKey               string   `json:logentriesAPIKey`
	LogentriesAPIEndpoint          string   `json:logentriesAPIEndpoint`
	TeamLogsetKey                  string   `json:teamLogsetKey`
	LogentriesHerokuLogStructureID string   `json:logentriesHerokuLogStructureID`
	RollbarAPIEndpoint             string   `json:rollbarAPIEndpoint`
	RollbarAccountAccessToken      string   `json:rollbarAccountAccessToken`
}

// Reader is used by other modules to read the application's configuration
var Reader Config

// InitConfig simply parses the `.config.json` file
// into a Config struct
// and loads the struct into the Reader global variable
func InitConfig() {
	file, err := ioutil.ReadFile("config/.config.json")
	if err != nil {
		fmt.Println("Err")
	}

	var config Config
	json.Unmarshal([]byte(file), &config)
	Reader = config
}
