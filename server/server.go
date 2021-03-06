package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

var (
	accessToken string
)

type (
	// Project struct
	Project struct {
		Name string `json:"name" xml:"name" form:"name" query:"name"`
		Body string `json:"body" xml:"body" form:"body" query:"body"`
	}

	// Column struct
	Column struct {
		Name string `json:"name" xml:"name" form:"name" query:"name"`
	}

	// Card struct
	Card struct {
		Note string `json:"note" xml:"note" form:"note" query:"note"`
	}

	// Card struct
	Moves struct {
		Position string `json:"position" xml:"position" form:"position" query:"position"`
	}

	// Config struct -> should be same as config.js
	ConfigObj struct {
		PAT string
	}
)

//TODO: get a proper access token after registering for client ID
func getAccessToken() {
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

// Sends a GET request to GitHub with URL provided
func sendGETReqToGH(url string, c echo.Context) error {
	//Create GET request object
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot access GH account!")
	}
	return sendReqToGH(req, c, http.StatusOK)
}

// Sends a POST req to GH
func sendReqToGH(req *http.Request, c echo.Context, statusCode int) error {
	//Set Authorization token header
	req.Header.Set("Accept", "application/vnd.github.inertia-preview+json")
	req.Header.Set("Authorization", "token "+accessToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot read response!")
	}

	defer resp.Body.Close()

	// Read from body of response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot read response!")
	}

	// All ok, sned 400
	return c.String(statusCode, string(body))
}

// Handles login requests - stores access token
// GET '/login'
// Returns 200 OK on success
func loginHandler(c echo.Context) error {
	getAccessToken()

	// If got status code, return 400 else return 401
	if accessToken != "" {
		return c.String(http.StatusOK, "SUCCESS")
	} else {
		return c.String(http.StatusUnauthorized, "REJECTED")
	}
}

// Returns a list of repo's for authenticated user
// GET /repos
// Returns 200 OK on success
func getAllRepos(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/user/repos", c)
}

// Creates a new project in the repo of the user
// POST /repos/:user/:repo/projects
// Returns 201 Created on success
// Req body = { "name" : "<name>", "body" : "<OPTIONAL body description>"}
func createNewProject(c echo.Context) error {
	// Create a new Project obj -> to be sent to GH in POST req
	p := new(Project)
	if err := c.Bind(p); err != nil {
		return err
	}

	// Convert p to JSON
	jsonObj, err := json.Marshal(p)
	if err != nil {
		print(err)
	}

	// Send POST req to GH with project object
	req, err := http.NewRequest("POST", "https://api.github.com/repos/"+c.Param("user")+"/"+c.Param("repo")+"/projects", bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusCreated)
}

// Returns all projects of some repo
// User must be authenticated before
// GET /repos/:ownerOfRepo/:repoName/projects
// Returns 200 OK on success
func getProjectsOfRepo(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/repos/"+c.Param("user")+"/"+c.Param("repo")+"/projects", c)
}

// Returns all details for a project
// GET /projects/:projectID
func getProjectDetails(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/projects/"+c.Param("projectID"), c)
}

// Updates a Project
// PATCH /projects/:projectID
// Returns 200 OK on success
// Req body = { "name" : "<OPTIONAL name>", "body" : "<OPTIONAL body description>"}
func updateProject(c echo.Context) error {
	// Create new Project object
	project := new(Project)
	if err := c.Bind(project); err != nil {
		return err
	}

	// Convert card to JSON
	jsonObj, err := json.Marshal(project)
	if err != nil {
		print(err)
	}

	// Send POST req to GH with card object
	req, err := http.NewRequest("PATCH", "https://api.github.com/projects/"+c.Param("projectID"), bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusOK)
}

// Deletes a project
// DELETE /projects/:projectID
// Returns 204 No Content on success
func deleteProject(c echo.Context) error {
	// Send POST req to GH with project object
	req, err := http.NewRequest("DELETE", "https://api.github.com/projects/"+c.Param("projectID"), nil)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusNoContent)
}

// Creates a project column
// POST /projects/:projectID/columns
// Returns 200 OK on success
// Req body = { "name" : "<name>" }
func createNewColumn(c echo.Context) error {
	// Create new column object
	col := new(Column)
	if err := c.Bind(col); err != nil {
		return err
	}

	// Convert col to JSON
	jsonObj, err := json.Marshal(col)
	if err != nil {
		print(err)
	}

	// Send POST req to GH with column object
	req, err := http.NewRequest("POST", "https://api.github.com/projects/"+c.Param("projectID")+"/columns", bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusCreated)
}

// Returns all Columns of some project
// User must be authenticated before
// GET /project/:projectID/columns
// Returns 200 OK on success
func getColumnsOfProject(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/projects/"+c.Param("projectID")+"/columns", c)
}

// Returns all details for a column
// GET /columns/:columnID
func getColumnDetails(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/projects/columns/"+c.Param("columnID"), c)
}

// Updates a column
// PATCH /column/:columnID
// Returns 200 OK on success
// Req body = { "name" : "<name>" }
func updateColumn(c echo.Context) error {
	// Create new card object
	column := new(Column)
	if err := c.Bind(column); err != nil {
		return err
	}

	// Convert card to JSON
	jsonObj, err := json.Marshal(column)
	if err != nil {
		print(err)
	}

	// Send POST req to GH with card object
	req, err := http.NewRequest("PATCH", "https://api.github.com/projects/columns/"+c.Param("columnID"), bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusOK)
}

