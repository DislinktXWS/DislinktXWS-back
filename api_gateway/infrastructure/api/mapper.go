package api

import (
	"github.com/dislinktxws-back/api_gateway/domain"
	pbPost "github.com/dislinktxws-back/common/proto/post_service"
	pbUser "github.com/dislinktxws-back/common/proto/user_service"
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

func mapNewPost(postPb *pbPost.Post) *domain.Post {
	post := &domain.Post{
		Content:  postPb.Content,
		Image:    postPb.Image,
		Date:     postPb.Date,
		User:     postPb.User,
		Likes:    postPb.Likes,
		Dislikes: postPb.Dislikes,
	}
	return post
}
