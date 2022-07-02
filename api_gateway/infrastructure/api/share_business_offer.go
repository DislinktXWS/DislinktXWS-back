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

type ShareBusinessOfferHandler struct {
	userClientAddress          string
	businessOfferClientAddress string
}

func NewShareBusinessOfferHandler(userClientAddress string, businessOfferClientAddress string) Handler {
	return &ShareBusinessOfferHandler{
		userClientAddress:          userClientAddress,
		businessOfferClientAddress: businessOfferClientAddress,
	}
}

func (handler *ShareBusinessOfferHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/shareBusinessOffer", handler.ShareBusinessOffer)
	if err != nil {
		panic(err)
	}
}

func (handler *ShareBusinessOfferHandler) ShareBusinessOffer(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var businessOfferDto domain.BusinessOfferDto

	err := json.NewDecoder(r.Body).Decode(&businessOfferDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	apiKey := r.Header.Get("ApiKey")
	userId, err := handler.getUserIdByApiKey(apiKey)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var businessOffer domain.BusinessOffer
	businessOffer.Id = 0
	businessOffer.AuthorId = userId
	businessOffer.Name = businessOfferDto.Name
	businessOffer.Position = businessOfferDto.Position
	businessOffer.Description = businessOfferDto.Description
	businessOffer.Industry = businessOfferDto.Industry

	offerId := handler.addOffer(businessOffer)
	fmt.Println("KREIRANA PONUDA")
	fmt.Println(offerId)

	for _, skill := range businessOfferDto.Skills {
		handler.addOfferSkill(skill, offerId)
	}

	response, err := json.Marshal("Success")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *ShareBusinessOfferHandler) getUserIdByApiKey(apiKey string) (string, error) {
	userClient := services.NewUserClient(handler.userClientAddress)
	user, err := userClient.GetByApiKey(context.TODO(), &user_proto.GetByApiKeyRequest{ApiKey: apiKey})
	if err != nil {
		panic(err)
	}
	id := user.User.Id
	return id, err
}

func (handler *ShareBusinessOfferHandler) addOffer(businessOffer domain.BusinessOffer) int64 {
	businessOfferClient := services.NewBusinessOfferClient(handler.businessOfferClientAddress)
	OfferPb := mapToOfferPb(&businessOffer)
	offer, err := businessOfferClient.InsertBusinessOffer(context.TODO(), &business_offer_proto.InsertOfferRequest{Offer: OfferPb})
	if err != nil {
		panic(err)
	}
	id := offer.Id
	return id
}

func (handler *ShareBusinessOfferHandler) addOfferSkill(skill domain.Skill, offerId int64) {
	businessOfferClient := services.NewBusinessOfferClient(handler.businessOfferClientAddress)
	SkillPb := mapToSkillPb(&skill)
	SkillPb.OfferId = offerId
	_, err := businessOfferClient.InsertSkill(context.TODO(), &business_offer_proto.InsertSkillsRequest{Skill: SkillPb})
	if err != nil {
		panic(err)
	}
}
