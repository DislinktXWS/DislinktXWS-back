package api

import (
	"context"
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
	businessOffer := mapNewBusinessOffer(request.Offer)
	err, offerId := handler.service.InsertBusinessOffer(businessOffer)
	if err != nil {
		return nil, err
	}
	return &pb.InsertOfferResponse{Id: offerId}, nil
}

func (handler *BusinessOfferHandler) InsertSkill(ctx context.Context, request *pb.InsertSkillsRequest) (*pb.InsertSkillsResponse, error) {
	skill := mapSkillFromOffer(request.Skill)
	err := handler.service.InsertSkill(skill)
	if err != nil {
		return nil, err
	}
	return &pb.InsertSkillsResponse{}, nil
}

func (handler *BusinessOfferHandler) GetBusinessOffers(ctx context.Context, request *pb.GetAllOffersRequest) (*pb.GetAllOffersResponse, error) {
	offers, err := handler.service.GetBusinessOffers()
	if err != nil {
		return nil, err
	}
	response := &pb.GetAllOffersResponse{
		Offers: []*pb.GetBusinessOffer{},
	}

	for _, offer := range offers {
		skills, err1 := handler.service.GetOfferSkills(offer.Id)
		if err1 != nil {
			return nil, err1
		}
		o := mapBusinessOffer(offer, skills)
		response.Offers = append(response.Offers, o)
	}

	return response, nil
}
