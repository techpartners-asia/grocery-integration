package zahii

import (
	"time"
)

type Base struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Branch struct {
	Base
	Name              string   `json:"name"`
	Active            bool     `json:"active"`
	AddressDesc       string   `json:"address_desc"`
	BagHoroo          string   `json:"bag_horoo"`
	Contact           string   `json:"contact"`
	Description       string   `json:"description"`
	ImagePaths        []string `json:"image_paths"`
	IsOpen            bool     `json:"is_open"`
	Latitude          float64  `json:"latitude"`
	Longitude         float64  `json:"longitude"`
	ScheduleStart     string   `json:"schedule_start"`
	ScheduleEnd       string   `json:"schedule_end"`
	SumDuureg         string   `json:"sum_duureg"`
	VeritechBranchID  string   `json:"veritech_branch_id"`
	CurrentOrderCount int      `json:"current_order_count"`
}

type Staff struct {
	Base
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	LastLoginDate time.Time `json:"last_login_date"`
	Active        bool      `json:"active"`
	BranchID      uint      `json:"branch_id"`
	Branch        *Branch   `json:"branch,omitempty"`
	RoleID        uint      `json:"role_id"`
}

type Customer struct {
	Base
	FirstName     string     `json:"first_name"`
	LastName      string     `json:"last_name"`
	Email         string     `json:"email"`
	Phone         string     `json:"phone"`
	BirthDate     *time.Time `json:"birth_date"`
	Gender        string     `json:"gender"`
	Register      string     `json:"register"`
	AuthUID       string     `json:"auth_uid"`
	LastLoginDate time.Time  `json:"last_login_date"`
	LastOrderID   uint       `json:"last_order_id"`
	Active        bool       `json:"active"`
	Point         int64      `json:"point"`
	PointBalance  int64      `json:"point_balance"`
	LevelID       uint       `json:"level_id"`
	IntegrationID string     `json:"integration_id"`
}

type Category struct {
	Base
	Name         string          `json:"name"`
	IsSet        bool            `json:"is_set"`
	Active       bool            `json:"active"`
	Rank         int             `json:"rank"`
	Color        string          `json:"color"`
	IconImage    string          `json:"icon_image"`
	ProductCount int             `json:"product_count"`
	Locales      []*CategoryLang `json:"locales,omitempty"`
}

type CategoryLang struct {
	ReferralID uint   `json:"referral_id"`
	Lang       string `json:"lang"`
	Name       string `json:"name"`
}

type Product struct {
	Base
	Name                   string           `json:"name"`
	Ingredient             string           `json:"ingredient"`
	Garnish                string           `json:"garnish"`
	Description            string           `json:"description"`
	Calories               float64          `json:"calories"`
	Sugar                  float64          `json:"sugar"`
	Fat                    float64          `json:"fat"`
	Protein                float64          `json:"protein"`
	ImagePaths             []string         `json:"image_paths"`
	DefaultPrice           float64          `json:"default_price"`
	CategoryID             uint             `json:"category_id"`
	Category               *Category        `json:"category,omitempty"`
	Active                 bool             `json:"active"`
	IsNewlyAdded           bool             `json:"is_newly_added"`
	TodaysSuggest          bool             `json:"todays_suggest"`
	IsRaw                  bool             `json:"is_raw"`
	IsNotSinglePurchasable bool             `json:"is_not_single_purchasable"`
	Popular                bool             `json:"popular"`
	Rank                   int              `json:"rank"`
	Rate                   float64          `json:"rate"`
	RateCount              int              `json:"rate_count"`
	Quantity               int              `json:"quantity"`
	IsPackage              bool             `json:"is_package"`
	StoreID                uint             `json:"store_id"`
	KeyWords               []string         `json:"key_words"`
	Price                  float64          `json:"price"`
	Branches               []*ProductBranch `json:"branches,omitempty"`
}

type ProductBranch struct {
	Base
	ProductID uint    `json:"product_id"`
	BranchID  uint    `json:"branch_id"`
	Branch    *Branch `json:"branch,omitempty"`
	Price     float64 `json:"price"`
	Qty       int     `json:"qty"`
	UnitPrice float64 `json:"unit_price"`
	IsActive  bool    `json:"is_active"`
}

