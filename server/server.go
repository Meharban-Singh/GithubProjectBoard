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

func loginHandler(c echo.Context) error {
	getAccessToken();
	
	if accessToken != "" {
		return c.String(http.StatusOK, "SUCCESS");
	} else {
		return c.String(http.StatusUnauthorized, "REJECTED")
	}
}

func getAllRepos(c echo.Context) error {
	//Get HTTP client 
	client := &http.Client{}
	
	//Create GET request object
	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot access GB account!");
	}
	
	//Set Authorization token header
	req.Header.Set("Authorization", "token " + accessToken)
	
	//Send request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot access GB account!");
	}
	
	//Read response and send it 
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot read response from GB!");
	}
	
	return c.String(http.StatusOK, string(data))
}

func main() {
	app := echo.New()

	app.GET("/login", loginHandler)
	//app.GET("/repos/:repo/projects", getProjectsOfRepo)
	app.GET("/repos", getAllRepos)

	//Start server
	app.Logger.Fatal(app.Start(":1010"))
}