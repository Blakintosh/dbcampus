package main

import (
	"context"
	"log"
	"os"
	"time"
	"errors"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {

	/*************** Setting up looger and authentication *********************/
	// Setting up logger
	logFile, err := os.OpenFile("log.log", os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	// Jira team url
	jiraURL := "https://groupseven.atlassian.net" // User's Jira URL
	email := "leon.pennington@warwick.ac.uk"
	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"

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
		log.Fatalln("Make sure you have inputted the correct email, API token and jira team url.")
	}
	log.Printf("Accessed project. Logged as: %v", user.EmailAddress)

	// Getting number of overdue tasks
	numOverDueTasks, err := getNumOverDueIssuesFromProject(email, token, jiraURL, "PIT", client)
	if err != nil {
		log.Fatalf("Error getting number of overdue tasks: %v\n", err)
	}
	log.Printf("Number of overdue tasks: %v", numOverDueTasks)

	// Getting priority of overdue tasks
	priorityOfTasks, err := getPriorities(email, token, jiraURL, "PIT", client)
	if err != nil {
		log.Fatalf("Error getting priority of overdue tasks: %v\n", err)
	}
	log.Printf("Priority of overdue tasks: %v", priorityOfTasks)

}

/************************ Authentication set up *******************************/
func authentication(email string, token string, url string) (bool, error) {

	// Jira team url
	jiraURL := url // User's Jira URL
	
	if jiraURL == "" {
		return false, errors.New("empty URL")
	}
	if token == "" {
		return false, errors.New("empty token")
	}
	if email == "" {
		return false, errors.New("empty email")
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
		return false, errors.New("Error creating Jira client")

	}

	// Getting current user to check if authentication was successful
	user, _, err := client.User.GetCurrentUser(context.Background())
	if err != nil {
		// log.Printf("Error getting current user: %v\n", err)
		// log.Fatalln("Make sure you have inputted the correct email, API token and jira team url.")
		return false, errors.New("Error getting current user")
	}

	log.Printf("Accessed project. Logged as: %v", user.EmailAddress)
	return true, nil
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
func getNumOverDueIssuesFromProject(email string, token string, url string, projectName string, client *jira.Client) (int, error) {
	counter := 0

	// Getting all issues
	issues, _, err := client.Issue.Search(context.Background(), "project = "+projectName, nil)
	if err != nil {
		log.Fatalf("Error getting issues: %v\n", err)
	}

	// Get the deadline of each issue
	for _, issue := range issues {
		// Get the deadline
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
			return -1, err
		}
		deadlineTime := time.Time(ti)

		// Check if the deadline is overdue deadline is of type time.Time
		if status == "In Progress" && deadlineTime.Before(time.Now()) {
			log.Printf("Issue %v is overdue", issue.Key)
			counter++
		}

	}
	return counter, nil
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

/******************** Getting all projects example ************************/

// Get all projects
// projects, _, err := client.Project.Get(context.Background(), "PIT")
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
