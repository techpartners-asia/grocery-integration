package zahii

type ImapService service

type ImapSsidResponse struct {
	BaseResponse
	Body string `json:"body"`
}

func (s *ImapService) SetLocationID(id string) *ImapService {
	s.locationID = id
	return s
}

func (s *ImapService) GetSsid() (*ImapSsidResponse, error) {
	var result ImapSsidResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/user/imap/ssid")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