type Store struct {
	Base
	Name      string `json:"name"`
	ImagePath string `json:"image_path"`
}

type OrderStatus string

const (
	OrderStatusNew              OrderStatus = "new"
	OrderStatusReceived         OrderStatus = "recieved"
	OrderStatusPreparing        OrderStatus = "preparing"
	OrderStatusPrepared         OrderStatus = "prepared"
	OrderStatusOutForDelivery   OrderStatus = "out_for_delivery"
	OrderStatusDelivered        OrderStatus = "delivered"
	OrderStatusCancelled        OrderStatus = "cancelled"
	OrderStatusCancelledDone    OrderStatus = "cancelled_done"
	OrderStatusWaitingForDriver OrderStatus = "waiting_for_driver"
)

type Order struct {
	Base
	OrderUID              string       `json:"order_uid"`
	OrderNumber           string       `json:"order_number"`
	BranchID              uint         `json:"branch_id"`
	Branch                *Branch      `json:"branch,omitempty"`
	CustomerID            uint         `json:"customer_id"`
	Customer              *Customer    `json:"customer,omitempty"`
	Status                OrderStatus  `json:"status"`
	Items                 []*OrderItem `json:"items,omitempty"`
	TotalPrice            float64      `json:"total_price"`
	SubtotalPrice         float64      `json:"subtotal_price"`
	DiscountPrice         float64      `json:"discount_price"`
	DeliverPrice          float64      `json:"deliver_price"`
	PaymentType           string       `json:"payment_type"`
	IsPaid                bool         `json:"is_paid"`
	Phone                 string       `json:"phone"`
	Description           string       `json:"description"`
	ScheduledDeliveryDate *string      `json:"scheduled_delivery_date"`
}

type OrderItem struct {
	Base
	OrderID      uint     `json:"order_id"`
	ProductID    uint     `json:"product_id"`
	Product      *Product `json:"product,omitempty"`
	ProductSetID uint     `json:"product_set_id"`
	PriceTotal   float64  `json:"price_total"`
	Qty          int      `json:"qty"`
	UnitPrice    float64  `json:"unit_price"`
	Note         string   `json:"note"`
}

type BaseResponse struct {
	Message string `json:"message"`
}

type ProductSet struct {
	Base
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	ImagePath     string            `json:"image_path"`
	Active        bool              `json:"active"`
	IsNewlyAdded  bool              `json:"is_newly_added"`
	TodaysSuggest bool              `json:"todays_suggest"`
	IsRaw         bool              `json:"is_raw"`
	TotalPrice    float64           `json:"total_price"`
	Rank          int               `json:"rank"`
	BranchID      uint              `json:"branch_id"`
	Items         []*ProductSetItem `json:"items,omitempty"`
}

type ProductSetItem struct {
	ID           uint       `json:"id"`
	ProductSetID uint       `json:"product_set_id"`
	Products     []*Product `json:"products,omitempty"`
	Qty          int64      `json:"qty"`
	UnitPrice    float64    `json:"unit_price"`
	HaveChoose   bool       `json:"have_choose"`
}

type LoyaltyLevel struct {
	Base
	Name            string  `json:"name"`
	Type            string  `json:"type"`
	MinPoint        int64   `json:"min_point"`
	MaxPoint        int64   `json:"max_point"`
	DiscountPercent float64 `json:"discount_percent"`
	IconPath        string  `json:"icon_path"`
}

type AppInfo struct {
	Base
	Version     string `json:"version"`
	Platform    string `json:"platform"`
	IsForce     bool   `json:"is_force"`
	Description string `json:"description"`
}

