package domain

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Surname     string             `bson:"surname"`
	Username    string             `bson:"username"`
	DateOfBirth string             `bson:"dateOfBirth"`
	Gender      string             `bson:"gender"`
	Email       string             `bson:"email"`
	Phone       string             `bson:"phone"`
	Biograpy    string             `bson:"biograpy"`
	IsPublic    bool               `bson:"isPublic"`
	Education   []Education        `bson:"education"`
	Experience  []Experience       `bson:"experience"`
	Interests   []string           `bson:"interests"`
	Skills      []Skill            `bson:"skills"`
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
	Novice SkillProficiency = iota
	AdvancedBeginner
	Proficient
	Expert
	Master
)

func (proficiency SkillProficiency) String() string {
	switch proficiency {
	case Novice:
		return "novice"
	case AdvancedBeginner:
		return "advanced beginner"
	case Proficient:
		return "proficient"
	case Expert:
		return "expert"
	case Master:
		return "master"
	}

	return "unknown"
}
