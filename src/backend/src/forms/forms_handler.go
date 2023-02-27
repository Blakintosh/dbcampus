package main

import (
	"context"
	"fmt"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/forms/v1"
)

func main() {
	// Set up authentication with Google
	ctx := context.Background()
	creds, err := google.FindDefaultCredentials(ctx, forms.FormsScope)
	if err != nil {
		fmt.Printf("Unable to get creds: %v", err)
		return
	}

	// Create a new Forms service using the OAuth2 token
	service, err := forms.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		fmt.Printf("Unable to create Forms service: %v", err)
		return
	}

	// Create a new form
	newForm := &forms.Form{
		Title:                "My Awesome Form",
		DescriptionPlainText: "This is a test form created using the Google Forms API.",
	}

	// Send the request to create the form
	response, err := service.Forms.Create(newForm).Do()
	if err != nil {
		fmt.Printf("Unable to create form: %v", err)
		return
	}

	fmt.Printf("Form created with ID %s", response.FormId)
}