package domain

type BusinessOffersGraph interface {
	InsertBusinessOffer(offer *BusinessOffer) (error, int64)
	InsertSkill(skill *SkillDTO) error
	GetBusinessOffers() ([]*BusinessOffer, error)
	GetOfferSkills(offerId int64) ([]*Skill, error)
	GetBusinessOfferRecommendations(recommend *Recommend) ([]*BusinessOffer, error)
}
