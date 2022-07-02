package persistence

import (
	"fmt"
	"github.com/dislinktxws-back/business_offer_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

type BusinessOffersDBGraph struct {
	session *neo4j.Session
}

func NewBusinessOffersGraph(session *neo4j.Session) domain.BusinessOffersGraph {
	return &BusinessOffersDBGraph{
		session: session,
	}
}

func (store *BusinessOffersDBGraph) InsertBusinessOffer(offer *domain.BusinessOffer) (error, int64) {
	var session = *store.session
	//_, offerId := addOfferNode(offer.AuthorId, offer.Name, offer.Position, offer.Description, offer.Industry)
	//fmt.Println(offerId)
	//_, err := session.WriteTransaction(work)
	offerId, err := addOffer(session, offer.AuthorId, offer.Name, offer.Position, offer.Description, offer.Industry)
	offers, _ := getAllOffers(session)
	for _, offer := range offers {
		fmt.Println(*offer)
	}
	fmt.Println(offerId)
	return err, offerId
}

func (store *BusinessOffersDBGraph) InsertSkill(skill *domain.SkillDTO) error {
	var session = *store.session
	_, err := session.WriteTransaction(addSkillNode(skill.OfferId, skill.Name, skill.Proficiency.String()))
	skills, _ := getOfferSkills(session, skill.OfferId)
	fmt.Println("Offer skills:")
	for _, s := range skills {
		fmt.Println(*s)
	}
	return err
}

func (store *BusinessOffersDBGraph) GetBusinessOffers() ([]*domain.BusinessOffer, error) {
	var session = *store.session
	offers, err := getAllOffers(session)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return offers, nil
}

func (store *BusinessOffersDBGraph) GetOfferSkills(offerId int64) ([]*domain.Skill, error) {
	var session = *store.session
	skills, err := getOfferSkills(session, offerId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return skills, err
}
