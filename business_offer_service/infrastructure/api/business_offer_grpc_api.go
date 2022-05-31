package api

import (
	"github.com/dislinktxws-back/business_offer_service/application"
	pb "github.com/dislinktxws-back/common/proto/business_offer_service"
)

type BusinessOfferHandler struct {
	pb.UnimplementedBusinessOffersServiceServer
	service *application.BusinessOfferService
}

func NewBusinessOfferHandler(service *application.BusinessOfferService) *BusinessOfferHandler {
	return &BusinessOfferHandler{
		service: service,
	}
}
