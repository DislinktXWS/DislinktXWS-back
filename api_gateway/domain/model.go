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

type Post struct {
	Id       string
	Content  string
	Image    string
	Date     string
	User     string
	Likes    []string
	Dislikes []string
	Comments []Comment
}

type Comment struct {
	PostId  string
	User    string
	Content string
}
