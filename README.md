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
    	"herokuAPIKey": "<your Heroku api key>",
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
    	"logentriesAPIKey": "<your Logentries API key>",
    	"logentriesAPIEndpoint": "https://rest.logentries.com",
    	"teamLogsetKey": "<the key of your team's logset on Logentries>",
    	"logentriesHerokuLogStructureID": "73de19ab-366b-4aa1-8f2d-d2b2128f1771",
    	"rollbarAPIEndpoint": "https://api.rollbar.com/api/1",
    	"rollbarAccountAccessToken": "<your Rollbar account access token>"
    }
  ```

After you've added your configuration, simple run `./service-setup` to run the tool.
