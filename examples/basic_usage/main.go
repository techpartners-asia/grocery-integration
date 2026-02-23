package main

import (
	"fmt"
	"log"

	"github.com/techpartners-asia/grocery-integration/zahii"
	"resty.dev/v3"
)

func main() {
	// Initialize the client
	client, err := zahii.NewClient(zahii.Config{
		BaseURL:  "https://api.example.com",
		Username: "super-app",
		Password: "password-here",

		// Optional: catch every request/response explicitly
		RequestResponseLogger: func(req *resty.Request, resp *resty.Response) {
			log.Printf("[ZAHII-SDK] %s %s -> HTTP %d (Took %v)",
				req.Method,
				req.URL,
				resp.StatusCode(),
				resp.Duration(),
			)
			log.Printf("Request Body: %v", req.Body)
			log.Printf("Response Body: %s", resp.String())
		},
	})
	if err != nil {
		log.Fatalf("Error initializing client: %v", err)
	}

	// Example 1: List categories (public, no auth required)
	categories, err := client.Category.List(zahii.ListCategoryRequest{Active: true})
	if err != nil {
		log.Fatalf("Error listing categories: %v", err)
	}
	fmt.Printf("Found %d categories\n", len(categories.Body))

	// Example 2: Create a comment (auth required)
	commentResp, err := client.User.Comment.Create(zahii.CreateCommentRequest{
		Body: "Great product!",
		Rate: 5,
	})
	if err != nil {
		log.Fatalf("Error creating comment: %v", err)
	}
	fmt.Printf("Comment creation status: %v\n", commentResp.Message)

	fmt.Println("Zahii SDK initialization successful!")
}