type JobApplication struct {
	Base
	FirstName     string   `json:"first_name"`
	LastName      string   `json:"last_name"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Register      string   `json:"register"`
	Gender        string   `json:"gender"`
	District      string   `json:"district"`
	Khoroo        string   `json:"khoroo"`
	Entrance      string   `json:"entrance"`
	DoorNumber    string   `json:"door_number"`
	MaritalStatus string   `json:"marital_status"`
	IsDriver      bool     `json:"is_driver"`
	LicenseType   string   `json:"license_type"`
	Education     string   `json:"education"`
	Attachment    []string `json:"attachment"`
	Position      string   `json:"position"`
	CustomerID    uint     `json:"customer_id"`
}

type Feedback struct {
	Base
	CategoryID   uint      `json:"category_id"`
	Category     *Category `json:"category,omitempty"`
	ServicedDate time.Time `json:"serviced_date"`
	Service      float64   `json:"service"`
	Note         string    `json:"note"`
	CustomerID   uint      `json:"customer_id"`
}

type Banner struct {
	Base
	Title          string   `json:"title"`
	Description    string   `json:"description"`
	HyperLink      string   `json:"hyper_link"`
	Active         bool     `json:"active"`
	HorizontalPath string   `json:"horizontal_path"`
	Type           string   `json:"type"`
	ProductID      uint     `json:"product_id"`
	Product        *Product `json:"product,omitempty"`
}

type StartUpSlider struct {
	Base
	HyperLink string `json:"hyper_link"`
	Active    bool   `json:"active"`
	ImagePath string `json:"image_path"`
}

type Invitation struct {
	Base
	Description string `json:"description"`
	Active      bool   `json:"active"`
	Point       int64  `json:"point"`
}

type InvitationCode struct {
	Base
	Code        string `json:"code"`
	Description string `json:"description"`
	Point       int64  `json:"point"`
	Limit       int    `json:"limit"`
	Active      bool   `json:"active"`
	UsedCount   int    `json:"used_count"`
}

type Address struct {
	Base
	Name     string `json:"name"`
	ParentID uint   `json:"parent_id"`
	Type     string `json:"type"`
}

type Lang struct {
	Code string `json:"code"`
	Name string `json:"name"`
}
type Wishlist struct {
	Base
	Name       string          `json:"name"`
	CustomerID uint            `json:"customer_id"`
	Items      []*WishlistItem `json:"items,omitempty"`
}

type JobApplicationRequest struct {
	Email         string   `json:"email"`
	FirstName     string   `json:"first_name"`
	LastName      string   `json:"last_name"`
	Phone         string   `json:"phone"`
	Position      string   `json:"position"`
	Register      string   `json:"register"`
	Attachment    []string `json:"attachment,omitempty"`
	District      string   `json:"district,omitempty"`
	DoorNumber    string   `json:"door_number,omitempty"`
	Education     string   `json:"education,omitempty"`
	Entrance      string   `json:"entrance,omitempty"`
	Gender        string   `json:"gender,omitempty"`
	IsDriver      bool     `json:"is_driver"`
	Khoroo        string   `json:"khoroo,omitempty"`
	LicenseType   string   `json:"license_type,omitempty"`
	MaritalStatus string   `json:"marital_status,omitempty"`
}

type WishlistItem struct {
	Base
	WishlistID uint     `json:"wishlist_id"`
	ProductID  uint     `json:"product_id"`
	Product    *Product `json:"product,omitempty"`
}

type Coupon struct {
	Base
	Code      string  `json:"code"`
	Amount    float64 `json:"amount"`
	EndDate   string  `json:"end_date"`
	IsActive  bool    `json:"is_active"`
	UsedCount int     `json:"used_count"`
}

type CustomerCoupon struct {
	Base
	CouponID   uint      `json:"coupon_id"`
	Coupon     *Coupon   `json:"coupon,omitempty"`
	CustomerID uint      `json:"customer_id"`
	Customer   *Customer `json:"customer,omitempty"`
	IsUsed     bool      `json:"is_used"`
	UsedAt     *string   `json:"used_at"`
	OrderUID   string    `json:"order_uid"`
}

type LoyaltyPromotion struct {
	Base
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Percent     float64 `json:"percent"`
}

type ListInvitationResponse struct {
	BaseResponse
	Body []*Invitation `json:"body"`
}

type ListLoyaltyPromotionResponse struct {
	BaseResponse
	Body []*LoyaltyPromotion `json:"body"`
}

type AuthenticateResponse struct {
	BaseResponse
	Body struct {
		Token string `json:"token"`
	} `json:"body"`
}

type CustomerResponse struct {
	BaseResponse
	Body *Customer `json:"body"`
}
