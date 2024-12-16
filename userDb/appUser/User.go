package appUser

type AppUser struct {
	Id       int
	Username string
	Password string
}

func New(username string, password string) *AppUser {
	return &AppUser{
		Username: username,
		Password: password,
	}
}
