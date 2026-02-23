package zahii

type CommentService service

type CreateCommentRequest struct {
	Body         string `json:"body"`
	CustomerID   uint   `json:"customer_id"`
	ProductID    uint   `json:"product_id,omitempty"`
	ProductSetID uint   `json:"product_set_id,omitempty"`
	Rate         int    `json:"rate"`
	Title        string `json:"title"`
}

func (s *CommentService) SetLocationID(id string) *CommentService {
	s.locationID = id
	return s
}

type CreateCommentResponse struct {
	BaseResponse
	Body struct {
		ID uint `json:"id"`
	} `json:"body"`
}

func (s *CommentService) Create(req CreateCommentRequest) (*CreateCommentResponse, error) {
	var result CreateCommentResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/user/comment/create")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type DeleteCommentRequest struct {
	ID uint `json:"id"`
}

func (s *CommentService) Delete(req DeleteCommentRequest) (*BaseResponse, error) {
	var result BaseResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Delete("/user/comment/delete")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

type ListCommentResponse struct {
	BaseResponse
	Body struct {
		Items []struct {
			ID    uint   `json:"id"`
			Body  string `json:"body"`
			Rate  int    `json:"rate"`
			Title string `json:"title"`
		} `json:"items"`
		Total int `json:"total"`
	} `json:"body"`
}

type ListCommentRequest struct {
	Limit int `json:"limit,omitempty"`
	Page  int `json:"page,omitempty"`
}

func (s *CommentService) List(req ListCommentRequest) (*ListCommentResponse, error) {
	var result ListCommentResponse
	_, err := s.client.newRequest(s.locationID).
		SetBody(req).
		SetResult(&result).
		Post("/user/comment/list")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
