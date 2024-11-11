package service

type User struct {
	ID   int
	Name string
}

type UserService struct {
	casdoorService *CasdoorService
}

func NewUserService(casdoorService *CasdoorService) *UserService {
	return &UserService{casdoorService: casdoorService}
}

func (s *UserService) GetUsers() []User {
	return []User{
		{ID: 1, Name: "User1"},
		{ID: 2, Name: "User2"},
	}
}
