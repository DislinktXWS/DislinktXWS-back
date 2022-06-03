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

func mapSkillFromOffer(skillPb *pb.Skill) *domain.SkillDTO {
	newSkill := &domain.SkillDTO{
		OfferId:     skillPb.OfferId,
		Name:        skillPb.Name,
		Proficiency: mapProficiency(skillPb.Proficiency),
	}
	return newSkill
}

func mapProficiency(proficiency pb.Skill_SkillProficiency) domain.SkillProficiency {
	switch proficiency {
	case pb.Skill_novice:
		return domain.Novice
	case pb.Skill_advancedBeginner:
		return domain.AdvancedBeginner
	case pb.Skill_proficient:
		return domain.Proficient
	case pb.Skill_expert:
		return domain.Expert
	case pb.Skill_master:
		return domain.Master
	}
	return domain.Novice
}

func mapBusinessOffer(offer *domain.BusinessOffer, skills []*domain.Skill) *pb.GetBusinessOffer {
	offerPb := &pb.GetBusinessOffer{
		Id:          offer.Id,
		AuthorId:    offer.AuthorId,
		Name:        offer.Name,
		Position:    offer.Position,
		Description: offer.Description,
		Industry:    offer.Industry,
		Skills:      mapSkillsForOffer(skills),
	}

	return offerPb
}

func mapSkillsForOffer(skills []*domain.Skill) []*pb.GetSkill {
	var skillsPb []*pb.GetSkill
	for _, skill := range skills {
		skillPb := &pb.GetSkill{
			Id:          skill.Id,
			Name:        skill.Name,
			Proficiency: mapProficiencyBackwards(skill.Proficiency),
		}
		skillsPb = append(skillsPb, skillPb)

	}
	return skillsPb
}

func mapProficiencyBackwards(proficiency domain.SkillProficiency) pb.GetSkill_SkillProficiency {
	switch proficiency {
	case domain.Novice:
		return pb.GetSkill_novice
	case domain.AdvancedBeginner:
		return pb.GetSkill_advancedBeginner
	case domain.Proficient:
		return pb.GetSkill_proficient
	case domain.Expert:
		return pb.GetSkill_expert
	case domain.Master:
		return pb.GetSkill_master
	}
	return pb.GetSkill_novice
}
