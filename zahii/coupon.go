package zahii

type CouponService service

type ListCouponResponse struct {
	BaseResponse
	Body []*CustomerCoupon `json:"body"`
}

func (s *CouponService) SetLocationID(id string) *CouponService {
	s.locationID = id
	return s
}

func (s *CouponService) ListActiveCoupons() (*ListCouponResponse, error) {
	var result ListCouponResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/coupon/active")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateCouponRequest struct {
	Amount   float64 `json:"amount"`
	Code     string  `json:"code"`
	EndDate  string  `json:"end_date"`
	IsActive bool    `json:"is_active"`
}

func (s *CouponService) CreateCoupon(req CreateCouponRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/coupon/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CouponService) ListHistory() (*ListCouponResponse, error) {
	var result ListCouponResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/customer/coupon/history")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
