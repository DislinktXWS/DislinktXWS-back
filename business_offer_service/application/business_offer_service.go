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

func (service *BusinessOfferService) InsertBusinessOffer(offer *domain.BusinessOffer) error {
	return service.graph.InsertBusinessOffer(offer)
}
