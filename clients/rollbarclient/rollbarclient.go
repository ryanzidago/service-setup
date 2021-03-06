package rollbarclient

import (
	"fmt"
	"log"
	"service_setup/config"
	"service_setup/httphelper"
	"strings"
)

var client string = "rollbar"

// CreateProject creates a new Rollbar project
func CreateProject(contextName string) map[string]interface{} {
	message := fmt.Sprint("Creating new ", contextName, " project on Rollbar ...", "\n")
	log.Println(message)

	projectName := fmt.Sprint(strings.Title(contextName))
	projectData := map[string]interface{}{
		"name": projectName,
	}

	url := fmt.Sprint(config.Reader.RollbarAPIEndpoint, "/projects")
	resp := httphelper.ExecutePostReqAndParseResp(projectData, url, client)
	return resp.(map[string]interface{})
}

// GetPostServerItemAccessTokens fetch the post_servier_item from a Rollbar project
func GetPostServerItemAccessTokens(rollbarProject map[string]interface{}) map[string]interface{} {
	message := fmt.Sprint("Fetching post_server_item (access tokens)", "\n")
	log.Println(message)

	rollbarProjectID := rollbarProject["result"].(map[string]interface{})["id"]
	url := fmt.Sprint(config.Reader.RollbarAPIEndpoint, "/project", "/", rollbarProjectID, "/access_tokens")
	resp := httphelper.ExecuteGetRequestAndParseResp(url, client)
	accessTokens := resp.(map[string]interface{})

	return accessTokens["result"].([]interface{})[0].(map[string]interface{})
}
