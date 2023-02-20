package main

import (
    "testing"

)

func TestAuthenticationDummyJira(t *testing.T) {
	email := "leon.pennington@warwick.ac.uk"
	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"
	jiraURL := "https://groupseven.atlassian.net"
	successful, err := authentication(email, token, jiraURL)

	if !successful {
        t.Fatalf(`error connecting and authenticating to dummy jira. Error: %v`, err)
    }
}

func TestAuthenticationNoEmail(t *testing.T) {
	email := ""
	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"
	jiraURL := "https://groupseven.atlassian.net"
	successful, err := authentication(email, token, jiraURL)

	if successful || err == nil{
        t.Fatalf(`authentication with empty email returns %t, %v. wanted false, error`,successful, err)
    }
}

func TestAuthenticationNoToken(t *testing.T) {
	email := "leon.pennington@warwick.ac.uk"
	token := ""
	jiraURL := "https://groupseven.atlassian.net"
	successful, err := authentication(email, token, jiraURL)

	if successful || err == nil{
        t.Fatalf(`authentication with empty token returns %t, %v. wanted false, error`,successful, err)
    }
}

func TestAuthenticationNoURL(t *testing.T) {
	email := "leon.pennington@warwick.ac.uk"
	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"
	jiraURL := ""
	successful, err := authentication(email, token, jiraURL)

	if successful || err == nil{
        t.Fatalf(`authentication with empty URL returns %t, %v. wanted false, error`,successful, err)
    }
}

// func TestAuthenticationNoURL(t *testing.T) {
// 	email := "leon.pennington@warwick.ac.uk"
// 	token := "ATATT3xFfGF0Z5Lxy5XQoNjI31MzitcRx8IZQ-idkqAnXnQJNIUSXr4eDy0UN7c-e9Ijg-8dN0fZZIFA60WHwWUcwX1uBTwtOzebHajNbc_rk46HIxkg9tzg9Zn-KYKnWx2ntx6YX2PFXhdoeiHyX7JgPcTWGFfr-5fpKWacxXyaNUWVUs0gizU=6AA42E90"
// 	jiraURL := ""
// 	successful, err := authentication(email, token, jiraURL)

// 	if successful || err == nil{
//         t.Fatalf(`authentication with empty URL returns %t, %v. wanted false, error`,successful, err)
//     }
// }