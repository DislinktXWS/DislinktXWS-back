package application

import "module/authentication_service/domain"

type AuthenticationService struct {
	store domain.AuthenticationStore
}

func NewAuthenticationService(store domain.AuthenticationStore) *AuthenticationService {
	return &AuthenticationService{
		store: store,
	}
}

func (service *AuthenticationService) Login(auth *domain.Auth) (int64, string, string) {
	return service.store.Login(auth)
}

func (service *AuthenticationService) Validate(token string) (int64, string, string) {
	return service.store.Validate(token)
}

func (service *AuthenticationService) Register(auth *domain.Auth) error {
	return service.store.Register(auth)
}

func (service *AuthenticationService) EditUsername(auth *domain.Auth) (*domain.Auth, error) {
	return service.store.EditUsername(auth)
}
