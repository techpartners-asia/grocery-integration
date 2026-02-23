package main

import (
	"fmt"
	"log"
	"time"

	"github.com/techpartners-asia/grocery-integration/zahii"
	"resty.dev/v3"
)

type TestContext struct {
	CategoryID uint
	ProductID  uint
	LocationID uint
	BranchID   uint
	OrderUID   string
}

func wait() {
	time.Sleep(3000 * time.Millisecond)
}

func main() {
	log.SetFlags(log.Ltime)
	ctx := &TestContext{}

	client, err := zahii.NewClient(zahii.Config{
		BaseURL:       "http://localhost:8000/api",
		Username:      "super-app",
		Password:      "rakbir-fihzo7-wevrIg",
		SuperAppToken: "0de7c10c9971bb0cc230765012ea672e",

		RequestResponseLogger: func(req *resty.Request, resp *resty.Response) {
			if resp.IsError() {
				log.Printf("ERR | %s %s -> %d | %s", req.Method, req.URL, resp.StatusCode(), resp.String())
			} else {
				log.Printf("OK  | %s %s -> %d", req.Method, req.URL, resp.StatusCode())
			}
		},
	})

	if err != nil {
		log.Fatalf("Failed to initialize client: %v", err)
	}

	// Explicitly authenticate to show the token for Postman testing
	log.Println("--- AUTHENTICATION ---")
	authResp, err := client.SuperApp.Authenticate.AuthenticateAndSetToken("0de7c10c9971bb0cc230765012ea672e")
	if err != nil {
		log.Printf("!!! Authentication failed: %v", err)
	} else {
		log.Printf(">>> AUTH TOKEN (for Postman): %s", authResp.Body.Token)
	}

	log.Println("\n--- STARTING REALISTIC TEST SCENARIOS ---")

	// 1. GUEST DISCOVERY SCENARIO
	if err := runGuestDiscovery(client, ctx); err != nil {
		log.Printf("[SCENARIO] Guest Discovery aborted: %v", err)
	}

	// 2. CUSTOMER PROFILE SCENARIO
	if err := runCustomerProfile(client, ctx); err != nil {
		log.Printf("[SCENARIO] Customer Profile aborted: %v", err)
	}

	// 3. WISHLIST MANAGEMENT SCENARIO
	if err := runWishlistScenario(client, ctx); err != nil {
		log.Printf("[SCENARIO] Wishlist Scenario aborted: %v", err)
	}

	// 4. REFERENCE & LOCATION SCENARIO
	if err := runLocationScenario(client, ctx); err != nil {
		log.Printf("[SCENARIO] Location Scenario aborted: %v", err)
	}

	// 5. ORDER CHECK & HISTORY SCENARIO
	if err := runOrderScenario(client, ctx); err != nil {
		log.Printf("[SCENARIO] Order Scenario aborted: %v", err)
	}

	// 6. COMMENT SCENARIO
	if err := runCommentScenario(client, ctx); err != nil {
		log.Printf("[SCENARIO] Comment Scenario aborted: %v", err)
	}

	log.Println("\n--- TESTING COMPLETE ---")
}

func runGuestDiscovery(c *zahii.Client, ctx *TestContext) error {
	log.Println("\n[SCENARIO] 🛍️  GUEST DISCOVERY")

	// 1. List Categories
	log.Println("   - Listing active categories...")
	catResp, err := c.Guest.Category.List(zahii.ListCategoryRequest{Active: true})
	if err != nil {
		return err
	}
	if len(catResp.Body) == 0 {
		return fmt.Errorf("no active categories found")
	}
	ctx.CategoryID = catResp.Body[0].ID
	log.Printf("   ✓ Found %d categories. Selected '%s' (ID: %d)", len(catResp.Body), catResp.Body[0].Name, ctx.CategoryID)
	wait()

	// 2. List Products in selected Category
	log.Printf("   - Listing products in Category %d...", ctx.CategoryID)
	prodResp, err := c.Guest.Product.List(zahii.ListProductRequest{CategoryID: ctx.CategoryID, Active: true})
	if err != nil {
		return err
	}
	if len(prodResp.Body) == 0 {
		log.Println("   ! No products in this category, trying global list...")
		prodResp, _ = c.Guest.Product.List(zahii.ListProductRequest{Active: true})
	}

	if len(prodResp.Body) > 0 {
		ctx.ProductID = prodResp.Body[0].ID
		log.Printf("   ✓ Selected product '%s' (ID: %d)", prodResp.Body[0].Name, ctx.ProductID)
	} else {
		return fmt.Errorf("no products available to test")
	}
	wait()

	// 3. Get Product Details
	log.Printf("   - Getting details for Product %d...", ctx.ProductID)
	_, err = c.Guest.Product.Get(ctx.ProductID, "")
	if err != nil {
		return err
	}
	wait()

	return nil
}

