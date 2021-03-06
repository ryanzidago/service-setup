# service-setup

This tool:
* creates a pipeline on Heroku, which will be owned by your Heroku Team
* creates a new Rollbar project
* for each environment (staging and production)
  * creates an app on Heroku (owned by your Heroku Team)
  * add the app to the pipeline
  * creates the `heroku-postgresql` add-on
  * installs default buildpacks
  * connects the app to the newly created Rollbar project
  * creates a new log for your team on Logentries
  * configure the Heroku app to drains its log to Logentries

To make it work:
* create the `.config.json` file in the `config/` directory
* the file should look like this:
  ```json
    {
    	"herokuAPIEndpoint": "https://api.heroku.com",
    	"herokuTeam": "<your team name on Heroku>",
    	"contextName": "<the name of the new application>",
    	"envs": [
    		"staging",
    		"production"
    	],
    	"buildpacks": [
    		"<the url of the buildpack>",
    		"<you can add more buildpacks>",
    	],
    	"logentriesAPIEndpoint": "https://rest.logentries.com",
    	"logentriesHerokuLogStructureID": "73de19ab-366b-4aa1-8f2d-d2b2128f1771",
    	"rollbarAPIEndpoint": "https://api.rollbar.com/api/1"
    }
  ```
* you will also need to set the following environment variables:
  * `HEROKU_API_KEY`
  * `LOGENTRIES_API_KEY`
  * `LOGENTRIES_LOGSET_KEY`
  * `ROLLBAR_ACCOUNT_ACCESS_TOKEN`

After you've added your configuration, simple run `go build service_setup.go` to build the binary and then `./service_setup` to run the tool.
