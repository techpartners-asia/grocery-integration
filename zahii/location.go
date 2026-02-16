package zahii

import "fmt"

type BranchService service

func (s *BranchService) SetLocationID(id string) *BranchService {
	s.locationID = id
	return s
}

type ListBranchResponse struct {
	BaseResponse
	Body []*Branch `json:"body"`
}

func (s *BranchService) ListBranch() (*ListBranchResponse, error) {
	var result ListBranchResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/branch/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type LocationService service

func (s *LocationService) SetLocationID(id string) *LocationService {
	s.locationID = id
	return s
}

type SaveLocationRequest struct {
	ID             uint    `json:"id,omitempty"`
	Name           string  `json:"name"`
	IsMain         bool    `json:"is_main"`
	SumDuureg      string  `json:"sum_duureg"`
	BagHoroo       string  `json:"bag_horoo"`
	EntranceNumber string  `json:"entrance_number,omitempty"`
	Apartment      string  `json:"apartment,omitempty"`
	Floor          string  `json:"floor,omitempty"`
	DoorNumber     string  `json:"door_number,omitempty"`
	EntranceCode   string  `json:"entrance_code,omitempty"`
	Note           string  `json:"note,omitempty"`
	Latitude       float64 `json:"latitude"`
	Longitude      float64 `json:"longitude"`
	BranchID       uint    `json:"branch_id"`
}

type LocationResponse struct {
	BaseResponse
	Body struct {
		ID             uint    `json:"id"`
		Name           string  `json:"name"`
		IsMain         bool    `json:"is_main"`
		SumDuureg      string  `json:"sum_duureg"`
		BagHoroo       string  `json:"bag_horoo"`
		EntranceNumber string  `json:"entrance_number"`
		Apartment      string  `json:"apartment"`
		Floor          string  `json:"floor"`
		DoorNumber     string  `json:"door_number"`
		EntranceCode   string  `json:"entrance_code"`
		Note           string  `json:"note"`
		Latitude       float64 `json:"latitude"`
		Longitude      float64 `json:"longitude"`
		CustomerID     uint    `json:"customer_id"`
		BranchID       uint    `json:"branch_id"`
		Rank           int     `json:"rank"`
	} `json:"body"`
}

func (s *LocationService) Create(req SaveLocationRequest) (*LocationResponse, error) {
	var result LocationResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/location/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LocationService) Delete(id uint) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Delete(fmt.Sprintf("/customer/location/delete/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListLocationResponse struct {
	BaseResponse
	Body []struct {
		ID        uint    `json:"id"`
		IsMain    bool    `json:"is_main"`
		Name      string  `json:"name"`
		SumDuureg string  `json:"sum_duureg"`
		BagHoroo  string  `json:"bag_horoo"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
		Rank      int     `json:"rank"`
	} `json:"body"`
}

func (s *LocationService) List() (*ListLocationResponse, error) {
	var result ListLocationResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/location/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LocationService) ListPolygon() (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/location/polygon/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *LocationService) Update(id uint, req SaveLocationRequest) (*LocationResponse, error) {
	var result LocationResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Put(fmt.Sprintf("/customer/location/update/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type UpdateLatLongRequest struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (s *LocationService) UpdateLatLong(id uint, req UpdateLatLongRequest) (*LocationResponse, error) {
	var result LocationResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Put(fmt.Sprintf("/customer/location/update-lat-long/%d", id))
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type UpdateRanksRequest struct {
	Ranks map[string]int `json:"ranks"`
}

func (s *LocationService) UpdateRanks(req UpdateRanksRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Put("/customer/location/update-ranks")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
