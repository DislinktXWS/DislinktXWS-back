package domain

import "time"

type UserRegistration struct {
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
	VerificationToken string
}

type UserBasicInfo struct {
	Id       string
	Name     string
	Surname  string
	Username string
}

type ConversationInfo struct {
	UserId                 string
	Username               string
	FirstName              string
	LastName               string
	LastMessageDate        time.Time
	LastMessageDateString  string
	NumberOfUnreadMessages int32
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
