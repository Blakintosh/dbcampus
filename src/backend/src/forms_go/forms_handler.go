package main

import (
	"context"
	"fmt"
	// "github.com/gookit/config/v2"
	"google.golang.org/api/option"
	"golang.org/x/oauth2"
	"google.golang.org/api/forms/v1"
)

func main() {
	config := &oauth2.Config{
		ClientID:     "your_google_client_id",
		ClientSecret: "your_google_client_secret",
	}
	// Set up authentication with Google
	// ctx := context.Background()
	// creds, err := google.FindDefaultCredentials(ctx, forms.FormsBodyScope)
	// if err != nil {
	// 	fmt.Printf("Unable to get creds: %v", err)
	// 	return
	// }
	ctx := context.Background()
	code := r.FormValue("code")
	token, err := config.Exchange(ctx, "code")
	formsService, err := forms.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	// Create a new Forms service using the OAuth2 token
	// ctx := context.Background()
	// service, err := forms.NewService(ctx)
	// service, err := forms.NewService(ctx, option.WithCredentials(creds))
	if err != nil {
		fmt.Printf("Unable to create Forms service: %v", err)
		return
	}

	// Create a new form
	newFormInfo := &forms.Info{
		Title:                "My Awesome Form",
		Description: "This is a test form created using the Google Forms API.",
	}
	newForm := &forms.Form{
		Info:    newFormInfo,
	}

	// Send the request to create the form
	response, err := formsService.Forms.Create(newForm).Do()
	if err != nil {
		fmt.Printf("Unable to create form: %v", err)
		return
	}

	fmt.Printf("Form created with ID %s", response.FormId)
}