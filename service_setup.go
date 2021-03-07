package main

import (
	"fmt"
	"log"
	"service_setup/clients/herokuclient"
	"service_setup/clients/logentriesclient"
	"service_setup/clients/rollbarclient"
	"service_setup/config"
)

func main() {
	log.Println("Initializing configurations ...")
	config.InitConfig()

	// owner of the pipeline
	// will be the Heroku Team
	team := herokuclient.GetTeam()
	pipeline := herokuclient.CreatePipeline(config.Reader.ContextName, team)

	// creating the Rollbar project
	rollbarProject := rollbarclient.CreateProject(config.Reader.ContextName)

	for _, env := range config.Reader.Envs {
		// bootstraping Heroku Apps and putting them in a pipeline
		appName := fmt.Sprint(config.Reader.ContextName, "-", env)
		app := herokuclient.CreateApp(appName)
		herokuclient.CoupleAppWithPipeline(app, pipeline, env)
		herokuclient.CreateAddon(app)
		herokuclient.AddBuildpacks(app, config.Reader.Buildpacks)

		// getting access token from newly created Rollbar project
		rollbarAccessToken := rollbarclient.GetPostServerItemAccessTokens(rollbarProject)
		herokuclient.ConfigureRollbar(rollbarAccessToken, env, app)

		// creating a Logentries log
		// and configuring draining logs from Heroku
		logData := logentriesclient.CreateLog(config.Reader.ContextName, env)
		herokuclient.CreateLogDrain(app, logData)

		log.Println("Project ", config.Reader.ContextName, " is bootstrapped.", "\n")

	}
}
