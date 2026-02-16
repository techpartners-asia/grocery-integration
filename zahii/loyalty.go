package zahii

import (
	"fmt"
)

type LoyaltyService service

type CreateInvitationRequest struct {
	ReceiverPhone string `json:"receiver_phone"`
}

func (s *LoyaltyService) SetLocationID(id string) *LoyaltyService {
	s.locationID = id
	return s
}

func (s *LoyaltyService) Create(req CreateInvitationRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/loyalty/invitation/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LoyaltyService) List() (*ListInvitationResponse, error) {
	var result ListInvitationResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/customer/loyalty/invitation/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LoyaltyService) ListPromotions(branchID uint) (*ListLoyaltyPromotionResponse, error) {
	var result ListLoyaltyPromotionResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/customer/loyalty/promotion/list/%d", branchID))
	if err != nil {
		return nil, err
	}
	return &result, nil
}
