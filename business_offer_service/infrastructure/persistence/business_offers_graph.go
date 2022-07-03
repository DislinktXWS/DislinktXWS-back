package persistence

import (
	"context"
	"fmt"
	"github.com/dislinktxws-back/business_offer_service/domain"
	"github.com/dislinktxws-back/business_offer_service/tracer"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type BusinessOffersDBGraph struct {
	session *neo4j.Session
}

func (store *BusinessOffersDBGraph) GetBusinessOfferRecommendations(recommend *domain.Recommend) ([]*domain.BusinessOffer, error) {
	var session = *store.session
	offers := []*domain.BusinessOffer{}
	for _, skill := range recommend.Skills {
		newOffers := getOffersBySkill(session, skill)
		for _, offer := range newOffers {
			offers = append(offers, offer)
		}
	}
	for _, industry := range recommend.Industry {
		newOffers := getOffersByIndustry(session, industry)
		for _, offer := range newOffers {
			offers = append(offers, offer)
		}
	}
	return offers, nil
}

func NewBusinessOffersGraph(session *neo4j.Session) domain.BusinessOffersGraph {
	return &BusinessOffersDBGraph{
		session: session,
	}
}

func (store *BusinessOffersDBGraph) InsertBusinessOffer(offer *domain.BusinessOffer, ctx context.Context) (error, int64) {
	span := tracer.StartSpanFromContext(ctx, "InsertBusinessOffer")
	defer span.Finish()
	var session = *store.session
	offerId, err := addOffer(session, offer.AuthorId, offer.Name, offer.Position, offer.Description, offer.Industry)
	offers, _ := getAllOffers(session)
	for _, offer := range offers {
		fmt.Println(*offer)
	}
	fmt.Println(offerId)
	return err, offerId
}

func (store *BusinessOffersDBGraph) InsertSkill(skill *domain.SkillDTO, ctx context.Context) error {
	span := tracer.StartSpanFromContext(ctx, "InsertSkill")
	defer span.Finish()
	var session = *store.session
	_, err := session.WriteTransaction(addSkillNode(skill.OfferId, skill.Name, skill.Proficiency.String()))
	skills, _ := getOfferSkills(session, skill.OfferId)
	fmt.Println("Offer skills:")
	for _, s := range skills {
		fmt.Println(*s)
	}
	return err
}

func (store *BusinessOffersDBGraph) GetBusinessOffers(ctx context.Context) ([]*domain.BusinessOffer, error) {
	span := tracer.StartSpanFromContext(ctx, "GetBusinessOffers")
	defer span.Finish()
	var session = *store.session
	offers, err := getAllOffers(session)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return offers, nil
}

func (store *BusinessOffersDBGraph) GetOfferSkills(offerId int64, ctx context.Context) ([]*domain.Skill, error) {
	span := tracer.StartSpanFromContext(ctx, "GetOfferSkills")
	defer span.Finish()
	var session = *store.session
	skills, err := getOfferSkills(session, offerId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return skills, err
}
