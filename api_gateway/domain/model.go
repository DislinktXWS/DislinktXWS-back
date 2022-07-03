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

type ApiKey struct {
	Username string
	ApiKey   string
}

type BusinessOfferDto struct {
	Id          int64
	AuthorId    string
	Name        string
	Position    string
	Description string
	Industry    string
	Skills      []Skill
}

type Skill struct {
	Id          int64
	Name        string
	Proficiency SkillProficiency
}

type SkillProficiency int64

type BusinessOffer struct {
	Id          int64
	AuthorId    string
	Name        string
	Position    string
	Description string
	Industry    string
}

type Notification struct {
	Id      string
	From    string //id
	To      string //id
	Date    string
	Content string
}

type NotificationsSettings struct {
	ChatNotificatons         bool
	ConnectionsNotifications bool
	PostNotifications        bool
}

type GetBusinessOffer struct {
	Id          int64
	AuthorId    string
	Name        string
	Position    string
	Description string
	Industry    string
	Skills      []GetSkill
}

type GetSkill struct {
	Id                     int64
	Name                   string
	SkillProficiencyString string
}
