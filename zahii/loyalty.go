package zahii

import (
	"fmt"
)

type UserLoyaltyService service

type CreateInvitationRequest struct {
	ReceiverPhone string `json:"receiver_phone"`
}

func (s *UserLoyaltyService) SetLocationID(id string) *UserLoyaltyService {
	s.locationID = id
	return s
}

func (s *UserLoyaltyService) CreateInvitation(req CreateInvitationRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/user/loyalty/invitation/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserLoyaltyService) ListInvitations() (*ListInvitationResponse, error) {
	var result ListInvitationResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/user/loyalty/invitation/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserLoyaltyService) ListPromotions(branchID uint) (*ListLoyaltyPromotionResponse, error) {
	var result ListLoyaltyPromotionResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/user/loyalty/promotion/list/%d", branchID))
	if err != nil {
		return nil, err
	}
	return &result, nil
}
