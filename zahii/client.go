package zahii

import (
	"fmt"
	"strings"

	"resty.dev/v3"
)

type Client struct {
	resty         *resty.Client
	superAppToken string
	token         string

	Customer struct {
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
}

func NewClient(config Config) (*Client, error) {
	if config.BaseURL == "" {
		return nil, fmt.Errorf("BaseURL is required")
	}

	r := resty.New()

	version := config.Version
	if version == "" {
		version = V1
	}

	baseURL := config.BaseURL

	cleanBaseURL := fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), strings.TrimPrefix(string(version), "/"))
	r.SetBaseURL(cleanBaseURL)

	r.SetBasicAuth(config.Username, config.Password)
	r.SetHeader("Location-Id", config.LocationID)

	c := &Client{resty: r, superAppToken: config.SuperAppToken}

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
	c.resty.SetHeader("Location-Id", id)
	return c
}

func (c *Client) SetAuthToken(token string) *Client {
	c.token = token
	c.resty.SetAuthToken(token)
	return c
}

func (c *Client) newBaseRequest(locationID string) *resty.Request {
	r := c.resty.R()
	if locationID != "" {
		r.SetHeader("Location-Id", locationID)
	}
	return r
}

func (c *Client) newRequest(locationID string) *resty.Request {
	if c.superAppToken != "" {
		_, _ = c.SuperApp.Authenticate.AuthenticateAndSetToken(c.superAppToken)
	}

	r := c.newBaseRequest(locationID)
	if c.token != "" {
		r.SetAuthToken(c.token)
	}
	return r
}

type service struct {
	client     *Client
	locationID string
}
