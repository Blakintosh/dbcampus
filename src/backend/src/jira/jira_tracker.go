package main

import (
	"context"
	"log"
	"os"

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

	// Authentication data, user has to input both email and API token
	tr := jira.BasicAuthTransport{
		Username: "<email>",     // User's Jira email. Don't include < and >.
		APIToken: "<API token>", // User's Jira API token. Don't include < and >.
	}

	// Creating Jira client
	client, err := jira.NewClient(jiraURL, tr.Client())
	if err != nil {
		log.Fatalf("Error creating Jira client: %v\n", err)

	}

	// Getting current user as a test
	user, _, err := client.User.GetCurrentUser(context.Background())
	if err != nil {
		log.Printf("Error getting current user: %v\n", err)
		log.Fatalln("Make sure you have inputted the correct email, API token and jira team url.")
	}

	log.Printf("Accessed project. Logged as: %v", user.EmailAddress)

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

}
