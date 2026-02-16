package zahii

import (
	"fmt"
)

type GuestCategoryService service

func (s *GuestCategoryService) SetLocationID(id string) *GuestCategoryService {
	s.locationID = id
	return s
}

type GuestProductService service

func (s *GuestProductService) SetLocationID(id string) *GuestProductService {
	s.locationID = id
	return s
}

type GuestCustomerService service

func (s *GuestCustomerService) SetLocationID(id string) *GuestCustomerService {
	s.locationID = id
	return s
}

type GuestLoyaltyService service

func (s *GuestLoyaltyService) SetLocationID(id string) *GuestLoyaltyService {
	s.locationID = id
	return s
}

type GuestOrderMessageService service

func (s *GuestOrderMessageService) SetLocationID(id string) *GuestOrderMessageService {
	s.locationID = id
	return s
}

type GuestStoreService service

func (s *GuestStoreService) SetLocationID(id string) *GuestStoreService {
	s.locationID = id
	return s
}

type GuestTagService service

func (s *GuestTagService) SetLocationID(id string) *GuestTagService {
	s.locationID = id
	return s
}

type ListCategoryRequest struct {
	Active bool `json:"active"`
}

type ListCategoryResponse struct {
	BaseResponse
	Body []*Category `json:"body"`
}

func (s *GuestCategoryService) List(req ListCategoryRequest) (*ListCategoryResponse, error) {
	var result ListCategoryResponse
	_, err := s.client.newRequest(s.locationID).
		SetQueryParams(map[string]string{
			"active": fmt.Sprintf("%v", req.Active),
		}).
		SetResult(&result).
		Get("/guest/category/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetCategoryResponse struct {
	BaseResponse
	Body *Category `json:"body"`
}

func (s *GuestCategoryService) Get(id uint) (*GetCategoryResponse, error) {
	var result GetCategoryResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/guest/category/get/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListProductRequest struct {
	Active                 bool   `json:"active,omitempty"`
	BranchID               uint   `json:"branch_id,omitempty"`
	CategoryID             uint   `json:"category_id,omitempty"`
	CategoryIDs            []uint `json:"category_ids,omitempty"`
	ClassificationCode     string `json:"classification_code,omitempty"`
	IsCityTax              bool   `json:"is_city_tax,omitempty"`
	IsClient               bool   `json:"is_client,omitempty"`
	IsNewlyAdded           bool   `json:"is_newly_added,omitempty"`
	IsNotSinglePurchasable bool   `json:"is_not_single_purchasable,omitempty"`
	IsPackage              bool   `json:"is_package,omitempty"`
	IsPromotion            bool   `json:"is_promotion,omitempty"`
	IsSale                 bool   `json:"is_sale,omitempty"`
	IsStoreProduct         bool   `json:"is_store_product,omitempty"`
	Limit                  int    `json:"limit,omitempty"`
	LocationID             string `json:"location_id,omitempty"`
	Page                   int    `json:"page,omitempty"`
	Popular                bool   `json:"popular,omitempty"`
	Search                 string `json:"search,omitempty"`
	StoreID                uint   `json:"store_id,omitempty"`
	TagID                  uint   `json:"tag_id,omitempty"`
	TaxProductCode         string `json:"tax_product_code,omitempty"`
	TaxType                string `json:"tax_type,omitempty"`
	TodaysSuggest          bool   `json:"todays_suggest,omitempty"`
}

type ListProductResponse struct {
	BaseResponse
	Body []*Product `json:"body"`
}

func (s *GuestProductService) List(req ListProductRequest) (*ListProductResponse, error) {
	var result ListProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/guest/product/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetProductResponse struct {
	BaseResponse
	Body *Product `json:"body"`
}

func (s *GuestProductService) Get(id uint, locationID string) (*GetProductResponse, error) {
	var result GetProductResponse
	reqLocationID := s.locationID
	if locationID != "" {
		reqLocationID = locationID
	}
	req := s.client.newRequest(reqLocationID).SetResult(&result)
	_, err := req.Get(fmt.Sprintf("/guest/product/get/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type UserExistsResponse struct {
	BaseResponse
	Body bool `json:"body"`
}

func (s *GuestCustomerService) IsUserExists(firebaseUID string) (*UserExistsResponse, error) {
	var result UserExistsResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/guest/customer/is-user-exists/%s", firebaseUID))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *GuestProductService) ListSimilar(id uint) (*ListProductResponse, error) {
	var result ListProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post(fmt.Sprintf("/guest/product/similar-product/list/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetProductSetResponse struct {
	BaseResponse
	Body *ProductSet `json:"body"`
}

func (s *GuestProductService) GetSetProduct(id uint) (*GetProductSetResponse, error) {
	var result GetProductSetResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/guest/product/set-product/get/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListProductSetResponse struct {
	BaseResponse
	Body []*ProductSet `json:"body"`
}

func (s *GuestProductService) ListSetProduct() (*ListProductSetResponse, error) {
	var result ListProductSetResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/guest/product/set-product/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *GuestProductService) GetStoreProduct(id uint) (*GetProductResponse, error) {
	var result GetProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/guest/product/store-product/get/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type TotalProductResponse struct {
	BaseResponse
	Body int `json:"body"`
}

func (s *GuestProductService) GetTotal() (*TotalProductResponse, error) {
	var result TotalProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/product/total")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListStoreResponse struct {
	BaseResponse
	Body []*Store `json:"body"`
}

func (s *GuestStoreService) List() (*ListStoreResponse, error) {
	var result ListStoreResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/guest/store/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListLoyaltyLevelResponse struct {
	BaseResponse
	Body []*LoyaltyLevel `json:"body"`
}

func (s *GuestLoyaltyService) ListLevels() (*ListLoyaltyLevelResponse, error) {
	var result ListLoyaltyLevelResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/loyalty/level/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *GuestLoyaltyService) ListInvitations() (*ListInvitationResponse, error) {
	var result ListInvitationResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/loyalty/invitation/get-all")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetInvitationResponse struct {
	BaseResponse
	Body *Invitation `json:"body"`
}

func (s *GuestLoyaltyService) GetInvitation() (*GetInvitationResponse, error) {
	var result GetInvitationResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/loyalty/invitation/get-invitation")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type OrderMessage struct {
	Base
	Title string `json:"title"`
	Body  string `json:"body"`
}

type GetOrderMessageResponse struct {
	BaseResponse
	Body *OrderMessage `json:"body"`
}

func (s *GuestOrderMessageService) Get() (*GetOrderMessageResponse, error) {
	var result GetOrderMessageResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/order-message/get")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListOrderMessageResponse struct {
	BaseResponse
	Body []*OrderMessage `json:"body"`
}

func (s *GuestOrderMessageService) List() (*ListOrderMessageResponse, error) {
	var result ListOrderMessageResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/order-message/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListTagResponse struct {
	BaseResponse
	Body []struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
	} `json:"body"`
}

func (s *GuestTagService) List() (*ListTagResponse, error) {
	var result ListTagResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/guest/tag/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
