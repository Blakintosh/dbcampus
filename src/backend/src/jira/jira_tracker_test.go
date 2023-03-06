package jira

import (
	"testing"
	// "net/http"
	// "os"
	// "strings"
	// "fmt"
)

func TestAuthenticationDummyJira(t *testing.T) {
	email := "leon.pennington@warwick.ac.uk"
	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"
	jiraURL := "https://groupseven.atlassian.net"
	_, err := authentication(email, token, jiraURL)

	if err != nil {
		t.Fatalf(`error connecting and authenticating to dummy jira. Error: %v`, err)
	}
}

func TestAuthenticationNoEmail(t *testing.T) {
	email := ""
	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"
	jiraURL := "https://groupseven.atlassian.net"
	_, err := authentication(email, token, jiraURL)

	if err == nil {
		t.Fatalf(`authentication with empty email returns error: %v. wanted error`, err)
	}
}

func TestAuthenticationNoToken(t *testing.T) {
	email := "leon.pennington@warwick.ac.uk"
	token := ""
	jiraURL := "https://groupseven.atlassian.net"
	_, err := authentication(email, token, jiraURL)

	if err == nil {
		t.Fatalf(`authentication with empty token returns error: %v. wanted error`, err)
	}
}

func TestAuthenticationNoURL(t *testing.T) {
	email := "leon.pennington@warwick.ac.uk"
	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"
	jiraURL := ""
	_, err := authentication(email, token, jiraURL)

	if err == nil {
		t.Fatalf(`authentication with empty URL returns error: %v. wanted error`, err)
	}
}

func TestGetProject(t *testing.T) {
	email := "leon.pennington@warwick.ac.uk"
	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"
	jiraURL := ""
	projectName := "PIT"
	client, errAuth := authentication(email, token, jiraURL)
	project, err := getProject(email, token, jiraURL, projectName, client)

	if errAuth != nil {
		t.Fatalf(`error authenticating`)
	}
	if err != nil || project == nil {
		t.Fatalf(`project not found: %v. wanted error`, err)
	}
}
