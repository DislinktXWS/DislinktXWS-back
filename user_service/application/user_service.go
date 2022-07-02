package application

import (
	"context"
	"fmt"
	"github.com/dislinktxws-back/user_service/domain"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	store        domain.UserStore
	orchestrator *InsertUserOrchestrator
}

func NewUserService(store domain.UserStore, orchestrator *InsertUserOrchestrator) *UserService {
	return &UserService{
		store:        store,
		orchestrator: orchestrator,
	}
}

func (service *UserService) Get(id primitive.ObjectID, ctx context.Context) (*domain.User, error) {
	return service.store.Get(id, ctx)
}

func (service *UserService) GetByApiKey(apiKey string, ctx context.Context) (*domain.User, error) {
	return service.store.GetByApiKey(apiKey, ctx)
}

func (service *UserService) GetByUsername(username string, ctx context.Context) (*domain.User, error) {
	return service.store.GetByUsername(username, ctx)
}

func (service *UserService) GetAll(ctx context.Context) ([]*domain.User, error) {
	return service.store.GetAll(ctx)
}

func (service *UserService) GetPublicUsers(ctx context.Context) ([]*domain.User, error) {
	return service.store.GetPublicUsers(ctx)
}

func (service *UserService) Insert(user *domain.User, ctx context.Context) (error, *domain.User) {
	err, newUser := service.store.Insert(user, ctx)
	fmt.Println("INSERTOVANO U USER")
	if err != nil {
		return err, nil
	}
	err = service.orchestrator.Start(newUser)
	if err != nil {
		return err, nil
	}
	return nil, newUser
}

func (service *UserService) Delete(id primitive.ObjectID) error {
	return service.store.Delete(id)
}

func (service *UserService) EditUser(user *domain.User, ctx context.Context) (*domain.User, error) {
	return service.store.EditUser(user, ctx)
}

func (service *UserService) EditUsername(user *domain.User, ctx context.Context) (*domain.User, error) {
	return service.store.EditUsername(user, ctx)
}

func (service *UserService) SetApiKey(username string, ctx context.Context) (string, error) {
	apiKey, _ := service.store.SetApiKey(username, ctx)
	return apiKey, nil
}

func (service *UserService) GetEducation(id primitive.ObjectID, ctx context.Context) (*[]domain.Education, error) {
	return service.store.GetEducation(id, ctx)
}

func (service *UserService) AddEducation(education *domain.Education, id primitive.ObjectID, ctx context.Context) (*domain.Education, error) {
	return service.store.AddEducation(education, id, ctx)
}

func (service *UserService) DeleteEducation(id primitive.ObjectID, index uint, ctx context.Context) error {
	return service.store.DeleteEducation(id, index, ctx)
}

func (service *UserService) GetExperience(id primitive.ObjectID, ctx context.Context) (*[]domain.Experience, error) {
	return service.store.GetExperience(id, ctx)
}

func (service *UserService) AddExperience(experience *domain.Experience, id primitive.ObjectID, ctx context.Context) (*domain.Experience, error) {
	return service.store.AddExperience(experience, id, ctx)
}

func (service *UserService) DeleteExperience(id primitive.ObjectID, index uint, ctx context.Context) error {
	return service.store.DeleteExperience(id, index, ctx)
}

func (service *UserService) GetInterests(id primitive.ObjectID, ctx context.Context) ([]string, error) {
	return service.store.GetInterests(id, ctx)
}

func (service *UserService) AddInterest(id primitive.ObjectID, interest string, ctx context.Context) error {
	return service.store.AddInterest(id, interest, ctx)
}

func (service *UserService) DeleteInterest(id primitive.ObjectID, index uint, ctx context.Context) error {
	return service.store.DeleteInterest(id, index, ctx)
}

func (service *UserService) GetSkills(id primitive.ObjectID, ctx context.Context) (*[]domain.Skill, error) {
	return service.store.GetSkills(id, ctx)
}

func (service *UserService) AddSkill(skill *domain.Skill, id primitive.ObjectID, ctx context.Context) (*domain.Skill, error) {
	return service.store.AddSkill(skill, id, ctx)
}

func (service *UserService) DeleteSkill(id primitive.ObjectID, index uint, ctx context.Context) error {
	return service.store.DeleteSkill(id, index, ctx)
}

func (service *UserService) SearchProfiles(search string, ctx context.Context) (*[]domain.User, error) {
	return service.store.SearchProfiles(search, ctx)
}

func (service *UserService) SetPrivacy(private bool, id primitive.ObjectID, ctx context.Context) error {
	return service.store.SetPrivacy(private, id, ctx)
}
