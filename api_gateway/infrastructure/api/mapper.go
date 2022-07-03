package api

import (
	"github.com/dislinktxws-back/api_gateway/domain"
	pbAuth "github.com/dislinktxws-back/common/proto/authentication_service"
	pbOffer "github.com/dislinktxws-back/common/proto/business_offer_service"
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
		Id:       postPb.Id,
		Content:  postPb.Content,
		Image:    postPb.Image,
		Date:     postPb.Date,
		User:     postPb.User,
		Likes:    postPb.Likes,
		Dislikes: postPb.Dislikes,
	}
	for _, comment := range postPb.Comments {
		post.Comments = append(post.Comments, domain.Comment{
			PostId:  comment.PostId,
			User:    comment.User,
			Content: comment.Content,
		})
	}

	return post
}

func mapToOfferPb(offer *domain.BusinessOffer) *pbOffer.BusinessOffer {
	offerPb := &pbOffer.BusinessOffer{
		Name:        offer.Name,
		AuthorId:    offer.AuthorId,
		Position:    offer.Position,
		Description: offer.Description,
		Industry:    offer.Industry,
	}
	return offerPb
}

func mapToSkillPb(offer *domain.Skill) *pbOffer.Skill {
	offerPb := &pbOffer.Skill{
		Name:        offer.Name,
		Proficiency: pbOffer.Skill_SkillProficiency(offer.Proficiency),
	}
	return offerPb
}

func mapToAuthPb(user *domain.UserRegistration) *pbAuth.Auth {
	authPb := &pbAuth.Auth{
		Id:                user.Id,
		Username:          user.Username,
		VerificationToken: user.VerificationToken,
		Email:             user.Email,
	}
	return authPb
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
