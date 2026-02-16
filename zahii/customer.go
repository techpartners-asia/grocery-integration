package zahii

type CustomerService service

func (s *CustomerService) SetLocationID(id string) *CustomerService {
	s.locationID = id
	return s
}
