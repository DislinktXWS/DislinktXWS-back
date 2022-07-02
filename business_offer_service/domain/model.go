package domain

type BusinessOffer struct {
	Id          int64
	AuthorId    string
	Name        string
	Position    string
	Description string
	Industry    string
}

type Skill struct {
	Id          int64
	Name        string
	Proficiency SkillProficiency
}

type SkillDTO struct {
	OfferId     int64
	Name        string
	Proficiency SkillProficiency
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

type Recommend struct {
	Skills   []string
	Industry []string
}
