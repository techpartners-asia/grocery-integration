package zahii

type OrderService service

func (s *OrderService) SetLocationID(id string) *OrderService {
	s.locationID = id
	return s
}

type CreateOrderRequest struct {
	BranchID              uint              `json:"branch_id"`
	CustomerCouponID      uint              `json:"customer_coupon_id,omitempty"`
	CustomerLocationID    uint              `json:"customer_location_id"`
	DeliverTimeID         uint              `json:"deliver_time_id"`
	ExternalPhone         string            `json:"external_phone,omitempty"`
	Items                 []OrderCreateItem `json:"items"`
	LocationID            string            `json:"location_id,omitempty"`
	Phone                 string            `json:"phone"`
	ScheduledDeliveryDate string            `json:"scheduled_delivery_date,omitempty"`
	Type                  string            `json:"type"`
	UsePoint              bool              `json:"use_point"`
}

type OrderCreateItem struct {
	ProductID               uint                    `json:"product_id"`
	Qty                     int                     `json:"qty"`
	Note                    string                  `json:"note,omitempty"`
	ProductSetID            uint                    `json:"product_set_id,omitempty"`
	ProductPromotionOptions []OrderPromotionOption  `json:"product_promotion_options,omitempty"`
	ProductSetChosen        []OrderProductSetChosen `json:"product_set_choosen,omitempty"`
}

type OrderPromotionOption struct {
	ProductID uint `json:"product_id"`
	Qty       int  `json:"qty"`
}

type OrderProductSetChosen struct {
	ProductID uint `json:"product_id"`
	SetItemID uint `json:"set_item_id"`
}

type CreateOrderResponse struct {
	BaseResponse
	Body *Order `json:"body"`
}

func (s *OrderService) CreateOrder(req CreateOrderRequest) (*CreateOrderResponse, error) {
	var result CreateOrderResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/order/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetOrderResponse struct {
	BaseResponse
	Body *Order `json:"body"`
}

func (s *OrderService) GetOrder(uid string) (*GetOrderResponse, error) {
	var result GetOrderResponse
	_, err := s.client.newRequest(s.locationID).
		SetPathParam("uid", uid).
		SetResult(&result).
		Get("/customer/order/get/{uid}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListOrderResponse struct {
	BaseResponse
	Body []*Order `json:"body"`
}

func (s *OrderService) ListActive() (*ListOrderResponse, error) {
	var result ListOrderResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/customer/order/active/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *OrderService) CheckOrder(req CreateOrderRequest) (*CreateOrderResponse, error) {
	var result CreateOrderResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/order/check")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type DeliverTime struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	Active    bool   `json:"active"`
}

type ListDeliverTimeResponse struct {
	BaseResponse
	Body []*DeliverTime `json:"body"`
}

func (s *OrderService) ListDeliverTime() (*ListDeliverTimeResponse, error) {
	var result ListDeliverTimeResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/order/deliver_time/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type EbarimtResponse struct {
	BaseResponse
	Body string `json:"body"`
}

func (s *OrderService) GetEbarimt(orderUID string) (*EbarimtResponse, error) {
	var result EbarimtResponse
	_, err := s.client.newRequest(s.locationID).
		SetPathParam("uid", orderUID).
		SetResult(&result).
		Get("/customer/order/ebarimt/{uid}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type OrderHistoryRequest struct {
	CreatedAt  []string          `json:"created_at,omitempty"`
	CustomerID uint              `json:"customer_id,omitempty"`
	Limit      int               `json:"limit,omitempty"`
	Page       int               `json:"page,omitempty"`
	Search     string            `json:"search,omitempty"`
	Sorter     map[string]string `json:"sorter,omitempty"`
}

func (s *OrderService) GetOrderHistory(req OrderHistoryRequest) (*ListOrderResponse, error) {
	var result ListOrderResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/order/history/cursor")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type OrderPaymentRequest struct {
	OrderUID    string `json:"order_uid"`
	PaymentType string `json:"payment_type"`
}

func (s *OrderService) CreateOrderPayment(req OrderPaymentRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/order/payment")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type OrderSuggestRequest struct {
	Items []OrderCreateItem `json:"items"`
}

func (s *OrderService) GetOrderSuggestions(req OrderSuggestRequest) (*ListProductResponse, error) {
	var result ListProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/order/suggest/items")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *OrderService) CancelOrder(orderUID string) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(map[string]string{"order_uid": orderUID}).
		SetResult(&result).
		Post("/customer/order/cancel")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
