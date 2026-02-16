package main

import (
	"fmt"
	"log"

	"github.com/techpartners-asia/grocery-integration/zahii"
)

func main() {
	// Initialize the client
	client, err := zahii.NewClient(zahii.Config{
		BaseURL:  "https://api.zahii.mn/api",
		Username: "super-app",
		Password: "password-here",
	})
	if err != nil {
		log.Fatalf("Error initializing client: %v", err)
	}

	// Example 1: List categories using Guest.Category service
	categories, err := client.Guest.Category.List(zahii.ListCategoryRequest{Active: true})
	if err != nil {
		log.Fatalf("Error listing categories: %v", err)
	}
	fmt.Printf("Found %d categories\n", len(categories.Body))

	// Example 2: Create a comment using Customer.Comment service
	// (Note: ListComment was changed to Create in this refactor to match Postman structure)
	commentResp, err := client.Customer.Comment.Create(zahii.CreateCommentRequest{
		Body: "Great product!",
		Rate: 5,
	})
	if err != nil {
		log.Fatalf("Error creating comment: %v", err)
	}
	fmt.Printf("Comment creation status: %v\n", commentResp.Message)

	fmt.Println("Zahii SDK initialization successful!")
}
