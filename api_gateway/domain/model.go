package domain

type UserRegistration struct {
	Id          string
	Name        string
	Surname     string
	Username    string
	Password    string
	DateOfBirth string
	Gender      string
	Email       string
	Phone       string
	Biography   string
}

type UserBasicInfo struct {
	Id       string
	Name     string
	Surname  string
	Username string
}
