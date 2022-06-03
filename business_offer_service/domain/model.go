package domain

type BusinessOffer struct {
	Id          int
	AuthorId    string
	Name        string
	Position    string
	Description string
	Industry    string
}

type Skill struct {
	Id          int
	OfferId     int
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
