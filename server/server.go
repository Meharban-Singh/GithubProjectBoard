package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

var accessToken string

func loginHandler(w http.ResponseWriter, r *http.Request) {
	//If URL is something other than login, return 404
    if r.URL.Path != "/login" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
	}
	
	//If not a GET request, return 404
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
	}
	
	//Get Authorization token from config.json
	data, err := ioutil.ReadFile("./config.json")
    if err != nil {
	  log.Fatal("FATAL: Cannot read configurations!", err)
	}

	//Object representing config.json file
	type ConfigObj struct {
		PAT string
	}
  
	var configs ConfigObj

	// unmarshall json file 
    err = json.Unmarshal(data, &configs)
    if err != nil {
        log.Fatal("FATAL: Cannot get authorization token!", err)
	}
	
	//Save access token 
	accessToken = configs.PAT

	w.WriteHeader(http.StatusOK);
}

func reposHandler(w http.ResponseWriter, r *http.Request) {
	//If URL is something other than /repos, return 404
    if r.URL.Path != "/repos" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
	}
	
	//If not a GET request, return 404
    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
	}
	
	//Get HTTP client 
	client := &http.Client{}
	
	//Create GET request object
	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		log.Fatal(err)
	}
	
	//Set Authorization token header
	req.Header.Set("Authorization", "token " + accessToken)
	
	//Send request
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	
	//Read response and send it 
	data, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	
	w.Write(data)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	
	//GET Repos
	http.HandleFunc("/repos", reposHandler)
	

	//LISTEN AND SERVER SERVER
	fmt.Printf("Starting server at port:1010")
	err := http.ListenAndServe(":1010", nil)
	if err != nil {
		log.Fatal(err)
	} 
}