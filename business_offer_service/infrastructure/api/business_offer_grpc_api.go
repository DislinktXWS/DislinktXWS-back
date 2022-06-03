package api

import (
	"context"
	"fmt"
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

func (handler *BusinessOfferHandler) InsertBusinessOffer(ctx context.Context, request *pb.InsertOfferRequest) (*pb.InsertOfferResponse, error) {
	fmt.Println("Uslo u handler")
	businessOffer := mapNewBusinessOffer(request.Offer)
	err := handler.service.InsertBusinessOffer(businessOffer)
	if err != nil {
		return nil, err
	}
	return &pb.InsertOfferResponse{}, nil
}
