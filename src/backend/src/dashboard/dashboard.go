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

type tmp struct {
	ProjectCode string `json:"projectCode"`
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
	TeamMeanExperience     float64   `json:"teamMeanExperience"`
	WeeklyTeamMeetings     float64   `json:"weeklyTeamMeetings"`
	ClientMeetingsPerMonth float64   `json:"clientMeetingsPerMonth"`
	JiraProjectID          string    `json:"jiraProjectID"`
	JiraEmail              string    `json:"jiraUsername"`
	JiraApiToken           string    `json:jiraApiToken`
	JiraURL                string    `json:"jiraURL"`
}

// One to one mapping of the json models
// SurveyFactor is a sub-model for elements of a SurveySummary.
type SurveyFactor struct {
	Name         string  `json:"name"`
	Question     string  `json:"question"`
	Satisfaction float64 `json:"satisfaction"`
}

// SurveySummary is a sub-model of a SoftwareProject.
type SurveySummary struct {
	Date        time.Time      `json:"date"`
	Factors     []SurveyFactor `json:"factors"`
	Suggestions []string       `json:"suggestions"`
}

// HealthInformation is a sub-model of a SoftwareProject.
type HealthInformation struct {
	Message          string   `json:"message"`
	Suggestions      []string `json:"suggestions"`
	PercentageHealth float64  `json:"percentageHealth"`
}

// SoftwareSurveys is a sub-model that stores all information on the surveys undertaken by a SoftwareProject.
type SoftwareSurveys struct {
	Client *SurveySummary    `json:"client,omitempty"`
	Team   *SurveySummary    `json:"team,omitempty"`
	Health HealthInformation `json:"health"`
}

// ProjectBudget is a sub-model of a SoftwareProject.
type ProjectBudget struct {
	Budget        float64           `json:"budget"`
	Spend         float64           `json:"spend"`
	SpendOverTime []float64         `json:"spendOverTime"`
	Health        HealthInformation `json:"health"`
}

// SoftwareProject is the primary model for the API request for a project.
type SoftwareProject struct {
	Code    string            `json:"code"`
	Name    string            `json:"name"`
	Health  HealthInformation `json:"health"`
	Surveys SoftwareSurveys   `json:"surveys"`
	Budget  ProjectBudget     `json:"budget"`
}

// DashboardPage is the handler for the main dashboard page
func ProjectPage(res http.ResponseWriter, req *http.Request) {
	// Get the project code and team manager ID from the request
	var projectCode string
	var username string

	var projectName string
	var projectSuccess float64
	var budget float64
	var customSpendings float64
	var deadline time.Time
	var monthlyExpenses float64
	var teamMeanExperience float64
	var weeklyTeamMeetings float64
	var clientMeetingsPerMonth float64
	var jiraProjectCode string
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

	var tempCode tmp
	// Read the project code from the request json
	err = json.NewDecoder(req.Body).Decode(&tempCode)
	if err != nil {
		log.Println("Error reading project code: ", err)
		http.Error(res, "Error reading project code", http.StatusInternalServerError)
		return
	}
	projectCode = tempCode.ProjectCode

	db, err := connector.ConnectDB()
	if err != nil {
		log.Println("Error connecting to database: ", err)
		http.Error(res, "Error connecting to database", http.StatusInternalServerError)
		return
	}
	defer connector.CloseDB(db)

	// Get the project name from the database
	err = db.QueryRow(`SELECT projectName FROM Project WHERE projectCode = $1 AND username = $2`, projectCode, username).Scan(&projectName)
	if err != nil {
		log.Println("Error getting project name: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
		return
	}

	// // Get the project success rate from the database
	// err = db.QueryRow(`SELECT projSuccess FROM Project WHERE projectCode = $1 AND username = $2`, projectCode, username).Scan(&projectSuccess)
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
	_, err = db.Exec(`UPDATE Project SET projSuccess = $1 WHERE projectCode = $2 AND username = $3`, projectSuccess, projectCode, username)
	if err != nil {
		log.Println("Error updating project success rate: ", err)
		http.Error(res, "Error updating project success rate", http.StatusInternalServerError)
	}

	// }

	// Execute the query
	err = db.QueryRow(`
				SELECT
					projectName,
					budget,
					customSpendings,
					monthlyExpenses,
					deadline,
					teamMeanExperience,
					weeklyTeamMeetings,
					clientMeetingsPerMonth,
					jiraProjectCode,
					jiraURL
				FROM Project
				WHERE projectCode = $1 AND username = $2;`, projectCode, username).
		Scan(&projectName, &budget, &customSpendings, &monthlyExpenses, &deadline, &teamMeanExperience, &weeklyTeamMeetings, &clientMeetingsPerMonth, &jiraProjectCode, &jiraURL)
	if err != nil {
		log.Println("Error getting project data: ", err)
		http.Error(res, "No project with that project code", http.StatusInternalServerError)
		return
	}

	// combine

	currentSpend := budget - ((((deadline.Sub(time.Now())).Hours())/730.5)*monthlyExpenses + customSpendings)

	// Put all the data we have about the project in a struct
	project := ProjectData{
		ProjectCode:            projectCode,
		ProjectName:            projectName,
		ProjSuccess:            float32(projectSuccess),
		Budget:                 budget,
		MonthlyExpenses:        monthlyExpenses,
		CustomSpendings:        customSpendings,
		BudgetSpent:            currentSpend,
		Deadline:               deadline,
		TeamMeanExperience:     teamMeanExperience,
		WeeklyTeamMeetings:     weeklyTeamMeetings,
		ClientMeetingsPerMonth: clientMeetingsPerMonth,
		JiraProjectID:          jiraProjectCode,
		JiraURL:                jiraURL,
	}

	log.Println("Project data: ", project)

	softwareProject := SoftwareProject{
		Code: projectCode,
		Name: projectName,
		Health: HealthInformation{
			Message:          "",
			Suggestions:      []string{},
			PercentageHealth: 0,
		},
		Surveys: SoftwareSurveys{
			Client: &SurveySummary{
				Date:        time.Now(),
				Factors:     []SurveyFactor{},
				Suggestions: []string{},
			},
			Team: &SurveySummary{
				Date:        time.Now(),
				Factors:     []SurveyFactor{},
				Suggestions: []string{},
			},
			Health: HealthInformation{
				Message:          "",
				Suggestions:      []string{},
				PercentageHealth: 0,
			},
		},
		Budget: ProjectBudget{
			Budget:        budget,
			Spend:         0,
			SpendOverTime: []float64{},
			Health: HealthInformation{
				Message:          "",
				Suggestions:      []string{},
				PercentageHealth: 0,
			},
		},
	}

	// Return the data to the frontend as a JSON
	res.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(res).Encode(softwareProject)
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
	err = db.QueryRow(`SELECT username FROM TeamManager WHERE sessionid=$1`, sessionID).Scan(&username)
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
	var projectCode string
	var projectName string
	for rows.Next() {
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

	err = db.QueryRow(`SELECT username FROM TeamManager WHERE sessionID=$1`, sessionID).Scan(&username)
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
