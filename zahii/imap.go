package zahii

type CustomerImapService service

type ImapSsidResponse struct {
	BaseResponse
	Body string `json:"body"`
}

func (s *CustomerImapService) SetLocationID(id string) *CustomerImapService {
	s.locationID = id
	return s
}

func (s *CustomerImapService) GetSsid() (*ImapSsidResponse, error) {
	var result ImapSsidResponse
	_, err := s.client.newRequest(s.locationID).
		SetResult(&result).
		Get("/customer/imap/ssid")
	if err != nil {
		return nil, err
	}
	return &result, nil
}
