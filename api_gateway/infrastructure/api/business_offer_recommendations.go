package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dislinktxws-back/api_gateway/domain"
	"github.com/dislinktxws-back/api_gateway/infrastructure/services"
	business_offer_proto "github.com/dislinktxws-back/common/proto/business_offer_service"
	user_proto "github.com/dislinktxws-back/common/proto/user_service"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type BusinessOfferRecommendationsHandler struct {
	userClientAddress          string
	businessOfferClientAddress string
}

func (handler BusinessOfferRecommendationsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/offers/offer/getRecommendations/{id}", handler.GetRecommendations)
	if err != nil {
		panic(err)
	}
}

func NewBusinessOfferRecommendationsHandler(userClientAddress string, businessOfferClientAddress string) Handler {
	return &BusinessOfferRecommendationsHandler{
		userClientAddress:          userClientAddress,
		businessOfferClientAddress: businessOfferClientAddress,
	}
}

func (handler *BusinessOfferRecommendationsHandler) GetRecommendations(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	fmt.Println("ID KORISNIKA(custom handler): " + id)
	skills := handler.getUserSkills(id)
	industries := handler.getUserExperience(id)

	recommend := Recommend{
		Skills:   skills,
		Industry: industries,
	}

	recommendations := handler.getBusinessOfferRecommendations(recommend)
	fmt.Println(recommendations)

	response, err := json.Marshal(recommendations)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *BusinessOfferRecommendationsHandler) getUserSkills(userId string) []string {
	userClient := services.NewUserClient(handler.userClientAddress)

	var skills []string
	skillsPb, _ := userClient.GetSkills(context.TODO(), &user_proto.GetSkillsRequest{Id: userId})
	fmt.Println(skillsPb)
	for _, s := range skillsPb.Skills {
		skills = append(skills, s.Name)
	}
	return skills
}

func (handler *BusinessOfferRecommendationsHandler) getUserExperience(userId string) []string {
	userClient := services.NewUserClient(handler.userClientAddress)
	fmt.Println("USLO U GETEXPERIENCE")
	var industries []string
	experiencePb, _ := userClient.GetExperience(context.TODO(), &user_proto.GetExperienceRequest{Id: userId})
	fmt.Println(experiencePb)
	for _, e := range experiencePb.Experience {
		industries = append(industries, e.Industry)
	}
	return industries
}

func (handler *BusinessOfferRecommendationsHandler) getBusinessOfferRecommendations(recommend Recommend) []domain.GetBusinessOffer {
	businessOfferClient := services.NewBusinessOfferClient(handler.businessOfferClientAddress)

	var businessOffers []domain.GetBusinessOffer
	businessOffersPb, _ := businessOfferClient.GetBusinessOfferRecommendations(context.TODO(), &business_offer_proto.RecommendationsRequest{Recommend: mapRecommend(&recommend)})
	for _, offerPb := range businessOffersPb.Offers {
		businessOffers = append(businessOffers, *mapOfferPb(offerPb))
	}
	return businessOffers
}

type Recommend struct {
	Skills   []string
	Industry []string
}

func mapRecommend(recommend *Recommend) *business_offer_proto.Recommend {
	recommendPb := &business_offer_proto.Recommend{
		Skills:   recommend.Skills,
		Industry: recommend.Industry,
	}
	return recommendPb
}

func mapOfferPb(offerPb *business_offer_proto.GetBusinessOffer) *domain.GetBusinessOffer {
	var skills []domain.GetSkill
	for _, s := range offerPb.Skills {
		skill := domain.GetSkill{
			Id:                     s.Id,
			Name:                   s.Name,
			SkillProficiencyString: s.Proficiency.String(),
		}
		skills = append(skills, skill)
	}
	offer := &domain.GetBusinessOffer{
		Id:          offerPb.Id,
		AuthorId:    offerPb.AuthorId,
		Name:        offerPb.Name,
		Position:    offerPb.Position,
		Description: offerPb.Description,
		Industry:    offerPb.Industry,
		Skills:      skills,
	}
	return offer
}
