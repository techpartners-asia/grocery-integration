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

	Product      *ProductService
	Category     *CategoryService
	Store        *StoreService
	Reference    *ReferenceService
	Loyalty      *LoyaltyService
	Tag          *TagService
	Branch       *BranchService
	OrderMessage *OrderMessageService

	User struct {
		Comment      *CommentService
		Coupon       *CouponService
		Imap         *ImapService
		Location     *LocationService
		Loyalty      *UserLoyaltyService
		Notification *NotificationService
		Order        *OrderService
		Profile      *ProfileService
		Reference    *UserReferenceService
		Wishlist     *WishlistService
	}

	Customer *CustomerService

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

	c.Product = &ProductService{client: c}
	c.Category = &CategoryService{client: c}
	c.Store = &StoreService{client: c}
	c.Reference = &ReferenceService{client: c}
	c.Loyalty = &LoyaltyService{client: c}
	c.Tag = &TagService{client: c}
	c.Customer = &CustomerService{client: c}
	c.Branch = &BranchService{client: c}
	c.OrderMessage = &OrderMessageService{client: c}

	c.User.Comment = &CommentService{client: c}
	c.User.Coupon = &CouponService{client: c}
	c.User.Imap = &ImapService{client: c}
	c.User.Location = &LocationService{client: c}
	c.User.Loyalty = &UserLoyaltyService{client: c}
	c.User.Notification = &NotificationService{client: c}
	c.User.Order = &OrderService{client: c}
	c.User.Profile = &ProfileService{client: c}
	c.User.Reference = &UserReferenceService{client: c}
	c.User.Wishlist = &WishlistService{client: c}

	c.SuperApp.Authenticate = &SuperAppAuthenticateService{client: c}

	return c, nil
}

func (c *Client) SetLocationID(id string) *Client {
	c.locationID = id
	c.resty.SetHeader("Location-Id", id)

	if c.Product != nil {
		c.Product.SetLocationID(id)
	}
	if c.Category != nil {
		c.Category.SetLocationID(id)
	}
	if c.Branch != nil {
		c.Branch.SetLocationID(id)
	}
	if c.User.Order != nil {
		c.User.Order.SetLocationID(id)
	}
	if c.User.Location != nil {
		c.User.Location.SetLocationID(id)
	}
	if c.User.Wishlist != nil {
		c.User.Wishlist.SetLocationID(id)
	}
	if c.User.Comment != nil {
		c.User.Comment.SetLocationID(id)
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
