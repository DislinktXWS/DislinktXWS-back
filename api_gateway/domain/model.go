package domain

type User struct {
	Id          string
	Name        string
	Surname     string
	Username    string
	Password    string
	DateOfBirth string
	Gender      string
	Email       string
	Phone       string
}

type UserBasicInfo struct {
	Id       string
	Name     string
	Surname  string
	Username string
}
