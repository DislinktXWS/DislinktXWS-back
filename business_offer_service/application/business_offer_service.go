package application

import "github.com/dislinktxws-back/business_offer_service/domain"

type BusinessOfferService struct {
	graph domain.BusinessOffersGraph
}

func NewBusinessOfferService(graph domain.BusinessOffersGraph) *BusinessOfferService {
	return &BusinessOfferService{
		graph: graph,
	}
}

func (service *BusinessOfferService) InsertBusinessOffer(offer *domain.BusinessOffer) (error, int64) {
	return service.graph.InsertBusinessOffer(offer)
}

func (service *BusinessOfferService) InsertSkill(skill *domain.SkillDTO) error {
	return service.graph.InsertSkill(skill)
}

func (service *BusinessOfferService) GetBusinessOffers() ([]*domain.BusinessOffer, error) {
	return service.graph.GetBusinessOffers()
}

func (service *BusinessOfferService) GetOfferSkills(offerId int64) ([]*domain.Skill, error) {
	return service.graph.GetOfferSkills(offerId)
}
