package api

import (
	"github.com/dislinktxws-back/api_gateway/domain"
	pbMessage "github.com/dislinktxws-back/common/proto/message_service"
	pbPost "github.com/dislinktxws-back/common/proto/post_service"
	pbUser "github.com/dislinktxws-back/common/proto/user_service"
)

func mapToUserPb(user *domain.UserRegistration) *pbUser.User {
	userPb := &pbUser.User{
		Id:                user.Id,
		Name:              user.Name,
		Surname:           user.Surname,
		Username:          user.Username,
		DateOfBirth:       user.DateOfBirth,
		Gender:            user.Gender,
		Email:             user.Email,
		Phone:             user.Phone,
		Biography:         user.Biography,
		VerificationToken: user.VerificationToken,
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

func mapNewConversationInfo(postPb *pbMessage.GetConversationResponse, userInfo *domain.UserBasicInfo) *domain.ConversationInfo {
	info := &domain.ConversationInfo{
		UserId:    userInfo.Id,
		Username:  userInfo.Username,
		FirstName: userInfo.Name,
		LastName:  userInfo.Surname,
	}
	return info
}
