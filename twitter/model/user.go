package model

type User struct {
	Id       string
	Username string
	Password string
	Key      string
	Token    string
}

type Credentials struct {
	Refresh_token string
	Grant_type    string
	Client_id     string
}

func NewCredentials(user User) Credentials {
	return Credentials{
		Refresh_token: user.Token,
		Grant_type:    "refresh_token",
		Client_id:     user.Key,
	}
}
