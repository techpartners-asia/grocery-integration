package zahii

type SuperAppAuthenticateService service

func (s *SuperAppAuthenticateService) SetLocationID(id string) *SuperAppAuthenticateService {
	s.locationID = id
	return s
}

func (s *SuperAppAuthenticateService) Authenticate(token string) (*AuthenticateResponse, error) {
	var result AuthenticateResponse
	_, err := s.client.newBaseRequest(s.locationID).
		SetPathParam("token", token).
		SetResult(&result).
		Get("/super-app/authenticate/{token}")
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (s *SuperAppAuthenticateService) AuthenticateAndSetToken(token string) (*AuthenticateResponse, error) {
	resp, err := s.Authenticate(token)
	if err != nil {
		return nil, err
	}

	if resp.Body.Token != "" {
		s.client.SetAuthToken(resp.Body.Token)
	}

	return resp, nil
}
