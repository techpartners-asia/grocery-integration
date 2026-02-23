package zahii

import (
	"fmt"
)

type CategoryService service

func (s *CategoryService) SetLocationID(id string) *CategoryService {
	s.locationID = id
	return s
}

type ProductService service

func (s *ProductService) SetLocationID(id string) *ProductService {
	s.locationID = id
	return s
}

type StoreService service

func (s *StoreService) SetLocationID(id string) *StoreService {
	s.locationID = id
	return s
}

type TagService service

func (s *TagService) SetLocationID(id string) *TagService {
	s.locationID = id
	return s
}

type LoyaltyService service

func (s *LoyaltyService) SetLocationID(id string) *LoyaltyService {
	s.locationID = id
	return s
}

type OrderMessageService service

func (s *OrderMessageService) SetLocationID(id string) *OrderMessageService {
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

func (s *CategoryService) List(req ListCategoryRequest) (*ListCategoryResponse, error) {
	var result ListCategoryResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/category/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetCategoryResponse struct {
	BaseResponse
	Body *Category `json:"body"`
}

func (s *CategoryService) Get(id uint) (*GetCategoryResponse, error) {
	var result GetCategoryResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/category/get/%d", id))
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

func (s *ProductService) List(req ListProductRequest) (*ListProductResponse, error) {
	var result ListProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/product/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetProductResponse struct {
	BaseResponse
	Body *Product `json:"body"`
}

func (s *ProductService) Get(id uint, locationID string) (*GetProductResponse, error) {
	var result GetProductResponse
	reqLocationID := s.locationID
	if locationID != "" {
		reqLocationID = locationID
	}
	req := s.client.newRequest(reqLocationID).SetResult(&result)
	_, err := req.Get(fmt.Sprintf("/product/get/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type UserExistsResponse struct {
	BaseResponse
	Body bool `json:"body"`
}

func (s *CustomerService) IsUserExists(firebaseUID string) (*UserExistsResponse, error) {
	var result UserExistsResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/auth/is-user-exists/%s", firebaseUID))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type SimilarProductRequest struct {
	ProductID uint `json:"product_id"`
	Limit     int  `json:"limit,omitempty"`
	Page      int  `json:"page,omitempty"`
}

func (s *ProductService) ListSimilar(req SimilarProductRequest) (*ListProductResponse, error) {
	var result ListProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/product/similar-product/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetProductSetResponse struct {
	BaseResponse
	Body *ProductSet `json:"body"`
}

func (s *ProductService) GetSetProduct(id uint) (*GetProductSetResponse, error) {
	var result GetProductSetResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/product/set-product/get/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListProductSetResponse struct {
	BaseResponse
	Body []*ProductSet `json:"body"`
}

func (s *ProductService) ListSetProduct() (*ListProductSetResponse, error) {
	var result ListProductSetResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/product/set-product/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ProductService) GetStoreProduct(id uint) (*GetProductResponse, error) {
	var result GetProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/product/store-product/get/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type TotalProductResponse struct {
	BaseResponse
	Body int `json:"body"`
}

type TotalProductRequest struct {
	CategoryID         *uint   `json:"category_id,omitempty"`
	ClassificationCode *string `json:"classification_code,omitempty"`
	IsCityTax          *bool   `json:"is_city_tax,omitempty"`
	TaxProductCode     *string `json:"tax_product_code,omitempty"`
	TaxType            *string `json:"tax_type,omitempty"`
	TagID              *uint   `json:"tag_id,omitempty"`
	CategoryIDs        []uint  `json:"category_ids,omitempty"`
	Popular            *bool   `json:"popular,omitempty"`
	IsNewlyAdded       *bool   `json:"is_newly_added,omitempty"`
	TodaysSuggest      *bool   `json:"todays_suggest,omitempty"`
	IsSale             *bool   `json:"is_sale,omitempty"`
	StoreID            *uint   `json:"store_id,omitempty"`
	IsPromotion        *bool   `json:"is_promotion,omitempty"`
	Active             *bool   `json:"active,omitempty"`
	BranchID           *uint   `json:"branch_id,omitempty"`
	IsPackage          *bool   `json:"is_package,omitempty"`
	Search             string  `json:"search,omitempty"`
}

func (s *ProductService) GetTotal(req TotalProductRequest) (*TotalProductResponse, error) {
	var result TotalProductResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/product/total")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListStoreResponse struct {
	BaseResponse
	Body []*Store `json:"body"`
}

type ListStoreRequest struct {
	Search string `json:"search,omitempty"`
	Limit  int    `json:"limit,omitempty"`
	Page   int    `json:"page,omitempty"`
}

func (s *StoreService) List(req ListStoreRequest) (*ListStoreResponse, error) {
	var result ListStoreResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/store/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListLoyaltyLevelResponse struct {
	BaseResponse
	Body []*LoyaltyLevel `json:"body"`
}

func (s *LoyaltyService) ListLevels() (*ListLoyaltyLevelResponse, error) {
	var result ListLoyaltyLevelResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/loyalty/level/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LoyaltyService) ListInvitations() (*ListInvitationResponse, error) {
	var result ListInvitationResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/loyalty/invitation/get-all")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetInvitationResponse struct {
	BaseResponse
	Body *Invitation `json:"body"`
}

func (s *LoyaltyService) GetInvitation() (*GetInvitationResponse, error) {
	var result GetInvitationResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/loyalty/invitation/get-invitation")
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

func (s *OrderMessageService) Get() (*GetOrderMessageResponse, error) {
	var result GetOrderMessageResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/order-message/get")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListOrderMessageResponse struct {
	BaseResponse
	Body []*OrderMessage `json:"body"`
}

func (s *OrderMessageService) List() (*ListOrderMessageResponse, error) {
	var result ListOrderMessageResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/order-message/list")
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

type ListTagRequest struct {
	Name  *string `json:"name,omitempty"`
	Limit int     `json:"limit,omitempty"`
	Page  int     `json:"page,omitempty"`
}

func (s *TagService) List(req ListTagRequest) (*ListTagResponse, error) {
	var result ListTagResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/tag/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
