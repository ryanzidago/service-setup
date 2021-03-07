package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

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

var Reader Config

func InitConfig() {
	file, err := ioutil.ReadFile("config/.config.json")
	if err != nil {
		fmt.Println("Err")
	}

	var config Config
	json.Unmarshal([]byte(file), &config)
	Reader = config
}
