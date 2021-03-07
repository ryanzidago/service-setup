package logentriesclient

import (
	"fmt"
	"log"
	"service_setup/config"
	"service_setup/httphelper"
	"strings"
)

var client string = "logentries"

func CreateLog(contextName string, env string) map[string]interface{} {
	message := fmt.Sprint("Creating log in Logentries for ", contextName, " ", env, "\n")
	log.Println(message)

	logName := fmt.Sprint(strings.Title(contextName), " ", strings.Title(env))

	postData := map[string]interface{}{
		"log": map[string]interface{}{
			"name": logName,
			"structures": [1]string{
				config.Reader.LogentriesHerokuLogStructureID,
			},
			"user_data":   map[string]string{},
			"source_type": "token",
			"tokens":      [0]string{},
			"logsets_info": [1]map[string]string{
				map[string]string{
					"id": config.Reader.TeamLogsetKey,
				},
			},
		},
	}

	url := fmt.Sprint(config.Reader.LogentriesAPIEndpoint, "/management/logs")
	resp := httphelper.ExecutePostReqAndParseResp(postData, url, client)
	return resp.(map[string]interface{})
}
