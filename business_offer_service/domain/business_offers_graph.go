package domain

import "context"

type BusinessOffersGraph interface {
	InsertBusinessOffer(offer *BusinessOffer, ctx context.Context) (error, int64)
	InsertSkill(skill *SkillDTO, ctx context.Context) error
	GetBusinessOffers(ctx context.Context) ([]*BusinessOffer, error)
	GetOfferSkills(offerId int64, ctx context.Context) ([]*Skill, error)
}
