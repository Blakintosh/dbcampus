package main

import (
	"context"
	"fmt"
	"system"

	jira "github.com/andygrunwald/go-jira/v2/cloud"
)

func main() {
	jiraURL := "https://go-jira-opensource.atlassian.net/"

	// Jira docs: https://support.atlassian.com/atlassian-account/docs/manage-api-tokens-for-your-atlassian-account/
	// Create a new API token: https://id.atlassian.com/manage-profile/security/api-tokens
	// tr := &http.Transport{
	// TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	// }
	// client1 := &http.Client{Transport: tr}
	tp := jira.BasicAuthTransport{
		Username: "Karim.Zeyads@warwick.ac.uk",
		APIToken: "ATATT3xFfGF0niFoT5pmrKVYdKFoYQ5Li4rYAubiFuv8nIPlben8R336h2PLu7Px37xfPHEXp2LaBxncoh8AACpDmVsV_ETtemD5zlhgF5uVfYEFPyJO-PnxMFvbEMnn66p7-uM-1FMUKPSp5Ev7a-f-0COVzuDnqfFokCFOj__rX3QLRvNjsXw=27D9CA4E",
	}
	client, err := jira.NewClient(jiraURL, tp.Client())
	if err != nil {
		fmt.Printf("Error at making a new client: %e ", err)
		system.Exit(-1)
	}

	u, _, err := client.User.GetCurrentUser(context.Background())
	if err != nil {
		fmt.Printf("Error at getting user: %e", err)
		system.Exit(-1)
	}

	fmt.Printf("Email: %v\n", u.EmailAddress)
	fmt.Println("Success!")
}
