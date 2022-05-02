package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) (error, *User)
	DeleteAll()
	EditUser(user *User) (*User, error)
	AddEducation(education *Education, id primitive.ObjectID) (*Education, error)
	DeleteEducation(id primitive.ObjectID, index uint) error
	AddExperience(experience *Experience, id primitive.ObjectID) (*Experience, error)
	DeleteExperience(id primitive.ObjectID, index uint) error
	AddInterest(id primitive.ObjectID, interest string) error
	DeleteInterest(id primitive.ObjectID, index uint) error
}
