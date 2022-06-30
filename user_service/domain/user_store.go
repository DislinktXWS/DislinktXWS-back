package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetByUsername(username string) (*User, error)
	GetByApiKey(apiKey string) (*User, error)
	GetAll() ([]*User, error)
	GetPublicUsers() ([]*User, error)
	Insert(user *User) (error, *User)
	Delete(id primitive.ObjectID) error
	DeleteAll()
	EditUser(user *User) (*User, error)
	EditUsername(user *User) (*User, error)
	SetApiKey(username string) (string, error)
	GetEducation(id primitive.ObjectID) (*[]Education, error)
	AddEducation(education *Education, id primitive.ObjectID) (*Education, error)
	DeleteEducation(id primitive.ObjectID, index uint) error
	GetExperience(id primitive.ObjectID) (*[]Experience, error)
	AddExperience(experience *Experience, id primitive.ObjectID) (*Experience, error)
	DeleteExperience(id primitive.ObjectID, index uint) error
	GetInterests(id primitive.ObjectID) ([]string, error)
	AddInterest(id primitive.ObjectID, interest string) error
	DeleteInterest(id primitive.ObjectID, index uint) error
	GetSkills(id primitive.ObjectID) (*[]Skill, error)
	AddSkill(skill *Skill, id primitive.ObjectID) (*Skill, error)
	DeleteSkill(id primitive.ObjectID, index uint) error
	SearchProfiles(search string) (*[]User, error)
	SetPrivacy(public bool, id primitive.ObjectID) error
}
