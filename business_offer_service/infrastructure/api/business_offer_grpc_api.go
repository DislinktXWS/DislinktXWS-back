package api

import (
	"context"
	"github.com/dislinktxws-back/business_offer_service/application"
	"github.com/dislinktxws-back/business_offer_service/tracer"
	pb "github.com/dislinktxws-back/common/proto/business_offer_service"
	otgo "github.com/opentracing/opentracing-go"
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
	trace       otgo.Tracer
)

func NewBusinessOfferHandler(service *application.BusinessOfferService) *BusinessOfferHandler {
	return &BusinessOfferHandler{
		service: service,
	}
}

func init() {
	trace, _ = tracer.Init("business-offer-service")
	otgo.SetGlobalTracer(trace)
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
	span := tracer.StartSpanFromContextMetadata(ctx, "InsertBusinessOffer")
	defer span.Finish()
	businessOffer := mapNewBusinessOffer(request.Offer)
	err, offerId := handler.service.InsertBusinessOffer(businessOffer, ctx)
	if err != nil {
		ErrorLogger.Println("Action: 25, Message: Cannot create business offer!")
		log.Println("Action: 25, Message: Cannot create business offer!")
		return nil, err
	}
	InfoLogger.Println("Action: 26, Message: User with id " + request.Offer.AuthorId + " created new business offer.")
	log.Println("Action: 26, Message: User with id " + request.Offer.AuthorId + " created new business offer.")
	return &pb.InsertOfferResponse{Id: offerId}, nil
}

func (handler *BusinessOfferHandler) InsertSkill(ctx context.Context, request *pb.InsertSkillsRequest) (*pb.InsertSkillsResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "InsertSKill")
	defer span.Finish()
	skill := mapSkillFromOffer(request.Skill)
	err := handler.service.InsertSkill(skill, ctx)
	if err != nil {
		ErrorLogger.Println("Action: 25, Message: Cannot create new skill!")
		log.Println("Action: 25, Message: Cannot create new skill!")
		return nil, err
	}
	InfoLogger.Println("Action: 27, Message: Created new skill for business offer.")
	log.Println("Action: 27, Message: Created new skill for business offer.")
	return &pb.InsertSkillsResponse{}, nil
}

func (handler *BusinessOfferHandler) GetBusinessOffers(ctx context.Context, request *pb.GetAllOffersRequest) (*pb.GetAllOffersResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "GetBusinessOffers")
	defer span.Finish()
	offers, err := handler.service.GetBusinessOffers(ctx)
	if err != nil {
		ErrorLogger.Println("Action: 2, Message: Business offers not found")
		log.Println("Action: 2, Message: Business offers not found")
		return nil, err
	}
	response := &pb.GetAllOffersResponse{
		Offers: []*pb.GetBusinessOffer{},
	}

	for _, offer := range offers {
		skills, err1 := handler.service.GetOfferSkills(offer.Id, ctx)
		if err1 != nil {
			return nil, err1
		}
		o := mapBusinessOffer(offer, skills)
		response.Offers = append(response.Offers, o)
	}

	return response, nil
}

func (handler *BusinessOfferHandler) GetBusinessOfferRecommendations(ctx context.Context, request *pb.RecommendationsRequest) (*pb.RecommendationsResponse, error) {
	recommend := mapRecommendation(request.Recommend)
	recommendations, _ := handler.service.GetBusinessOfferRecommendations(recommend)
	response := &pb.RecommendationsResponse{
		Offers: []*pb.GetBusinessOffer{},
	}

	for _, recommendation := range recommendations {
		skills, err1 := handler.service.GetOfferSkills(recommendation.Id)
		if err1 != nil {
			return nil, err1
		}
		o := mapBusinessOffer(recommendation, skills)
		response.Offers = append(response.Offers, o)
	}

	return response, nil
}
