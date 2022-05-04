package domain

type AuthenticationStore interface {
	Login(auth *Auth) (status int64, error string, token string)
	Validate(token string) (status int64, error string, username string)
	Register(auth *Auth) error
}
