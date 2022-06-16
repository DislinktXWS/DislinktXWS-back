package domain

type AuthenticationStore interface {
	Login(auth *Auth) (status int64, error string, token string, isTwoFactorEnabled bool)
	PasswordlessLogin(verificationToken string) (status int64, error string, token string)
	GenerateVerificationToken(email string) error
	Validate(token string) (status int64, error string, username string)
	Register(auth *Auth) error
	EditUsername(auth *Auth) (*Auth, error)
	ChangePassword(auth *Auth) error
	AccountRecovery(email string) (status int64, error string)
	ChangeTwoFactorAuth(username string) (qrCode string, error string)
	GetTwoFactorAuth(username string) bool
	VerifyTwoFactorAuthToken(username string, twoAuthToken string) (status int64, error string, JWTtoken string)
}
