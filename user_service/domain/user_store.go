package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type UserStore interface {
	Get(id primitive.ObjectID) (*User, error)
	GetAll() ([]*User, error)
	Insert(user *User) error
	DeleteAll()
	EditUser(user *User) (*User, error)
	AddEducation(education *Education, id primitive.ObjectID) (*Education, error)
}
