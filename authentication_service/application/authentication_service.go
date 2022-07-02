package application

import (
	"context"
	"github.com/dislinktxws-back/authentication_service/domain"
)

type AuthenticationService struct {
	store domain.AuthenticationStore
}

func NewAuthenticationService(store domain.AuthenticationStore) *AuthenticationService {
	return &AuthenticationService{
		store: store,
	}
}

func (service *AuthenticationService) Login(auth *domain.Auth, ctx context.Context) (int64, string, string, bool) {
	return service.store.Login(auth, ctx)
}

func (service *AuthenticationService) Validate(token string, ctx context.Context) (int64, string, string) {
	return service.store.Validate(token, ctx)
}

func (service *AuthenticationService) Register(auth *domain.Auth) error {
	return service.store.Register(auth)
}

func (service *AuthenticationService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *AuthenticationService) EditUsername(auth *domain.Auth, ctx context.Context) (*domain.Auth, error) {
	return service.store.EditUsername(auth, ctx)
}

func (service *AuthenticationService) PasswordlessLogin(verificationToken string, ctx context.Context) (int64, string, string) {
	return service.store.PasswordlessLogin(verificationToken, ctx)
}

func (service *AuthenticationService) GenerateVerificationToken(email string, ctx context.Context) error {
	return service.store.GenerateVerificationToken(email, ctx)
}

func (service *AuthenticationService) ChangePassword(auth *domain.Auth, ctx context.Context) error {
	return service.store.ChangePassword(auth, ctx)
}

func (service *AuthenticationService) AccountRecovery(email string, ctx context.Context) (int64, string) {
	return service.store.AccountRecovery(email, ctx)
}

func (service *AuthenticationService) ChangeTwoFactorAuth(username string, ctx context.Context) (string, string) {
	return service.store.ChangeTwoFactorAuth(username, ctx)
}

func (service *AuthenticationService) GetTwoFactorAuth(username string, ctx context.Context) bool {
	return service.store.GetTwoFactorAuth(username, ctx)
}

func (service *AuthenticationService) VerifyTwoFactorAuthToken(username string, token string, ctx context.Context) (int64, string, string) {
	return service.store.VerifyTwoFactorAuthToken(username, token, ctx)
}
