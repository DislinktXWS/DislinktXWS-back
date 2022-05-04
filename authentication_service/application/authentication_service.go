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