// Deletes a project column
// DELETE /projects/columns/:column_id
// Returns 204 No Content on success
func deleteColumn(c echo.Context) error {
	// Send DELETE request to GH with columnID
	req, err := http.NewRequest("DELETE", "https://api.github.com/projects/columns/"+c.Param("columnID"), nil)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusNoContent)
}

// Moves a Column to a specific position
// POST /columns/:columnID/moves
// Returns 200 OK on success
// Req body = { "position" : "first"} 
// Other options for position: last, after:colID
func moveColumn(c echo.Context) error {
	// Create new moves object
	moves := new(Moves)
	if err := c.Bind(moves); err != nil {
		return err
	}

	// Convert moves to JSON
	jsonObj, err := json.Marshal(moves)
	if err != nil {
		print(err)
	}

	// Send POST req to GH with moves object
	req, err := http.NewRequest("POST", "https://api.github.com/projects/columns/"+c.Param("columnID")+"/moves", bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusCreated)
}

// Creates a project card
// POST /columns/:column_id/cards
// Returns 201 Created on success
// Req body = { "note" : "<new card note>" }
func createNewCard(c echo.Context) error {
	// Create new card object
	card := new(Card)
	if err := c.Bind(card); err != nil {
		return err
	}

	// Convert card to JSON
	jsonObj, err := json.Marshal(card)
	if err != nil {
		print(err)
	}

	// Send POST req to GH with card object
	req, err := http.NewRequest("POST", "https://api.github.com/projects/columns/"+c.Param("columnID")+"/cards", bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusCreated)
}

// Returns all cards of some column
// User must be authenticated before
// GET /column/:columnID/cards
// Returns 200 OK on success
func getCardsOfColumn(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/projects/columns/"+c.Param("columnID")+"/cards", c)
}

// Returns all details for a card
// GET /cards/:cardID
func getCardDetails(c echo.Context) error {
	return sendGETReqToGH("https://api.github.com/projects/columns/cards/"+c.Param("cardID"), c)
}

// Updates a card
// PATCH /cards/:cardID
// Returns 200 OK on success
// Req body = { "note" : "<new card note>" }
func updateCard(c echo.Context) error {
	// Create new card object
	card := new(Card)
	if err := c.Bind(card); err != nil {
		return err
	}

	// Convert card to JSON
	jsonObj, err := json.Marshal(card)
	if err != nil {
		print(err)
	}

	// Send POST req to GH with card object
	req, err := http.NewRequest("PATCH", "https://api.github.com/projects/columns/cards/"+c.Param("cardID"), bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusOK)
}

// Delete a project card
// DELETE /cards/:card_id
// Returns Status: 204 No Content on success
func deleteCard(c echo.Context) error {
	// Send DELETE request to GH with cardID
	req, err := http.NewRequest("DELETE", "https://api.github.com/projects/columns/cards/"+c.Param("cardID"), nil)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusNoContent)
}

// Moves a Card to a specific position
// POST /cards/:cardID/moves
// Returns 200 OK on success
// Req body = { "position" : "top"} 
// Other options for position: bottom, after:cardID
func moveCard(c echo.Context) error {
	// Create new moves object
	moves := new(Moves)
	if err := c.Bind(moves); err != nil {
		return err
	}

	// Convert moves to JSON
	jsonObj, err := json.Marshal(moves)
	if err != nil {
		print(err)
	}

	// Send POST req to GH with moves object
	req, err := http.NewRequest("POST", "https://api.github.com/projects/columns/cards/"+c.Param("cardID")+"/moves", bytes.NewBuffer(jsonObj))
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "FATAL: Cannot send request!")
	}
	return sendReqToGH(req, c, http.StatusCreated)
}

// Main function
func main() {
	// Create a new echo object
	app := echo.New()

	// Routes
	app.GET("/login", loginHandler)

	app.GET("/repos", getAllRepos)
	
	// PROJECTS 
	app.POST("/repos/:user/:repo/projects", createNewProject)
	app.GET("/repos/:user/:repo/projects", getProjectsOfRepo)
	app.GET("/projects/:projectID", getProjectDetails)
	app.PATCH("/projects/:projectID", updateProject)
	app.DELETE("/projects/:projectID", deleteProject)
	
	// COLOUMNS 
	app.POST("/projects/:projectID/columns", createNewColumn)
	app.GET("/projects/:projectID/columns", getColumnsOfProject)
	app.GET("/columns/:columnID", getColumnDetails)
	app.PATCH("/columns/:columnID", updateColumn)
	app.DELETE("/columns/:columnID", deleteColumn)
	app.POST("/columns/:columnID/moves", moveColumn)
	
	// CARDS 
	app.POST("/columns/:columnID/cards", createNewCard)
	app.GET("/columns/:columnID/cards", getCardsOfColumn)
	app.GET("/cards/:cardID", getCardDetails)
	app.PATCH("/cards/:cardID", updateCard)
	app.DELETE("/cards/:cardID", deleteCard)
	app.POST("/cards/:cardID/moves", moveCard)

	//Start server
	app.Logger.Fatal(app.Start(":1010"))
}
