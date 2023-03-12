package dashboard

import (
	auth "authentication"
	"connector"
	"encoding/json"
	"equation"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

/**************************** Entities/Main Objects ***************************/
// Manager data
type Manager struct {
	TeammanagerID            int    `json:"teammanagerid"`
	Username                 string `json:"username"`
	Password                 string `json:"password"`
	SessionID                string `json:"sessionid"`
	ManagerSuccessPercentage string `json:"managersuccesspercentage"`
}

// Dashboard data
type DashboardData struct {
	ProjectCode            string    `json:"projectCode"`
	Username               string    `json:"username"`
	ProjectName            string    `json:"projectName"`
	ProjectCodes           []string  `json:"projectCodes"`
	ProjSuccess            float32   `json:"projSuccess"`
	Budget                 float64   `json:"budget"`
	MonthlyExpenses        float64   `json:"monthlyExpenses"`
	CustomSpendings        float64   `json:"customSpendings"`
	BudgetSpent            float64   `json:"budgetSpent"`
	Deadline               time.Time `json:"deadline"`
	TeamMeanExperience     float64   `json:"teamMeanExperience"`
	WeeklyTeamMeetings     float64   `json:"weeklyTeamMeetings"`
	ClientMeetingsPerMonth float64   `json:"clientMeetingsPerMonth"`
	JiraURL                string    `json:"jiraURL"`
}

// Dashboard data
type ProjectData struct {
	ProjectCode            string    `json:"projectCode"`
	ProjectName            string    `json:"projectName"`
	Budget                 float64   `json:"budget"`
	MonthlyExpenses        float64   `json:"monthlyExpenses"`
	CustomSpendings        float64   `json:"customSpendings"`
	Deadline               time.Time `json:"deadline"`
	ManagerExperience      float64   `json:"managerExperience"`
	TeamMeanExperience     float64   `json:"teamMeanExperience"`
	WeeklyTeamMeetings     float64   `json:"weeklyTeamMeetings"`
	ClientMeetingsPerMonth float64   `json:"clientMeetingsPerMonth"`
	JiraProjectID          string    `json:"jiraProjectID"`
	JiraEmail              string    `json:"jiraUsername"`
	JiraApiToken           string    `json:"jiraApiToken"`
	JiraURL                string    `json:"jiraURL"`
}

// DashboardPage is the handler for the main dashboard page
func DashboardPage(res http.ResponseWriter, req *http.Request) {
	// Get the project code and team manager ID from the request
	var projectCode string
	var username string

	var projectName string
	var projectSuccess float64
	var totalBudget float64
	var customSpendings float64
	var deadline time.Time
	var monthlyExpenses float64
	var teamMeanExperience float64
	var weeklyTeamMeetings float64
	var clientMeetingsPerMonth float64
	var jiraURL string

	var projectCodes []string

	// Get the username from the session cookie
	session, err := auth.Store.Get(req, "session")
	if err != nil {
		log.Println("Error getting session. User might not be logged in: ", err)
		res.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Check if the user is logged in
	if session.Values["username"] == nil || session.Values["authenticated"] != true {
		log.Println("User is not logged in")
		http.Error(res, "User is not logged in", http.StatusUnauthorized)
		return
	}

	// Get the username from the session cookie
	username = session.Values["username"].(string)

	// Read the project code from the request json
	err = json.NewDecoder(req.Body).Decode(&projectCode)
	if err != nil {
		log.Println("Error reading project code: ", err)
		http.Error(res, "Error reading project code", http.StatusInternalServerError)
	}

	db, err := connector.ConnectDB()
	if err != nil {
		log.Println("Error connecting to database: ", err)
		http.Error(res, "Error connecting to database", http.StatusInternalServerError)
	}
	defer connector.CloseDB(db)

	// Get the projects that has the same team manager
	rows, err := db.Query(`SELECT "projectCode" FROM "Project" WHERE "username" = $1`, username)
	if err != nil {
		log.Println("Error getting projects: ", err)
		http.Error(res, "Error getting projects", http.StatusInternalServerError)
	}
	defer rows.Close()

	// Put the project codes in an array
	for rows.Next() {
		err := rows.Scan(&projectCodes)
		if err != nil {
			log.Println("Error scanning project codes: ", err)
			http.Error(res, "Error scanning project codes", http.StatusInternalServerError)
		}
	}

	// Return the project codes as a json
	if err != nil {
		log.Println("Error encoding project codes: ", err)
		http.Error(res, "Error encoding project codes", http.StatusInternalServerError)
	}

	// Get the project name from the database
	err = db.QueryRow(`SELECT "projectName" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&projectName)
	if err != nil {
		log.Println("Error getting project name: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// // Get the project success rate from the database
	// err = db.QueryRow(`SELECT "projSuccess" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&projectSuccess)
	// if err != nil {
	// 	log.Println("Error getting project success rate: ", err)
	// 	http.Error(res, "No project with that project code", http.StatusInternalServerError)
	// }
	// if projectSuccess == 0 {

	// Decided to claculate the project success rate every time the user opens the dashboard
	log.Println("Project success rate wasn't calculated before")
	projectSuccess, err = equation.GetPercentage(username, projectCode)
	if err != nil {
		log.Println("Error calculating project success rate: ", err)
		http.Error(res, "Error calculating project success rate", http.StatusInternalServerError)
	}
	// Update the project success rate in the database
	_, err = db.Exec(`UPDATE "Project" SET "projSuccess" = $1 WHERE "projectCode" = $2 AND "username" = $3`, projectSuccess, projectCode, username)
	if err != nil {
		log.Println("Error updating project success rate: ", err)
		http.Error(res, "Error updating project success rate", http.StatusInternalServerError)
	}

	// }

	// Get the total budget from the database
	err = db.QueryRow(`SELECT "budget" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&totalBudget)
	if err != nil {
		log.Println("Error getting total budget: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// Get the additional spendings from the database
	err = db.QueryRow(`SELECT "customSpendings" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&customSpendings)
	if err != nil {
		log.Println("Error getting current budget: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// Get the monthly expenses from the database
	err = db.QueryRow(`SELECT "monthlyExpenses" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&monthlyExpenses)
	if err != nil {
		log.Println("Error getting monthly expenses: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// Get the deadline from the database
	err = db.QueryRow(`SELECT "deadline" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&deadline)
	if err != nil {
		log.Println("Error getting deadline: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// Get the team mean experience from the database
	err = db.QueryRow(`SELECT "teamMeanExperience" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&teamMeanExperience)
	if err != nil {
		log.Println("Error getting team mean experience: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// Get the weekly team meetings from the database
	err = db.QueryRow(`SELECT "weeklyTeamMeetings" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&weeklyTeamMeetings)
	if err != nil {
		log.Println("Error getting weekly team meetings: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// Get the client meetings per month from the database
	err = db.QueryRow(`SELECT "clientMeetingsPerMonth" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&clientMeetingsPerMonth)
	if err != nil {
		log.Println("Error getting client meetings per month: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// Get the Jira URL from the database
	err = db.QueryRow(`SELECT "jiraURL" FROM "Project" WHERE "projectCode" = $1 AND "username" = $2`, projectCode, username).Scan(&jiraURL)
	if err != nil {
		log.Println("Error getting Jira URL: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	currentSpend := totalBudget - ((((deadline.Sub(time.Now())).Hours())/730.5)*monthlyExpenses + customSpendings)

	// Put all the data we have about the project in a struct
	project := DashboardData{
		ProjectCode:            projectCode,
		Username:               username,
		ProjectName:            projectName,
		ProjectCodes:           projectCodes,
		ProjSuccess:            float32(projectSuccess),
		Budget:                 totalBudget,
		MonthlyExpenses:        monthlyExpenses,
		CustomSpendings:        customSpendings,
		BudgetSpent:            currentSpend,
		Deadline:               deadline,
		TeamMeanExperience:     teamMeanExperience,
		WeeklyTeamMeetings:     weeklyTeamMeetings,
		ClientMeetingsPerMonth: clientMeetingsPerMonth,
		JiraURL:                jiraURL,
	}

	// Return the data to the frontend as a JSON
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(project)
	if err != nil {
		log.Println("Error encoding project data: ", err)
		http.Error(res, "Error encoding project data", http.StatusInternalServerError)
	}
	res.WriteHeader(http.StatusOK)

}

// CreateProject creates a new project in the database
func CreateProject(res http.ResponseWriter, req *http.Request) {
	// Receive the data from the frontend using json
	var project ProjectData
	var username string

	// Access the request body
	reqBody, err := ioutil.ReadAll(req.Body)
	fmt.Println("reqBody register: ", string(reqBody))
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to read input.", 500)
	}
	err = json.Unmarshal([]byte(reqBody), &project)
	if err != nil {
		log.Println(err)
		http.Error(res, "Server error, unable to read input.", 500)
	}

	// Connect to the database
	db, err := connector.ConnectDB()
	if err != nil {
		log.Println("Error connecting to the database: ", err)
		http.Error(res, "Error connecting to the database", http.StatusInternalServerError)
	}
	defer connector.CloseDB(db)

	// Get the username from the session
	session, err := auth.Store.Get(req, "session")

	if err != nil || session.Values["authenticated"] == nil {
		log.Println("Unable to get session: ", err)
		http.Error(res, "Unable to get session", http.StatusInternalServerError)
	}

	sessionID := session.Values["sessionId"].(string)
	log.Println("Session ID: ", sessionID)

	err = db.QueryRow(`SELECT "username" FROM "teammanager" WHERE sessionid=$1`, sessionID).Scan(&username)
	if err != nil {
		log.Println("Error checking session: ", err)
	}

	// Put the project data in the database
	_, err = db.Exec(`INSERT INTO project (projectCode, username, projectName, budget, monthlyExpenses, customSpendings, deadline, teamMeanExperience, weeklyTeamMeetings, clientMeetingsPerMonth,jiraProjectCode, jiraURL) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`, project.ProjectCode, username, project.ProjectName, project.Budget, project.MonthlyExpenses, project.CustomSpendings, project.Deadline, project.TeamMeanExperience, project.WeeklyTeamMeetings, project.ClientMeetingsPerMonth, project.JiraProjectID, project.JiraURL)
	if err != nil {
		log.Println("Error inserting project data: ", err)
		http.Error(res, "Error inserting project data", http.StatusInternalServerError)
	}

	res.WriteHeader(http.StatusOK)
}
