package zahii

import (
	"fmt"
)

type WishlistService service

func (s *WishlistService) SetLocationID(id string) *WishlistService {
	s.locationID = id
	return s
}

type ListWishlistResponse struct {
	BaseResponse
	Body []*Wishlist `json:"body"`
}

func (s *WishlistService) List() (*ListWishlistResponse, error) {
	var result ListWishlistResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/wishlist/base/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateWishlistRequest struct {
	Name         string `json:"name"`
	ProductID    uint   `json:"product_id,omitempty"`
	ProductSetID uint   `json:"product_set_id,omitempty"`
}

func (s *WishlistService) Create(req CreateWishlistRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/wishlist/base/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *WishlistService) Delete(id uint) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(&result).
		Delete("/customer/wishlist/base/delete/{id}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetWishlistResponse struct {
	BaseResponse
	Body *Wishlist `json:"body"`
}

func (s *WishlistService) Get(id uint) (*GetWishlistResponse, error) {
	var result GetWishlistResponse
	_, err := s.client.newRequest(s.locationID).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(&result).
		Get("/customer/wishlist/base/get/{id}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *WishlistService) Update(id uint, req CreateWishlistRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetBody(req).
		SetResult(&result).
		Put("/customer/wishlist/base/update/{id}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListWishlistItemResponse struct {
	BaseResponse
	Body []*WishlistItem `json:"body"`
}

func (s *WishlistService) ListItem() (*ListWishlistItemResponse, error) {
	var result ListWishlistItemResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/wishlist/item/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *WishlistService) AddItem(id uint, productID uint) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetBody(map[string]uint{"product_id": productID}).
		SetResult(&result).
		Post("/customer/wishlist/item/add/{id}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *WishlistService) DeleteItem(id uint) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetPathParam("id", fmt.Sprintf("%d", id)).
		SetResult(&result).
		Delete("/customer/wishlist/item/delete/{id}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
