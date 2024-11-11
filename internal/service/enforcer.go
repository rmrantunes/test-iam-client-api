package service

type EnforcerService struct {
	casdoorService *CasdoorService
}

func NewEnforcerService(casdoorService *CasdoorService) *EnforcerService {
	return &EnforcerService{casdoorService: casdoorService}
}

func (s *EnforcerService) Enforce(input *EnforceInput) (*EnforceResult, error) {
	return s.casdoorService.Enforce(input)
}
