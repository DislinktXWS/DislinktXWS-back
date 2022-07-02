package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserStore interface {
	Get(id primitive.ObjectID, ctx context.Context) (*User, error)
	GetByUsername(username string, ctx context.Context) (*User, error)
	GetByApiKey(apiKey string, ctx context.Context) (*User, error)
	GetAll(ctx context.Context) ([]*User, error)
	GetPublicUsers(ctx context.Context) ([]*User, error)
	Insert(user *User, ctx context.Context) (error, *User)
	Delete(id primitive.ObjectID) error
	DeleteAll(ctx context.Context)
	EditUser(user *User, ctx context.Context) (*User, error)
	EditUsername(user *User, ctx context.Context) (*User, error)
	SetApiKey(username string, ctx context.Context) (string, error)
	GetEducation(id primitive.ObjectID, ctx context.Context) (*[]Education, error)
	AddEducation(education *Education, id primitive.ObjectID, ctx context.Context) (*Education, error)
	DeleteEducation(id primitive.ObjectID, index uint, ctx context.Context) error
	GetExperience(id primitive.ObjectID, ctx context.Context) (*[]Experience, error)
	AddExperience(experience *Experience, id primitive.ObjectID, ctx context.Context) (*Experience, error)
	DeleteExperience(id primitive.ObjectID, index uint, ctx context.Context) error
	GetInterests(id primitive.ObjectID, ctx context.Context) ([]string, error)
	AddInterest(id primitive.ObjectID, interest string, ctx context.Context) error
	DeleteInterest(id primitive.ObjectID, index uint, ctx context.Context) error
	GetSkills(id primitive.ObjectID, ctx context.Context) (*[]Skill, error)
	AddSkill(skill *Skill, id primitive.ObjectID, ctx context.Context) (*Skill, error)
	DeleteSkill(id primitive.ObjectID, index uint, ctx context.Context) error
	SearchProfiles(search string, ctx context.Context) (*[]User, error)
	SetPrivacy(public bool, id primitive.ObjectID, ctx context.Context) error
}
