package application

import (
	"module/user_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store domain.UserStore
}

func NewUserService(store domain.UserStore) *UserService {
	return &UserService{
		store: store,
	}
}

func (service *UserService) Get(id primitive.ObjectID) (*domain.User, error) {
	return service.store.Get(id)
}

func (service *UserService) GetAll() ([]*domain.User, error) {
	return service.store.GetAll()
}

func (service *UserService) Insert(user *domain.User) error {
	return service.store.Insert(user)
}

func (service *UserService) EditUser(user *domain.User) (*domain.User, error) {
	return service.store.EditUser(user)
}

func (service *UserService) AddEducation(education *domain.Education, id primitive.ObjectID) (*domain.Education, error) {
	return service.store.AddEducation(education, id)
}

func (service *UserService) DeleteEducation(id primitive.ObjectID, index uint) error {
	return service.store.DeleteEducation(id, index)
}

func (service *UserService) AddExperience(experience *domain.Experience, id primitive.ObjectID) (*domain.Experience, error) {
	return service.store.AddExperience(experience, id)
}

func (service *UserService) DeleteExperience(id primitive.ObjectID, index uint) error {
	return service.store.DeleteExperience(id, index)
}
