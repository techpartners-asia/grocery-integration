package zahii

type CustomerProfileService service

func (s *CustomerProfileService) SetLocationID(id string) *CustomerProfileService {
	s.locationID = id
	return s
}

func (s *CustomerProfileService) GetCredit() (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/profile/credit")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type UpdateProfileRequest struct {
	BirthDate string `json:"birth_date,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Gender    string `json:"gender,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Register  string `json:"register,omitempty"`
}

func (s *CustomerProfileService) Update(req UpdateProfileRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Put("/customer/profile/update")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type AgeCheckRequest struct {
	BirthDate string `json:"birth_date"`
}

func (s *CustomerProfileService) AgeCheck(req AgeCheckRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/profile/age_check")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type PointHistoryRequest struct {
	Limit  int               `json:"limit"`
	Page   int               `json:"page"`
	Sorter map[string]string `json:"sorter"`
}

func (s *CustomerProfileService) GetPointHistory(req PointHistoryRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/profile/point/history/cursor")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
