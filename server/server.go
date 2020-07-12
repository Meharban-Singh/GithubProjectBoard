package main

import (
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"

	"github.com/labstack/echo"
)

var (
	accessToken string
)

//TODO: get a proper access token after registering for client ID
func getAccessToken() {
	//Object representing config.json file
	type ConfigObj struct {
		PAT string
	}
	
	var config ConfigObj
	
	//Get Authorization token from config.json
	data, err := ioutil.ReadFile("./config.json")
	if err != nil {
	  log.Fatal("FATAL: Cannot read configurations!", err)
	}
	
	// unmarshall json file 
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal("FATAL: Cannot get authorization token!", err)
	}
	
	//Save access token 
	accessToken = config.PAT
}

func sendGETReqToGH(url string, c echo.Context) error {
	//Get HTTP client 
	client := &http.Client{}

	//Create GET request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot access GH account!");
	}
	
	//Set Authorization token header
	req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")
	req.Header.Set("Authorization", "token " + accessToken)
	
	//Send request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot access GH account!");
	}

	//Read response and send it 
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot read response from GH!");
	}

	// All ok, sned 400
	return c.String(http.StatusOK, string(data))
}

// Handles login requests - stores access token
// URL '/login'
// Returns status codes 400 or 401
func loginHandler(c echo.Context) error {
	getAccessToken();
	
	// If got status code, return 400 else return 401
	if accessToken != "" {
		return c.String(http.StatusOK, "SUCCESS");
	} else {
		return c.String(http.StatusUnauthorized, "REJECTED")
	}
}

// Returns a list of repo's for authenticated user
// URL: /repos
// Returns status codes 400 or 500 for errors
func getAllRepos(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/user/repos", c)
}

// Returns all projects of some repo
// User must be authenticated before
// URL: /repos/:ownerOfRepo/:repoName/projects
// Returns status codes 400 or 500
func getProjectsOfRepo(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/repos/" + c.Param("user") + "/" + c.Param("repo") + "/projects", c)
}

func getProjectDetails(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/projects/" + c.Param("projectID"), c)
}

// Main function
func main() {
	// Create a new echo object
	app := echo.New()

	// Routes
	app.GET("/login", loginHandler)
	app.GET("/repos/:user/:repo/projects", getProjectsOfRepo)
	app.GET("/projects/:projectID", getProjectDetails)
	app.GET("/repos", getAllRepos)

	//Start server
	app.Logger.Fatal(app.Start(":1010"))
}