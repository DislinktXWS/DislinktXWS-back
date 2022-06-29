package insert_user

type User struct {
	Id                string
	Name              string
	Surname           string
	Username          string
	Password          string
	DateOfBirth       string
	Gender            string
	Email             string
	Phone             string
	Biography         string
	IsPublic          bool
	VerificationToken string
	Education         []Education
	Experience        []Experience
	Interests         []string
	Skills            []Skill
	ApiKey            string
}

type Education struct {
	School       string
	Degree       string
	FieldOfStudy string
	StartDate    string
	EndDate      string
	Grade        float32
	Description  string
}

type Experience struct {
	Title       string
	CompanyName string
	StartDate   string
	EndDate     string
	Industry    string
	Description string
}

type Skill struct {
	Name        string
	Proficiency SkillProficiency
}

type SkillProficiency int8

type InsertUserCommandType int8

const (
	InsertUserAuthentication InsertUserCommandType = iota
	RollbackInsertUserAuthentication
	InsertUserNode
	UnknownCommand
)

type InsertUserCommand struct {
	User User
	Type InsertUserCommandType
}

type InsertUserReplyType int8

const (
	UserAuthenticationInserted InsertUserReplyType = iota
	UserAuthenticationNotInserted
	UserAuthenticationRolledBack
	UserNodeInserted
	UserNodeNotInserted
	UnknownReply
)

type InsertUserReply struct {
	User User
	Type InsertUserReplyType
}
