package domain

import "context"

type AuthenticationStore interface {
	Login(auth *Auth, ctx context.Context) (status int64, error string, token string, isTwoFactorEnabled bool)
	PasswordlessLogin(verificationToken string, ctx context.Context) (status int64, error string, token string)
	GenerateVerificationToken(email string, ctx context.Context) error
	Validate(token string, ctx context.Context) (status int64, error string, username string)
	Register(auth *Auth) error
	Delete(id string) error
	EditUsername(auth *Auth, ctx context.Context) (*Auth, error)
	ChangePassword(auth *Auth, ctx context.Context) error
	AccountRecovery(email string, ctx context.Context) (status int64, error string)
	ChangeTwoFactorAuth(username string, ctx context.Context) (qrCode string, error string)
	GetTwoFactorAuth(username string, ctx context.Context) bool
	VerifyTwoFactorAuthToken(username string, twoAuthToken string, ctx context.Context) (status int64, error string, JWTtoken string)
}
