package main

import (
	"context"
	"fmt"
	"os"

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
		Username: "<username>",
		APIToken: "<api_token>",
	}
	client, err := jira.NewClient(jiraURL, tp.Client())
	if err != nil {
		fmt.Printf("Error at making a new client: %e ", err)
		os.Exit(-1)

	}

	u, _, err := client.User.GetCurrentUser(context.Background())
	if err != nil {
		fmt.Printf("Error at getting user: %e", err)
		os.Exit(-1)
	}

	fmt.Printf("Email: %v\n", u.EmailAddress)
	fmt.Println("Success!")
}
