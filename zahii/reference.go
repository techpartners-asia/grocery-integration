package zahii

import (
	"fmt"
)

type GuestReferenceService service

func (s *GuestReferenceService) SetLocationID(id string) *GuestReferenceService {
	s.locationID = id
	return s
}

type CustomerReferenceService service

func (s *CustomerReferenceService) SetLocationID(id string) *CustomerReferenceService {
	s.locationID = id
	return s
}

type ListAddressResponse struct {
	BaseResponse
	Body []*Address `json:"body"`
}

type ListAddressRequest struct {
	OrderField       int    `json:"order_field,omitempty"`
	ParentID         uint   `json:"parent_id,omitempty"`
	ParentRegionName string `json:"parent_region_name,omitempty"`
	RegionName       string `json:"region_name,omitempty"`
	RegionTypeID     int    `json:"region_type_id,omitempty"`
}

func (s *GuestReferenceService) ListAddresses(req ListAddressRequest) (*ListAddressResponse, error) {
	var result ListAddressResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/guest/reference/address/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListBannerResponse struct {
	BaseResponse
	Body []*Banner `json:"body"`
}

func (s *GuestReferenceService) ListBanners(branchID uint) (*ListBannerResponse, error) {
	var result ListBannerResponse
	_, err := s.client.newRequest(s.locationID).
		SetQueryParam("branch_id", fmt.Sprintf("%d", branchID)).
		SetResult(&result).
		Get("/guest/reference/banner/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateFeedbackRequest struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func (s *GuestReferenceService) CreateFeedback(req CreateFeedbackRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/guest/reference/feedback/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListLangResponse struct {
	BaseResponse
	Body []*Lang `json:"body"`
}

func (s *GuestReferenceService) ListLangs() (*ListLangResponse, error) {
	var result ListLangResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/reference/lang/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetAppInfoResponse struct {
	BaseResponse
	Body *AppInfo `json:"body"`
}

func (s *GuestReferenceService) GetAppInfo() (*GetAppInfoResponse, error) {
	var result GetAppInfoResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/reference/app/get")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *GuestReferenceService) GetHelp() (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/reference/help/get")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *GuestReferenceService) GetOrg(orgNo string) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/guest/reference/org/get/%s", orgNo))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListStartUpSliderResponse struct {
	BaseResponse
	Body []*StartUpSlider `json:"body"`
}

func (s *GuestReferenceService) ListStartUpSliders() (*ListStartUpSliderResponse, error) {
	var result ListStartUpSliderResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/reference/start-up-slider/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *GuestReferenceService) GetTermOfService() (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/guest/reference/term-of-service/get")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type PolygonCheckRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (s *GuestReferenceService) CheckPolygon(req PolygonCheckRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/guest/reference/polygon/check")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *CustomerReferenceService) CreateJobApplication(req JobApplicationRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/reference/job/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
