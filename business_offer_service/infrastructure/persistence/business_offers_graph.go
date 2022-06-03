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

func (store *BusinessOffersDBGraph) InsertBusinessOffer(offer *domain.BusinessOffer) error {
	var session = *store.session
	_, err := session.WriteTransaction(addOfferNode(offer.AuthorId, offer.Name, offer.Position, offer.Description, offer.Industry))
	//fmt.Println(addOffer(session, offer.AuthorId, offer.Name, offer.Position, offer.Description, offer.Industry))
	//fmt.Println(getAllOffers(session))
	offers, _ := getAllOffers(session)
	for _, offer := range offers {
		fmt.Println(offer)
	}
	return err
}
