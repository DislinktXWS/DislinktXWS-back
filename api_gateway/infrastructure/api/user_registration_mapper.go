package api

import (
	"module/api_gateway/domain"
	pbPost "module/common/proto/post_service"
	pbUser "module/common/proto/user_service"
	domainP "module/post_service/domain"
)

func mapToUserPb(user *domain.UserRegistration) *pbUser.User {
	userPb := &pbUser.User{
		Id:          user.Id,
		Name:        user.Name,
		Surname:     user.Surname,
		Username:    user.Username,
		DateOfBirth: user.DateOfBirth,
		Gender:      user.Gender,
		Email:       user.Email,
		Phone:       user.Phone,
		Biography:   user.Biography,
	}
	return userPb
}

func mapNewPost(postPb *pbPost.Post) *domainP.Post {
	post := &domainP.Post{
		Content:  postPb.Content,
		Image:    postPb.Image,
		Date:     postPb.Date,
		User:     postPb.User,
		Likes:    postPb.Likes,
		Dislikes: postPb.Dislikes,
	}
	return post
}
