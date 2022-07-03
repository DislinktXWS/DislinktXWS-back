package application

import (
	"context"
	"github.com/dislinktxws-back/business_offer_service/domain"
)

type BusinessOfferService struct {
	graph domain.BusinessOffersGraph
}

func NewBusinessOfferService(graph domain.BusinessOffersGraph) *BusinessOfferService {
	return &BusinessOfferService{
		graph: graph,
	}
}

func (service *BusinessOfferService) InsertBusinessOffer(offer *domain.BusinessOffer, ctx context.Context) (error, int64) {
	return service.graph.InsertBusinessOffer(offer, ctx)
}

func (service *BusinessOfferService) InsertSkill(skill *domain.SkillDTO, ctx context.Context) error {
	return service.graph.InsertSkill(skill, ctx)
}

func (service *BusinessOfferService) GetBusinessOffers(ctx context.Context) ([]*domain.BusinessOffer, error) {
	return service.graph.GetBusinessOffers(ctx)
}

func (service *BusinessOfferService) GetOfferSkills(offerId int64, ctx context.Context) ([]*domain.Skill, error) {
	return service.graph.GetOfferSkills(offerId, ctx)
}

func (service *BusinessOfferService) GetBusinessOfferRecommendations(recommend *domain.Recommend) ([]*domain.BusinessOffer, error) {
	return service.graph.GetBusinessOfferRecommendations(recommend)
}
