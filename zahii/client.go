package zahii

import (
	"fmt"
	"log"
	"strings"

	"resty.dev/v3"
)

type Client struct {
	resty         *resty.Client
	superAppToken string
	token         string
	locationID    string

	Customer struct {
		Branch       *BranchService
		Comment      *CustomerCommentService
		Coupon       *CouponService
		Imap         *CustomerImapService
		Location     *LocationService
		Loyalty      *LoyaltyService
		Notification *NotificationService
		Order        *OrderService
		Profile      *CustomerProfileService
		Reference    *CustomerReferenceService
		Wishlist     *WishlistService
	}

	Guest struct {
		Category     *GuestCategoryService
		Customer     *GuestCustomerService
		Loyalty      *GuestLoyaltyService
		OrderMessage *GuestOrderMessageService
		Product      *GuestProductService
		Reference    *GuestReferenceService
		Store        *GuestStoreService
		Tag          *GuestTagService
	}

	SuperApp struct {
		Authenticate *SuperAppAuthenticateService
	}
}

type APIVersion string

const (
	V1 APIVersion = "v1"
)

type Config struct {
	BaseURL       string
	Version       APIVersion
	Username      string
	Password      string
	LocationID    string
	SuperAppToken string
	RestyClient   *resty.Client
	ErrorHandler  func(resp *resty.Response) error

	// Optional hook to inspect every request and response pair after completion
	RequestResponseLogger func(req *resty.Request, resp *resty.Response)
}

func NewClient(config Config) (*Client, error) {
	if config.BaseURL == "" {
		return nil, fmt.Errorf("BaseURL is required")
	}

	r := config.RestyClient
	if r == nil {
		r = resty.New()
	}

	version := config.Version
	if version == "" {
		version = V1
	}

	baseURL := config.BaseURL

	cleanBaseURL := fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), strings.TrimPrefix(string(version), "/"))
	r.SetBaseURL(cleanBaseURL)
	r.SetResponseBodyUnlimitedReads(true)
	r.SetAllowMethodDeletePayload(true)

	r.SetBasicAuth(config.Username, config.Password)
	r.SetHeader("Location-Id", config.LocationID)

	if config.ErrorHandler != nil {
		r.AddResponseMiddleware(func(c *resty.Client, resp *resty.Response) error {
			if resp.IsError() {
				return config.ErrorHandler(resp)
			}
			return nil
		})
	}

	if config.RequestResponseLogger != nil {
		r.AddResponseMiddleware(func(c *resty.Client, resp *resty.Response) error {
			config.RequestResponseLogger(resp.Request, resp)
			return nil
		})
	}

	c := &Client{resty: r, superAppToken: config.SuperAppToken, locationID: config.LocationID}

	c.Customer.Branch = &BranchService{client: c}
	c.Customer.Comment = &CustomerCommentService{client: c}
	c.Customer.Coupon = &CouponService{client: c}
	c.Customer.Imap = &CustomerImapService{client: c}
	c.Customer.Location = &LocationService{client: c}
	c.Customer.Loyalty = &LoyaltyService{client: c}
	c.Customer.Notification = &NotificationService{client: c}
	c.Customer.Order = &OrderService{client: c}
	c.Customer.Profile = &CustomerProfileService{client: c}
	c.Customer.Reference = &CustomerReferenceService{client: c}
	c.Customer.Wishlist = &WishlistService{client: c}

	c.Guest.Category = &GuestCategoryService{client: c}
	c.Guest.Customer = &GuestCustomerService{client: c}
	c.Guest.Loyalty = &GuestLoyaltyService{client: c}
	c.Guest.OrderMessage = &GuestOrderMessageService{client: c}
	c.Guest.Product = &GuestProductService{client: c}
	c.Guest.Reference = &GuestReferenceService{client: c}
	c.Guest.Store = &GuestStoreService{client: c}
	c.Guest.Tag = &GuestTagService{client: c}

	c.SuperApp.Authenticate = &SuperAppAuthenticateService{client: c}

	return c, nil
}

func (c *Client) SetLocationID(id string) *Client {
	c.locationID = id
	c.resty.SetHeader("Location-Id", id)

	// Propagate to services
	if c.Customer.Branch != nil {
		c.Customer.Branch.SetLocationID(id)
	}
	if c.Customer.Order != nil {
		c.Customer.Order.SetLocationID(id)
	}
	if c.Customer.Location != nil {
		c.Customer.Location.SetLocationID(id)
	}
	if c.Customer.Wishlist != nil {
		c.Customer.Wishlist.SetLocationID(id)
	}
	if c.Customer.Comment != nil {
		c.Customer.Comment.SetLocationID(id)
	}

	if c.Guest.Product != nil {
		c.Guest.Product.SetLocationID(id)
	}
	if c.Guest.Category != nil {
		c.Guest.Category.SetLocationID(id)
	}

	return c
}

func (c *Client) SetAuthToken(token string) *Client {
	c.token = token
	c.resty.SetAuthToken(token)
	return c
}

func (c *Client) newBaseRequest(locationID string) *resty.Request {
	r := c.resty.R()
	locID := locationID
	if locID == "" {
		locID = c.locationID
	}
	if locID != "" {
		r.SetHeader("Location-Id", locID)
		// log.Printf("[SDK DEBUG] Set Location-Id header to %s", locID)
	}
	return r
}

func (c *Client) newRequest(locationID string) *resty.Request {
	if c.superAppToken != "" && c.token == "" {
		_, _ = c.SuperApp.Authenticate.AuthenticateAndSetToken(c.superAppToken)
	}

	r := c.newBaseRequest(locationID)
	if c.token != "" {
		r.SetAuthToken(c.token)
	}
	log.Printf("[SDK DEBUG] Request Headers for %s: %v", r.URL, r.Header)
	return r
}

type service struct {
	client     *Client
	locationID string
}
