package user

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
