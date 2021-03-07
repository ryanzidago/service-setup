package httphelper

import (
  "net/http"
  "encoding/json"
  "log"
  "io/ioutil"
  "fmt"
  "bytes"
  "service_setup/config"
)

func ExecutePostReqAndParseResp(data map[string]interface{}, url string, client string) interface{} {
	jsonData := toJSON(data)
	req := newPostRequest(jsonData, url)
  return executeReqAndParseResp(req, client)
}

func ExecutePatchReqAndParseResp(data map[string]interface{}, url string, client string) interface{}  {
  jsonData := toJSON(data)
  req := newPatchRequest(jsonData, url)
  return executeReqAndParseResp(req, client)
}

func ExecutePutReqAndParseResp(data map[string]interface{}, url string, client string) interface{} {
	jsonData := toJSON(data)
	req := newPutRequest(jsonData, url)
  return executeReqAndParseResp(req, client)
}

func ExecuteGetRequestAndParseResp(url string, client string) interface{} {
	req := newGetRequest(url)
	return executeReqAndParseResp(req, client)
}

func executeReqAndParseResp(req *http.Request, client string) interface{} {
  addHeadersToRequestForClient(req, client)
  resp := sendRequest(req)
  parsedResp := readRespBody(resp)
  log.Println(fmt.Sprint(parsedResp, "\n"))

  return parsedResp
}

func sendRequest(req *http.Request) *http.Response {
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

func newPostRequest(jsonData []byte, url string) *http.Request {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func newPatchRequest(jsonData []byte, url string) *http.Request {
  req, err := http.NewRequest("PATCH", url, bytes.NewBuffer(jsonData))
  if err != nil {
    log.Fatal(err)
  }
  return req
}

func newPutRequest(jsonData []byte, url string) *http.Request {
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonData))
	if err != nil {
		log.Fatal(err)
	}
	return req
}

func newGetRequest(url string) *http.Request  {
  req, err := http.NewRequest("GET", url, nil)
  if err != nil {
    log.Fatal(err)
  }
  return req
}

func readRespBody(resp *http.Response) interface{} {
	bytes, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	var parsedResp interface{}
	_ = json.Unmarshal([]byte(bytes), &parsedResp)
	return parsedResp
}

func toJSON(data interface{}) []byte {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return jsonData
}

func selectAPIEndpointForClient(client string) string {
  switch client {
  case "heroku":
    return config.Reader.HerokuAPIEndpoint
  case "rollbar":
    return config.Reader.RollbarAPIEndpoint
  case "logentries":
    return config.Reader.LogentriesAPIEndpoint
  default:
    return "error"
  }
}

func addHeadersToRequestForClient(req *http.Request, client string)  {
  req.Header.Add("User-Agent", "go version go1.16 linux/amd64")
	req.Header.Add("Content-Type", "application/json")

  switch client {
  case "heroku":
    bearerToken := fmt.Sprint("Bearer ", config.Reader.HerokuAPIKey)
    req.Header.Add("Accept", "application/vnd.heroku+json; version=3")
  	req.Header.Add("Authorization", bearerToken)
  case "rollbar":
    req.Header.Add("X-Rollbar-Access-Token", config.Reader.RollbarAccountAccessToken)
  case "logentries":
    req.Header.Add("x-api-key", config.Reader.LogentriesAPIKey)
  }
}
