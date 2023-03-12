package dashboard

import (
	auth "authentication"
	"connector"
	"database/sql"
	"encoding/json"
	"equation"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
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
	ProjectCode string `json:"code"`
	ProjectName string `json:"name"`
}

type DashboardDataList struct {
	DashboardData []DashboardData `json:"data"`
}

// Project data
type ProjectData struct {
	ProjectCode            string    `json:"projectCode"`
	ProjectName            string    `json:"projectName"`
	ProjSuccess            float32   `json:"projSuccess"`
	Budget                 float64   `json:"budget"`
	MonthlyExpenses        float64   `json:"monthlyExpenses"`
	CustomSpendings        float64   `json:"customSpendings"`
	BudgetSpent            float64   `json:"budgetSpent"`
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
func ProjectPage(res http.ResponseWriter, req *http.Request) {
	// Get the project code and team manager ID from the request
	var projectCode string
	var username string

	var projectName string
	var projectSuccess float64
	var totalBudget float64
	var customSpendings float64
	var deadline time.Time
	var monthlyExpenses float64
	var managerExperience float64
	var teamMeanExperience float64
	var weeklyTeamMeetings float64
	var clientMeetingsPerMonth float64
	var jiraProjectID string
	var jiraURL string

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

	// Execute the query
	err = db.QueryRow(`
				SELECT
					"projectName",
					"budget",
					"customSpendings",
					"monthlyExpenses",
					"deadline",
					"managerExperience",
					"teamMeanExperience",
					"weeklyTeamMeetings",
					"clientMeetingsPerMonth",
					"jiraProjectCode",
					"jiraURL"
				FROM "Project"
				WHERE "projectCode" = $1 AND "username" = $2;`, projectCode, username).
		Scan(&projectName, &totalBudget, &customSpendings, &monthlyExpenses, &deadline, &managerExperience, &teamMeanExperience, &weeklyTeamMeetings, &clientMeetingsPerMonth, &jiraProjectID, &jiraURL)
	if err != nil {
		log.Println("Error getting project data: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
	}

	// combine

	currentSpend := totalBudget - ((((deadline.Sub(time.Now())).Hours())/730.5)*monthlyExpenses + customSpendings)

	// Put all the data we have about the project in a struct
	project := ProjectData{
		ProjectCode:            projectCode,
		ProjectName:            projectName,
		ProjSuccess:            float32(projectSuccess),
		Budget:                 totalBudget,
		MonthlyExpenses:        monthlyExpenses,
		CustomSpendings:        customSpendings,
		BudgetSpent:            currentSpend,
		Deadline:               deadline,
		ManagerExperience:      managerExperience,
		TeamMeanExperience:     teamMeanExperience,
		WeeklyTeamMeetings:     weeklyTeamMeetings,
		ClientMeetingsPerMonth: clientMeetingsPerMonth,
		JiraProjectID:          jiraProjectID,
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

func DashboardPage(res http.ResponseWriter, req *http.Request) {
	var projectCodes []string
	var projectNames []string
	var projects []DashboardData = []DashboardData{}
	var projectList DashboardDataList
	var username string

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

	// Get the username from the session
	err = db.QueryRow(`SELECT "username" FROM "teammanager" WHERE sessionid=$1`, sessionID).Scan(&username)
	if err != nil {
		log.Println("Error checking session: ", err)
		http.Error(res, "Error getting username", http.StatusInternalServerError)
	}

	// Get the project codes from the database
	rows, err := db.Query(`SELECT projectCode, projectName FROM Project WHERE username = $1`, username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No projects found for user: ", username)
			// encode the empty struct to JSON
			res.Header().Set("Content-Type", "application/json")
			err = json.NewEncoder(res).Encode(projects)
			if err != nil {
				log.Println("Error encoding project data: ", err)
				http.Error(res, "Error encoding project data", http.StatusInternalServerError)
			}
			res.WriteHeader(http.StatusOK)

		}
		log.Println("Error getting project codes: ", err)
		http.Error(res, "Error getting project codes", http.StatusInternalServerError)
	}
	defer rows.Close()

	// Scan the project codes and names
	log.Println("rows.next: ", rows.Next())
	for rows.Next() {
		var projectCode string
		var projectName string
		err = rows.Scan(&projectCode, &projectName)
		if err != nil {
			log.Println("Error scanning project codes: ", err)
			http.Error(res, "Error scanning project codes", http.StatusInternalServerError)
		}
		projectCodes = append(projectCodes, projectCode)
		projectNames = append(projectNames, projectName)
	}

	// for each project code and name, add it to the struct
	for i := 0; i < len(projectCodes); i++ {
		projects = append(projects, DashboardData{
			ProjectCode: projectCodes[i],
			ProjectName: projectNames[i],
		})
		log.Println("Project code: ", projectCodes[i])
		log.Println("Project name: ", projectNames[i])
	}
	projectList.DashboardData = projects

	// Return the data to the frontend as a JSON
	res.Header().Set("Content-Type", "application/json")

	err = json.NewEncoder(res).Encode(projectList)
	if err != nil {
		log.Println("Error encoding project data: ", err)
		http.Error(res, "Error encoding project data to json", http.StatusInternalServerError)
	}
	log.Println("Successfully sent project data to frontend. " + strconv.Itoa(len(projects)) + " projects found.")
	// print json to console
	b, err := json.Marshal(projectList)
	if err != nil {
		log.Println("Error encoding project data: ", err)
		http.Error(res, "Error encoding project data", http.StatusInternalServerError)
	}
	log.Println(string(b))
	res.WriteHeader(http.StatusOK)
}

// CreateProject creates a new project in the database
func CreateProject(res http.ResponseWriter, req *http.Request) {
	// Receive the data from the frontend using json
	var project ProjectData
	var username string

	// Access the request body
	reqBody, err := ioutil.ReadAll(req.Body)
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
