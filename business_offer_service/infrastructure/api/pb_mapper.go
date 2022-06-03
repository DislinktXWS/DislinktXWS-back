package api

import (
	"fmt"
	"github.com/dislinktxws-back/business_offer_service/domain"
	pb "github.com/dislinktxws-back/common/proto/business_offer_service"
)

func mapNewBusinessOffer(offerPb *pb.BusinessOffer) *domain.BusinessOffer {
	businessOffer := &domain.BusinessOffer{
		Id:          0,
		AuthorId:    offerPb.AuthorId,
		Name:        offerPb.Name,
		Position:    offerPb.Position,
		Description: offerPb.Description,
		Industry:    offerPb.Industry,
	}

	fmt.Println(businessOffer)
	return businessOffer
}

func mapSkillFromOffer(skillPb *pb.Skill) *domain.Skill {
	newSkill := &domain.Skill{
		OfferId:     0,
		Name:        "",
		Proficiency: 0,
	}
	return newSkill
}
