package herokuclient

import (
	"fmt"
	"log"
	"service_setup/config"
	"service_setup/httphelper"
)

var client string = "heroku"

func GetTeam() map[string]interface{} {
	message := fmt.Sprint("Getting Heroku Team ...", "\n")
	log.Println(message)

	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/teams", "/", config.Reader.HerokuTeam)
	resp := httphelper.ExecuteGetRequestAndParseResp(url, client)
	return resp.(map[string]interface{})
}

func CreateAddon(app map[string]interface{}) map[string]interface{} {
	message := fmt.Sprint("Adding heroku-postgresql to ", app["name"], "\n")
	log.Println(message)

	addonData := map[string]interface{}{
		"plan":    "heroku-postgresql",
		"confirm": app["name"],
	}

	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/apps", "/", app["name"], "/addons")
	resp := httphelper.ExecutePostReqAndParseResp(addonData, url, client)
	return resp.(map[string]interface{})
}

func AddBuildpacks(app map[string]interface{}, buildpacks []string) []interface{} {
	message := fmt.Sprint("Adding default buildpacks to ", app["name"], "\n")
	log.Println(message)

	var buildpackMaps [3]map[string]interface{}

	for index := range buildpackMaps {
		buildpackMaps[index] = map[string]interface{}{"buildpack": buildpacks[index]}
	}

	buildpackData := map[string]interface{}{
		"updates": buildpackMaps,
	}

	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/apps", "/", app["id"], "/", "buildpack-installations")
	resp := httphelper.ExecutePutReqAndParseResp(buildpackData, url, client)
	return resp.([]interface{})
}

func GetAccountInfo() map[string]interface{} {
	message := fmt.Sprint("Getting account information ", "\n")
	log.Println(message)

	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/account")
	resp := httphelper.ExecuteGetRequestAndParseResp(url, client)
	return resp.(map[string]interface{})
}

func CreatePipeline(pipelineName string, owner map[string]interface{}) map[string]interface{} {
	message := fmt.Sprint("Creating pipeline ", "\n")
	log.Println(message)

	pipelineData := map[string]interface{}{
		"name": pipelineName,
		"owner": map[string]interface{}{
			"id":   owner["id"],
			"type": "team",
		},
	}

	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/pipelines")
	resp := httphelper.ExecutePostReqAndParseResp(pipelineData, url, client)
	return resp.(map[string]interface{})
}

func CreateApp(contextName string) map[string]interface{} {
	message := fmt.Sprint("Creating ", contextName, " application", "\n")
	log.Println(message)

	appData := map[string]interface{}{
		"name":   contextName,
		"region": "eu",
		"team":   config.Reader.HerokuTeam,
	}

	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/teams/apps")
	resp := httphelper.ExecutePostReqAndParseResp(appData, url, client)
	return resp.(map[string]interface{})
}

func CoupleAppWithPipeline(app map[string]interface{}, pipeline map[string]interface{}, stage string) map[string]interface{} {
	message := fmt.Sprint("Coupling ", app["name"], " with pipeline ...", "\n")
	log.Println(message)

	pipelineCouplingData := map[string]interface{}{
		"app":      app["id"],
		"pipeline": pipeline["id"],
		"stage":    stage,
	}

	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/pipeline-couplings")
	resp := httphelper.ExecutePostReqAndParseResp(pipelineCouplingData, url, client)
	return resp.(map[string]interface{})
}

func ConfigureRollbar(rollbarAccessToken map[string]interface{}, env string, app map[string]interface{}) map[string]interface{} {
	message := fmt.Sprint("Configuring Rollbar on Heroku ...", "\n")
	log.Println(message)

	envVars := map[string]interface{}{
		"ROLLBAR_ACESS_TOKEN": rollbarAccessToken["access_token"],
		"ROLLBAR_ENVIRONMENT": env,
	}
	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/apps", "/", app["id"], "/config-vars")
	resp := httphelper.ExecutePatchReqAndParseResp(envVars, url, client)
	return resp.(map[string]interface{})
}

func CreateLogDrain(app, logData map[string]interface{}) map[string]interface{} {
	message := fmt.Sprint("Configuring log draining from Heroku to Logentries ...", "\n")
	log.Println(message)

	logToken := logData["log"].(map[string]interface{})["tokens"].([]interface{})[0]
	logDrainURL := fmt.Sprint(config.Reader.LogentriesAPIEndpoint, "/v1/drains", "/", logToken)
	logDrainData := map[string]interface{}{
		"url": logDrainURL,
	}

	url := fmt.Sprint(config.Reader.HerokuAPIEndpoint, "/apps", "/", app["id"], "/log-drains")
	resp := httphelper.ExecutePostReqAndParseResp(logDrainData, url, client)
	return resp.(map[string]interface{})
}
