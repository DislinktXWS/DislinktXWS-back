package api

import (
	"context"
	"github.com/dislinktxws-back/business_offer_service/application"
	pb "github.com/dislinktxws-back/common/proto/business_offer_service"
	"log"
	"os"
)

type BusinessOfferHandler struct {
	pb.UnimplementedBusinessOffersServiceServer
	service *application.BusinessOfferService
}

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

func NewBusinessOfferHandler(service *application.BusinessOfferService) *BusinessOfferHandler {
	return &BusinessOfferHandler{
		service: service,
	}
}

func init() {
	infoFile, err := os.OpenFile("info.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	InfoLogger = log.New(infoFile, "INFO: ", log.LstdFlags|log.Lshortfile)

	errFile, err1 := os.OpenFile("error.log", os.O_APPEND|os.O_WRONLY, 0666)
	if err1 != nil {
		log.Fatal(err1)
	}
	ErrorLogger = log.New(errFile, "ERROR: ", log.LstdFlags|log.Lshortfile)
}

func (handler *BusinessOfferHandler) InsertBusinessOffer(ctx context.Context, request *pb.InsertOfferRequest) (*pb.InsertOfferResponse, error) {
	businessOffer := mapNewBusinessOffer(request.Offer)
	err, offerId := handler.service.InsertBusinessOffer(businessOffer)
	if err != nil {
		ErrorLogger.Println("Cannot create business offer!")
		return nil, err
	}
	InfoLogger.Println("User with id " + request.Offer.AuthorId + " created new business offer.")
	return &pb.InsertOfferResponse{Id: offerId}, nil
}

func (handler *BusinessOfferHandler) InsertSkill(ctx context.Context, request *pb.InsertSkillsRequest) (*pb.InsertSkillsResponse, error) {
	skill := mapSkillFromOffer(request.Skill)
	err := handler.service.InsertSkill(skill)
	if err != nil {
		ErrorLogger.Println("Cannot create new skill!")
		return nil, err
	}
	InfoLogger.Println("Created new skill for business offer.")

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
