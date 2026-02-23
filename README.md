# Zahii Grocery Integration SDK

A Go client library for accessing the Zahii Grocery API.

## Installation

```bash
go get github.com/techpartners-asia/grocery-integration
```

## Overview

This SDK provides a convenient, strongly-typed wrapper around the Zahii Grocery API. It uses `resty.dev/v3` for underlying HTTP networking and exposes three main service sets:
- **Guest**: Services accessible indiscriminately or with basic read scopes (e.g., categories, stores, products).
- **Customer**: Authenticated services used for customer actions (e.g., orders, wishlist, profile).
- **SuperApp**: Services related to SuperApp authentication flows.

## Usage

### Initialization

To use the SDK, import the `zahii` package and initialize a client with your configuration:

```go
package main

import (
	"fmt"
	"log"

	"github.com/techpartners-asia/grocery-integration/zahii"
)

func main() {
	client, err := zahii.NewClient(zahii.Config{
		BaseURL:  "https://api.zahii.mn/api",
		Username: "super-app",
		Password: "password-here",
		// LocationID: "optional-location-id",
		// Version: zahii.V1, // defaults to "v1"
		
		// Optional: Global HTTP Error Handler Hook
		// ErrorHandler: func(resp *resty.Response) error {
		//   return fmt.Errorf("API Error: Status %d - %s", resp.StatusCode(), resp.String())
		// },
	})
	if err != nil {
		log.Fatalf("Error initializing client: %v", err)
	}

	// Make API calls...
}
```

### Authentication & Headers
You can configure the client inline to update headers post-initialization:
```go
// Switch location dynamically context
client.SetLocationID("store-1234")

// Setup token for a Customer authenticated session
client.SetAuthToken("eyJhbGciOi...")
```

### Global Error Handling
You can optionally intercept and normalize `resty.Response` errors using the `ErrorHandler` in the SDK configuration. This is automatically invoked whenever a non-2xx status code is returned. The error you return from this hook will be passed directly back as the `err` result from any SDK method call.

```go
type APIError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

// 1. Configure the Client
client, _ := zahii.NewClient(zahii.Config{
    // ...
    ErrorHandler: func(resp *resty.Response) error {
        var rawErr APIError
        if err := json.Unmarshal(resp.Body(), &rawErr); err == nil && rawErr.Message != "" {
            return fmt.Errorf("API failed with %d: %s", resp.StatusCode(), rawErr.Message)
        }
        return fmt.Errorf("request failed with status %d: %s", resp.StatusCode(), resp.String())
    },
})

// 2. The error from the handler directly propagates to SDK caller methods
_, err := client.Customer.Profile.GetProfile(zahii.InfoRequestDTO{})
if err != nil {
    // If the server answered 401 Unauthorized with {"message": "Invalid token"}, 
    // this prints: "API failed with 401: Invalid token"
    log.Println(err.Error())
}
```

### Examples

#### Calling a Guest Service (Listing Categories)
```go
categories, err := client.Guest.Category.List(zahii.ListCategoryRequest{Active: true})
if err != nil {
    log.Fatalf("Error: %v", err)
}
fmt.Printf("Found %d categories\n", len(categories.Body))
```

#### Calling a Customer Service (Creating a Comment)
```go
commentResp, err := client.Customer.Comment.Create(zahii.CreateCommentRequest{
    Body: "Great product!",
    Rate: 5,
})
if err != nil {
    log.Fatalf("Error: %v", err)
}
fmt.Println(commentResp.Message)
```

For more details, see `examples/basic_usage/main.go`.

## Capabilities

### `Guest` Services
- `Category`: List and manage categories.
- `Customer`: Handle customer registration/status context.
- `Loyalty`: Access general loyalty policies and tiers.
- `OrderMessage`: Messaging utilities for orders.
- `Product`: List, search, and view single products.
- `Reference`: Configuration and common reference data.
- `Store`: Location and active branches data.
- `Tag`: Product metadata tags and labels.

### `Customer` Services
- `Comment`: Product and shopping/driver reviews.
- `Coupon`: View and apply coupons.
- `Imap`: Map-related interactions.
- `Location`: Managing customer addresses.
- `Loyalty`: Earned points and history.
- `Notification`: In-app notification feeds.
- `Order`: Creation, modification, and history of grocery deliveries.
- `Profile`: Manage account details.
- `Reference`: Secure reference sets.
- `Wishlist`: Saved items.

### `SuperApp` Services
- `Authenticate`: Super-app authentication.

## License
Refer to the `LICENSE` file for details.
