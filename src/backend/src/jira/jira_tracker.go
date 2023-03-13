package jira

import (
	"context"
	"errors"
	"log"
	"time"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

type ProjectData struct {
	ProjectName string `json:"projectName"` // Name of the project
	ProjectID   string `json:"projectID"`   // ID of the project
	NumTasks    int    `json:"numTasks"`    // Number of tasks in the project
	NumOverdue  int    `json:"numOverdue"`  // Number of overdue tasks in the project
}

/************************ Authentication set up *******************************/
func authentication(email string, token string, url string) (*jira.Client, error) {

	// Jira team url
	jiraURL := url // User's Jira URL

	if jiraURL == "" {
		return nil, errors.New("empty URL")
	}
	if token == "" {
		return nil, errors.New("empty token")
	}
	if email == "" {
		return nil, errors.New("empty email")
	}

	// Authentication data, user has to input both email and API token
	tr := jira.BasicAuthTransport{
		Username: email, // User's Jira email
		APIToken: token,
	}

	// Creating Jira client
	client, err := jira.NewClient(jiraURL, tr.Client())
	if err != nil {
		// log.Fatalf("Error creating Jira client: %v\n", err)
		return nil, errors.New("error creating Jira client")

	}

	// Getting current user to check if authentication was successful
	user, _, err := client.User.GetCurrentUser(context.Background())
	if err != nil {
		// log.Printf("Error getting current user: %v\n", err)
		// log.Fatalln("Make sure you have inputted the correct email, API token and jira team url.")
		return nil, errors.New("error getting current user")
	}

	log.Printf("Accessed project. Logged as: %v", user.EmailAddress)
	return client, nil
}

/******************* Gets project data through its name ***********************/
func getProject(email string, token string, url string, projectName string, client *jira.Client) (*jira.Project, error) {
	// Authentication
	authentication(email, token, url)

	// Getting all projects
	project, _, err := client.Project.Get(context.Background(), projectName)
	if err != nil {
		log.Fatalf("Error getting projects: %v\n", err)
	}

	return project, err
}

/********************* Getting number of overdue tasks *************************/
func GetNumOverDueIssuesFromProject(projectName string, jiraURL string, email string, token string) (int, int, error) {
	// Jira team url
	// jiraURL := "https://groupseven.atlassian.net" // User's Jira URL
	// email := "Karim.Zeyada@warwick.ac.uk"
	// token := "ATATT3xFfGF0HNtow0fIs24CsTvCYbEG5RkrnO9UaayuQCfn_K797qIKQ8TRtJitAayzDld3JZHuB88ujP_cTFQctzuWHS-luFE9A48EjMJWa5TLiXjvzXEuynPTCtLGH5eweIvwwQvxCbCGZoIcJ2f0FvHPzn_dLDdUpZbwFUPIFdXlGfWYxQs=B6C3BA30"

	// Authentication data, user has to input both email and API token
	tr := jira.BasicAuthTransport{
		Username: email, // User's Jira email
		APIToken: token,
	}

	// Creating Jira client
	client, err := jira.NewClient(jiraURL, tr.Client())
	if err != nil {
		log.Fatalf("Error creating Jira client: %v\n", err)

	}

	// Getting current user to check if authentication was successful
	user, _, err := client.User.GetCurrentUser(context.Background())
	if err != nil {
		log.Printf("Error getting current user: %v\n", err)
		log.Println("Make sure you have inputted the correct email, API token and jira team url.")
	}
	if err == nil {
		log.Printf("Accessed project. Logged as: %v", user.EmailAddress)
		counter := 0
		tasks := 0

		// Getting all issues
		issues, _, err := client.Issue.Search(context.Background(), "project = "+projectName, nil)
		if err != nil {
			log.Printf("Error getting issues: %v\n", err)
		}

		// Get the deadline of each issue
		for _, issue := range issues {
			// Get the deadline
			tasks++
			deadline := issue.Fields.Duedate
			status := issue.Fields.Status.Name
			deadlineByte, _ := deadline.MarshalJSON()
			deadlineByte = deadlineByte[1 : len(deadlineByte)-1]

			if string(deadlineByte) == "null" {
				log.Printf("Issue %v has no deadline", issue.Key)
				continue
			}
			ti, err := time.Parse("2006-01-02", string(deadlineByte))
			if err != nil {
				log.Printf("Error parsing time: %v", err)
				return -1, -1, err
			}
			deadlineTime := time.Time(ti)

			// Check if the deadline is overdue deadline is of type time.Time
			if status == "In Progress" && deadlineTime.Before(time.Now()) {
				log.Printf("Issue %v is overdue", issue.Key)
				counter++
			}

		}
		return counter, tasks, nil
	}

	return -1, -1, err
}

/**************** Get priorties of all issues in a project ********************/
func getPriorities(email string, token string, url string, projectName string, client *jira.Client) ([]string, error) {
	// Getting all issues
	issues, _, err := client.Issue.Search(context.Background(), "project = "+projectName, nil)
	if err != nil {
		log.Fatalf("Error getting issues: %v\n", err)
	}

	// Create a slice to store all priorities
	priorities := make([]string, 0)

	// Get the priority of each issue
	for _, issue := range issues {
		// Get the priority
		priority := issue.Fields.Priority.Name
		log.Printf("Priority: %v", priority)
		priorities = append(priorities, priority)

	}

	return priorities, nil
}

// func makeProjectData(projectName string, projectID string) *ProjectData {
// 	numTasks, numOverdue, err := getNumOverDueIssuesFromProject(projectName)
// 	if err != nil {
// 		log.Fatalf("Error getting number of overdue tasks: %v", err)
// 	}
// 	projectData := &ProjectData{
// 		ProjectName: projectName,
// 		ProjectID:   projectID,
// 		NumTasks:    numTasks,
// 		NumOverdue:  numOverdue,
// 	}

// 	return projectData
// }

// func MakeProjectData(res http.ResponseWriter, req *http.Request) {
// 	// Get the project name from the request
// 	var projectData *ProjectData
// 	if req.Method == "POST" {
// 		// Access the request body
// 		reqBody, err := ioutil.ReadAll(req.Body)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 		err = json.Unmarshal([]byte(reqBody), &projectData)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	}

// 	projectData = makeProjectData(projectData.ProjectName, projectData.ProjectID)

// 	// Convert the project data to json
// 	jsonData, err := json.Marshal(projectData)
// 	if err != nil {
// 		log.Fatalf("Error converting project data to json: %v", err)
// 	}
// 	// Send the project data to the frontend
// 	res.Header().Set("Content-Type", "application/json")
// 	res.Write(jsonData)
// }

/******************** Getting all projects example ************************/

// Get all projects
// Jira team url
// var jiraURL = "https://groupseven.atlassian.net" // User's Jira URL
// var email = "Karim.Zeyada@warwick.ac.uk"
// var token = "ATATT3xFfGF0HNtow0fIs24CsTvCYbEG5RkrnO9UaayuQCfn_K797qIKQ8TRtJitAayzDld3JZHuB88ujP_cTFQctzuWHS-luFE9A48EjMJWa5TLiXjvzXEuynPTCtLGH5eweIvwwQvxCbCGZoIcJ2f0FvHPzn_dLDdUpZbwFUPIFdXlGfWYxQs=B6C3BA30"

// // Authentication data, user has to input both email and API token
// var tr = jira.BasicAuthTransport{
// 	Username: email, // User's Jira email
// 	APIToken: token,
// }

// // Creating Jira client
// var client, err = jira.NewClient(jiraURL, tr.Client())
// if err != nil {
// 	log.Fatalf("Error creating Jira client: %v\n", err)

// }
// var projects, _, err = client.Project.Get(context.Background(), "PIT")
// if err != nil {
// 	log.Fatalf("Error getting projects: %v\n", err)
// }
// print all project data
// fmt.Println("Project name: ", projects.Name)
// fmt.Println("Project key: ", projects.Key)
// fmt.Println("Project ID: ", projects.ID)
// fmt.Println("Project URL: ", projects.Self)
// fmt.Println("Project description: ", projects.Description)
// fmt.Println("Project lead: ", projects.Lead.Name)
// fmt.Println("Project lead email: ", projects.Lead.EmailAddress)
// fmt.Println("Project lead URL: ", projects.Lead.Self)
// fmt.Println("Project issues types: ", projects.IssueTypes)
// fmt.Println("Project projectCategory: ", projects.ProjectCategory)
// fmt.Println("Project url: ", projects.URL)
// fmt.Println("Project roles: ", projects.Roles)

/***************** Getting all issues in project PIT **********************/
// Get all issues
// issues, _, err := client.Issue.Search(context.Background(), "project = PIT", nil)
// if err != nil {
// 	panic(err)
// }

// // print all issues data
// for _, issue := range issues {
// 	fmt.Println("Issue name: ", issue.Fields.Project.Name)
// 	fmt.Println("Issue key: ", issue.Key)
// 	fmt.Println("Issue ID: ", issue.ID)
// 	fmt.Println("Issue URL: ", issue.Self)
// 	fmt.Println("Issue summary: ", issue.Fields.Summary)
// 	fmt.Println("Issue description: ", issue.Fields.Description)
// 	fmt.Println("Issue status: ", issue.Fields.Status.Name)
// 	fmt.Println("Issue status URL: ", issue.Fields.Status.Self)
// 	fmt.Println("Issue status description: ", issue.Fields.Status.Description)
// 	fmt.Println("Issue status icon URL: ", issue.Fields.Status.IconURL)
// 	fmt.Println("Issue status statusCategory: ", issue.Fields.Status.StatusCategory.Name)
// 	fmt.Println("Issue status statusCategory URL: ", issue.Fields.Status.StatusCategory.Self)
// 	fmt.Println("Issue status statusCategory key: ", issue.Fields.Status.StatusCategory.Key)
// }
