package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Surname     string             `bson:"surname"`
	Username    string             `bson:"username"`
	Password    string             `bson:"password"`
	DateOfBirth string             `bson:"dateOfBirth"`
	Gender      string             `bson:"gender"`
	Email       string             `bson:"email"`
	Phone       string             `bson:"phone"`
	Education   []Education        `bson:"education"`
	Experience  []Experience       `bson:"experience"`
	Interests   []string           `bson:"interests"`
}

type Education struct {
	School       string  `bson:"school"`
	Degree       string  `bson:"degree"`
	FieldOfStudy string  `bson:"fieldOfStudy"`
	StartDate    string  `bson:"startDate"`
	EndDate      string  `bson:"endDate"`
	Grade        float32 `bson:"grade"`
	Description  string  `bson:"description"`
}

type Experience struct {
	Title       string `bson:"title"`
	CompanyName string `bson:"companyName"`
	StartDate   string `bson:"startDate"`
	EndDate     string `bson:"endDate"`
	Industry    string `bson:"industry"`
	Description string `bson:"description"`
}

type Skill struct {
	Name        string           `bson:"name"`
	Proficiency SkillProficiency `bson:"proficiency"`
}

type SkillProficiency int8

const (
	novice SkillProficiency = iota
	advancedBeginner
	proficient
	expert
	master
)

func (proficiency SkillProficiency) String() string {
	switch proficiency {
	case novice:
		return "novice"
	case advancedBeginner:
		return "advanced beginner"
	case proficient:
		return "proficient"
	case expert:
		return "expert"
	case master:
		return "master"
	}

	return "unknown"
}
