package zahii

import (
	"time"
)

type Base struct {
	ID        uint       `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
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
	Type              string   `json:"type,omitempty"`
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
	Provider      string        `json:"provider"`
	FirstName     string        `json:"first_name"`
	LastName      string        `json:"last_name"`
	Email         string        `json:"email"`
	Phone         string        `json:"phone"`
	BirthDate     *time.Time    `json:"birth_date"`
	Gender        string        `json:"gender"`
	Register      string        `json:"register"`
	AuthUID       string        `json:"auth_uid"`
	LastLoginDate time.Time     `json:"last_login_date"`
	PushToken     string        `json:"-"`
	LastOrderID   uint          `json:"last_order_id"`
	Active        bool          `json:"active"`
	Point         int64         `json:"point"`
	PointBalance  int64         `json:"point_balance"`
	LevelID       uint          `json:"level_id"`
	Level         *LoyaltyLevel `json:"level,omitempty"`
	IntegrationID string        `json:"integration_id"`
	PhoneOrder    string        `json:"phone_order"`
}

type Category struct {
	Base
	Name          string            `json:"name"`
	IsSet         bool              `json:"is_set"`
	Active        bool              `json:"active"`
	Rank          int               `json:"rank"`
	Color         string            `json:"color"`
	IconImage     string            `json:"icon_image"`
	VerticalImage string            `json:"vertical_image"`
	ProductCount  int               `json:"product_count"`
	Locales       []*CategoryLang   `json:"locales,omitempty"`
	Promo         *LoyaltyPromotion `json:"promo,omitempty"`
}

type CategoryLang struct {
	ReferralID uint   `json:"referral_id"`
	Lang       string `json:"lang"`
	Name       string `json:"name"`
}

type Product struct {
	Base
	Name                   string            `json:"name"`
	Ingredient             string            `json:"ingredient"`
	Garnish                string            `json:"garnish"`
	Description            string            `json:"description"`
	Calories               float64           `json:"calories"`
	Sugar                  float64           `json:"sugar"`
	Fat                    float64           `json:"fat"`
	Protein                float64           `json:"protein"`
	ImagePaths             []string          `json:"image_paths"`
	DefaultPrice           float64           `json:"default_price"`
	CategoryID             uint              `json:"category_id"`
	Category               *Category         `json:"category,omitempty"`
	PromoProduct           *interface{}      `json:"promo_product,omitempty"` // simplified
	Active                 bool              `json:"active"`
	IsNewlyAdded           bool              `json:"is_newly_added"`
	TodaysSuggest          bool              `json:"todays_suggest"`
	IsRaw                  bool              `json:"is_raw"`
	IsNotSinglePurchasable bool              `json:"is_not_single_purchasable"`
	Popular                bool              `json:"popular"`
	Rank                   int               `json:"rank"`
	Rate                   float64           `json:"rate"`
	RateCount              int               `json:"rate_count"`
	Quantity               int               `json:"quantity"`
	IsPackage              bool              `json:"is_package"`
	StoreID                uint              `json:"store_id"`
	KeyWords               []string          `json:"key_words"`
	Price                  float64           `json:"price"`
	AvailableBranches      []string          `json:"available_branches,omitempty"`
	Tags                   *[]interface{}    `json:"tags,omitempty"`
	IsRestricted           bool              `json:"is_restricted"`
	ProductPromotion       *LoyaltyPromotion `json:"product_promotion,omitempty"`
	Branches               []*ProductBranch  `json:"branches,omitempty"`
}

type ProductBranch struct {
	Base
	ProductID    uint    `json:"product_id"`
	BranchID     uint    `json:"branch_id"`
	Branch       *Branch `json:"branch,omitempty"`
	Price        float64 `json:"price"`
	Qty          int     `json:"qty"`
	UnitPrice    float64 `json:"unit_price"`
	IsActive     bool    `json:"is_active"`
	TempQuantity int     `json:"temp_quantity,omitempty"`
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
	OrderUID              string         `json:"order_uid"`
	BranchID              uint           `json:"branch_id"`
	Branch                *Branch        `json:"branch,omitempty"`
	CustomerID            uint           `json:"customer_id"`
	Customer              *Customer      `json:"customer,omitempty"`
	StaffID               uint           `json:"staff_id,omitempty"`
	Staff                 *Staff         `json:"staff,omitempty"`
	CookStaffID           uint           `json:"cook_staff_id,omitempty"`
	CookStaff             *Staff         `json:"cook_staff,omitempty"`
	DriverID              uint           `json:"driver_id,omitempty"`
	Driver                *Staff         `json:"driver,omitempty"`
	Phone                 string         `json:"phone"`
	Description           string         `json:"description"`
	PaymentType           string         `json:"payment_type"`
	PanchanCount          int            `json:"panchan_count"`
	OrderNumber           string         `json:"order_number"`
	Type                  string         `json:"type"`
	DeliverTimeID         uint           `json:"deliver_time_id"`
	Status                OrderStatus    `json:"status"`
	Items                 []*OrderItem   `json:"items,omitempty"`
	IsPaid                bool           `json:"is_paid"`
	TotalPrice            float64        `json:"total_price"`
	PriceWithoutCoupon    int64          `json:"price_without_coupon"`
	OnlyDiscount          float64        `json:"only_discount"`
	Deliver               *interface{}   `json:"deliver,omitempty"` // simplified
	SubtotalPrice         float64        `json:"subtotal_price"`
	DiscountPrice         float64        `json:"discount_price"`
	SpendPoint            int64          `json:"spend_point"`
	GivenPoint            int64          `json:"given_point"`
	LevelDiscountPercent  float64        `json:"level_discount_percent"`
	LevelDiscount         float64        `json:"level_discount"`
	CouponDiscount        float64        `json:"coupon_discount"`
	DeliverPrice          float64        `json:"deliver_price"`
	CustomerCouponID      uint           `json:"customer_coupon_id"`
	CustomerCouponCode    string         `json:"customer_coupon_code"`
	Payments              []*interface{} `json:"payments,omitempty"` // simplified
	DeliveringTime        *float64       `json:"delivering_time"`
	UsedLoyalty           bool           `json:"used_loyalty"`
	CancelTaskID          string         `json:"cancel_task_id"`
	ScheduledDeliveryDate *string        `json:"scheduled_delivery_date"`
	IsVeritechSynced      bool           `json:"is_veritech_synced"`
	IsSuperApp            bool           `json:"is_super_app"`
	ExternalCouponAmount  float64        `json:"external_coupon_amount"`
}

type OrderItem struct {
	Base
	OrderID                uint        `json:"order_id"`
	Order                  *Order      `json:"order,omitempty"`
	ProductID              uint        `json:"product_id"`
	Product                *Product    `json:"product,omitempty"`
	ProductSetID           uint        `json:"product_set_id"`
	ProductSet             *ProductSet `json:"product_set,omitempty"`
	PriceTotal             float64     `json:"price_total"`
	Qty                    int         `json:"qty"`
	PriceSize              string      `json:"price_size"`
	Note                   string      `json:"note"`
	PromoID                uint        `json:"promo_id"`
	Discountpercent        float64     `json:"discount_percent"`
	AfterDiscountTotal     int64       `json:"after_discount_total"`
	IsHistory              bool        `json:"is_history"`
	SwitchedItemID         uint        `json:"switched_item_id"`
	PackageID              uint        `json:"package_id"`
	IsPromotion            bool        `json:"is_promotion"`
	AfterDiscountUnitPrice float64     `json:"after_discount_unit_price"`
	UnitPrice              float64     `json:"unit_price"`
	UnitVat                float64     `json:"unit_vat"`
	StoreID                uint        `json:"store_id"`
	Store                  *Store      `json:"store,omitempty"`
	IsManualAdded          bool        `json:"is_manual_added"`
	IsReturned             bool        `json:"is_returned"`
	ReturnedQty            int         `json:"returned_qty"`
	ReturnedStaffID        uint        `json:"returned_staff_id"`
	ManualAddedStaffID     uint        `json:"manual_added_staff_id"`
}

type BaseResponse struct {
	Message string `json:"message"`
}

type ProductSet struct {
	Base
	Name          string            `json:"name"`
	Description   string            `json:"description"`
	ImagePath     string            `json:"image_path"`
	Items         []*ProductSetItem `json:"items,omitempty"`
	Locales       []*interface{}    `json:"locales,omitempty"`
	Active        bool              `json:"active"`
	IsNewlyAdded  bool              `json:"is_newly_added"`
	TodaysSuggest bool              `json:"todays_suggest"`
	IsRaw         bool              `json:"is_raw"`
	TotalPrice    float64           `json:"total_price"`
	Rank          int               `json:"rank"`
	Is21          bool              `json:"is_21"`
	Is18          bool              `json:"is_18"`
	Rate          float64           `json:"rate"`
	BranchID      uint              `json:"branch_id"`
	Branch        *Branch           `json:"branch,omitempty"`
	ParentID      uint              `json:"parent_id"`
	Parent        *ProductSet       `json:"parent,omitempty"`
	Branches      []*ProductSet     `json:"branches,omitempty"`
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
	FirstName     string    `json:"first_name"`
	LastName      string    `json:"last_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Register      string    `json:"register"`
	Gender        string    `json:"gender"`
	District      string    `json:"district"`
	Khoroo        string    `json:"khoroo"`
	Entrance      string    `json:"entrance"`
	DoorNumber    string    `json:"door_number"`
	MaritalStatus string    `json:"marital_status"`
	IsDriver      bool      `json:"is_driver"`
	LicenseType   string    `json:"license_type"`
	Education     string    `json:"education"`
	Attachment    []string  `json:"attachment,omitempty"`
	Position      string    `json:"position"`
	CustomerID    uint      `json:"customer_id"`
	Customer      *Customer `json:"customer,omitempty"`
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
	Description         string         `json:"description"`
	Active              bool           `json:"active"`
	Point               int64          `json:"point"`
	CustomerInvitations []*interface{} `json:"customer_invitations,omitempty"`
	MaxInvitation       int            `json:"max_invitation"`
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
	Name         string  `json:"name"`
	ParentID     uint    `json:"parent_id"`
	Type         string  `json:"type"`
	RegionName   string  `json:"region_name"`
	RegionTypeID int     `json:"region_type_id"`
	OrderField   int     `json:"order_field"`
	DeliverPrice float64 `json:"deliver_price"`
	BranchID     uint    `json:"branch_id"`
	Branch       *Branch `json:"branch,omitempty"`
	Active       bool    `json:"active"`
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
	WishlistID   uint        `json:"wishlist_id"`
	ProductID    uint        `json:"product_id"`
	Product      *Product    `json:"product,omitempty"`
	ProductSetID uint        `json:"product_set_id"`
	ProductSet   *ProductSet `json:"product_set,omitempty"`
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
	Title           string     `json:"title"`
	StartDate       *string    `json:"start_date,omitempty"`
	EndDate         *string    `json:"end_date,omitempty"`
	TargetProductID uint       `json:"target_product_id"`
	TargetProduct   *Product   `json:"target_product,omitempty"`
	IsActive        bool       `json:"is_active"`
	BuyQty          int        `json:"buy_qty"`
	FreeQty         int        `json:"free_qty"`
	Options         []*Product `json:"options,omitempty"`
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
		User  struct {
			Email     string `json:"email"`
			Phone     string `json:"phone"`
			EbarimtNo string `json:"ebarimt_no"`
			ID        string `json:"id"`
		} `json:"user"`
	} `json:"body"`
}

type CustomerResponse struct {
	BaseResponse
	Body *Customer `json:"body"`
}
