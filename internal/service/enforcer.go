package service

type EnforcerService struct {
	casdoorService *CasdoorService
}

func NewEnforcerService(inject *ServiceInjectInput) *EnforcerService {
	return &EnforcerService{casdoorService: inject.CasdoorService}
}

func (s *EnforcerService) Enforce(input *EnforceInput) (*EnforceResult, error) {
	return s.casdoorService.Enforce(input)
}
