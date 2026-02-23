package zahii

import (
	"fmt"
)

type ReferenceService service

func (s *ReferenceService) SetLocationID(id string) *ReferenceService {
	s.locationID = id
	return s
}

type UserReferenceService service

func (s *UserReferenceService) SetLocationID(id string) *UserReferenceService {
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

func (s *ReferenceService) ListAddresses(req ListAddressRequest) (*ListAddressResponse, error) {
	var result ListAddressResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/reference/address/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListBannerResponse struct {
	BaseResponse
	Body []*Banner `json:"body"`
}

func (s *ReferenceService) ListBanners(branchID uint) (*ListBannerResponse, error) {
	var result ListBannerResponse
	_, err := s.client.newRequest(s.locationID).
		SetQueryParam("branch_id", fmt.Sprintf("%d", branchID)).
		SetResult(&result).
		Get("/reference/banner/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type CreateFeedbackRequest struct {
	Body  string `json:"body"`
	Title string `json:"title"`
}

func (s *ReferenceService) CreateFeedback(req CreateFeedbackRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/reference/feedback/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListLangResponse struct {
	BaseResponse
	Body []*Lang `json:"body"`
}

func (s *ReferenceService) ListLangs() (*ListLangResponse, error) {
	var result ListLangResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/reference/lang/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type GetAppInfoResponse struct {
	BaseResponse
	Body *AppInfo `json:"body"`
}

func (s *ReferenceService) GetAppInfo() (*GetAppInfoResponse, error) {
	var result GetAppInfoResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/reference/app/get")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ReferenceService) GetHelp() (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/reference/help/get")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ReferenceService) GetOrg(orgNo string) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get(fmt.Sprintf("/reference/org/get/%s", orgNo))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListStartUpSliderResponse struct {
	BaseResponse
	Body []*StartUpSlider `json:"body"`
}

func (s *ReferenceService) ListStartUpSliders() (*ListStartUpSliderResponse, error) {
	var result ListStartUpSliderResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/reference/start-up-slider/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *ReferenceService) GetTermOfService() (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/reference/term-of-service/get")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type PolygonCheckRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (s *ReferenceService) CheckPolygon(req PolygonCheckRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/reference/polygon/check")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *UserReferenceService) CreateJobApplication(req JobApplicationRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/user/reference/job/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