func runCustomerProfile(c *zahii.Client, ctx *TestContext) error {
	log.Println("\n[SCENARIO] 👤 CUSTOMER PROFILE")

	// 1. Get Profile
	log.Println("   - Fetching current profile...")
	_, err := c.Customer.Profile.GetProfile(zahii.InfoRequestDTO{})
	if err != nil {
		return err
	}
	wait()

	// 2. Point History
	log.Println("   - Fetching point history...")
	_, err = c.Customer.Profile.GetPointHistory(zahii.PointHistoryRequest{Limit: 10, Page: 1})
	if err != nil {
		return err
	}
	wait()

	return nil
}

func runWishlistScenario(c *zahii.Client, ctx *TestContext) error {
	log.Println("\n[SCENARIO] ❤️  WISHLIST MANAGEMENT")
	if ctx.ProductID == 0 {
		return fmt.Errorf("no product ID available from discovery")
	}

	// 1. Create Wishlist
	log.Printf("   - Creating wishlist 'TEST RUN' with product %d...", ctx.ProductID)
	_, err := c.Customer.Wishlist.Create(zahii.CreateWishlistRequest{
		Name:      "TEST RUN",
		ProductID: ctx.ProductID,
	})
	if err != nil {
		return err
	}
	wait()

	// 2. List Wishlists
	log.Println("   - Verifying creation in list...")
	list, err := c.Customer.Wishlist.List()
	if err != nil {
		return err
	}
	var myWishlistID uint
	for _, w := range list.Body {
		if w.Name == "TEST RUN" {
			myWishlistID = w.ID
			break
		}
	}

	if myWishlistID == 0 {
		log.Println("   ! Created wishlist not found in list (might be delayed), skipping delete")
	} else {
		log.Printf("   ✓ Found wishlist ID: %d. Deleting...", myWishlistID)
		wait()
		c.Customer.Wishlist.Delete(myWishlistID)
	}
	wait()

	return nil
}

func runLocationScenario(c *zahii.Client, ctx *TestContext) error {
	log.Println("\n[SCENARIO] 📍 REFERENCE & LOCATION")

	// 1. Polygon Check
	log.Println("   - Checking polygon for Ulaanbaatar center...")
	c.Guest.Reference.CheckPolygon(zahii.PolygonCheckRequest{
		Latitude:  47.9221,
		Longitude: 106.9155,
	})
	wait()

	// 2. List Branches
	log.Println("   - Listing available branches...")
	branches, err := c.Customer.Branch.ListBranch()
	if err != nil {
		return err
	}
	if len(branches.Body) > 0 {
		ctx.BranchID = branches.Body[0].ID
		log.Printf("   ✓ Selected branch '%s' (ID: %d)", branches.Body[0].Name, ctx.BranchID)
	}
	wait()

	// 3. List Customer Locations
	log.Println("   - Listing customer locations...")
	locs, err := c.Customer.Location.List()
	if err != nil {
		return err
	}
	if len(locs.Body) > 0 {
		ctx.LocationID = locs.Body[0].ID
		log.Printf("   ✓ Selected existing location '%s' (ID: %d)", locs.Body[0].Name, ctx.LocationID)
	}
	wait()

	return nil
}

func runOrderScenario(c *zahii.Client, ctx *TestContext) error {
	log.Println("\n[SCENARIO] 📦 ORDER FLOW")

	// 1. History
	log.Println("   - Fetching order history...")
	history, err := c.Customer.Order.GetOrderHistory(zahii.OrderHistoryRequest{Limit: 5, Page: 1})
	if err != nil {
		return err
	}
	if len(history.Body) > 0 {
		ctx.OrderUID = history.Body[0].OrderUID
		log.Printf("   ✓ Found previous order (UID: %s)", ctx.OrderUID)
	}
	wait()

	// 2. Check Order (Simulation)
	if ctx.BranchID != 0 && ctx.ProductID != 0 && ctx.LocationID != 0 {
		log.Printf("   - Simulating order check for Product %d at Branch %d...", ctx.ProductID, ctx.BranchID)
		c.Customer.Order.CheckOrder(zahii.CreateOrderRequest{
			BranchID:           ctx.BranchID,
			CustomerLocationID: ctx.LocationID,
			Phone:              "99112233",
			Type:               "delivery",
			Items: []zahii.OrderCreateItem{
				{ProductID: ctx.ProductID, Qty: 1},
			},
		})
	}
	wait()

	return nil
}

func runCommentScenario(c *zahii.Client, ctx *TestContext) error {
	log.Println("\n[SCENARIO] 💬 COMMENTING")
	if ctx.ProductID == 0 {
		return fmt.Errorf("no product ID available")
	}

	// 1. Create Comment
	log.Printf("   - Posting comment on Product %d...", ctx.ProductID)
	resp, err := c.Customer.Comment.Create(zahii.CreateCommentRequest{
		ProductID: ctx.ProductID,
		Body:      "Excellent quality and fast service!",
		Rate:      5,
		Title:     "Highly Recommended",
	})
	if err != nil {
		return err
	}
	wait()

	// 2. Delete it (Cleanup)
	if resp.Body.ID != 0 {
		log.Printf("   - Deleting test comment ID %d...", resp.Body.ID)
		c.Customer.Comment.Delete(zahii.DeleteCommentRequest{ID: resp.Body.ID})
	}
	wait()

	return nil
}
