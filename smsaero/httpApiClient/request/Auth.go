package apiRequest

type BasicAuth struct {
	User     string
	Password string
}

func NewBasicAuth(user, password string) *BasicAuth {
	return &BasicAuth{User: user, Password: password}
}
