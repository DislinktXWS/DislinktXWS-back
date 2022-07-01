package application

import "github.com/dislinktxws-back/authentication_service/domain"

type AuthenticationService struct {
	store domain.AuthenticationStore
}

func NewAuthenticationService(store domain.AuthenticationStore) *AuthenticationService {
	return &AuthenticationService{
		store: store,
	}
}

func (service *AuthenticationService) Login(auth *domain.Auth) (int64, string, string, bool) {
	return service.store.Login(auth)
}

func (service *AuthenticationService) Validate(token string) (int64, string, string) {
	return service.store.Validate(token)
}

func (service *AuthenticationService) Register(auth *domain.Auth) error {
	return service.store.Register(auth)
}

func (service *AuthenticationService) Delete(id string) error {
	return service.store.Delete(id)
}

func (service *AuthenticationService) EditUsername(auth *domain.Auth) (*domain.Auth, error) {
	return service.store.EditUsername(auth)
}

func (service *AuthenticationService) PasswordlessLogin(verificationToken string) (int64, string, string) {
	return service.store.PasswordlessLogin(verificationToken)
}

func (service *AuthenticationService) GenerateVerificationToken(email string) error {
	return service.store.GenerateVerificationToken(email)
}

func (service *AuthenticationService) ChangePassword(auth *domain.Auth) error {
	return service.store.ChangePassword(auth)
}

func (service *AuthenticationService) AccountRecovery(email string) (int64, string) {
	return service.store.AccountRecovery(email)
}

func (service *AuthenticationService) ChangeTwoFactorAuth(username string) (string, string) {
	return service.store.ChangeTwoFactorAuth(username)
}

func (service *AuthenticationService) GetTwoFactorAuth(username string) bool {
	return service.store.GetTwoFactorAuth(username)
}

func (service *AuthenticationService) VerifyTwoFactorAuthToken(username string, token string) (int64, string, string) {
	return service.store.VerifyTwoFactorAuthToken(username, token)
}
