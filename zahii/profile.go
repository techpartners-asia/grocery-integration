package zahii

type ProfileService service

func (s *ProfileService) SetLocationID(id string) *ProfileService {
	s.locationID = id
	return s
}

func (s *ProfileService) GetCredit() (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/user/profile/credit")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type InfoRequestDTO struct {
	BirthDate string `json:"birth_date,omitempty"`
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	Gender    string `json:"gender,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Register  string `json:"register,omitempty"`
	PushToken string `json:"push_token,omitempty"`
}

func (s *ProfileService) GetProfile(req InfoRequestDTO) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/user/info")
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

func (s *ProfileService) Update(req UpdateProfileRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Put("/user/profile/update")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type AgeCheckRequest struct {
	BirthDate string `json:"birth_date"`
}

func (s *ProfileService) AgeCheck(req AgeCheckRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/user/profile/age_check")
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

func (s *ProfileService) GetPointHistory(req PointHistoryRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/user/profile/point/history/cursor")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
