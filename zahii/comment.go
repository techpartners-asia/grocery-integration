package zahii

type CustomerCommentService service

type CreateCommentRequest struct {
	Body         string `json:"body"`
	CustomerID   uint   `json:"customer_id"`
	ProductID    uint   `json:"product_id,omitempty"`
	ProductSetID uint   `json:"product_set_id,omitempty"`
	Rate         int    `json:"rate"`
	Title        string `json:"title"`
}

func (s *CustomerCommentService) SetLocationID(id string) *CustomerCommentService {
	s.locationID = id
	return s
}

func (s *CustomerCommentService) Create(req CreateCommentRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/customer/comment/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type DeleteCommentRequest struct {
	ID uint `json:"id"`
}

func (s *CustomerCommentService) Delete(req DeleteCommentRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Delete("/customer/comment/delete")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListCommentResponse struct {
	BaseResponse
	Body []struct {
		ID    uint   `json:"id"`
		Body  string `json:"body"`
		Rate  int    `json:"rate"`
		Title string `json:"title"`
	} `json:"body"`
}

func (s *CustomerCommentService) List() (*ListCommentResponse, error) {
	var result ListCommentResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Post("/customer/comment/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
