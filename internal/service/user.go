package service_users

type User struct {
	ID   int
	Name string
}

func GetUsers() []User {
	return []User{
		{ID: 1, Name: "User1"},
		{ID: 2, Name: "User2"},
	}
}
