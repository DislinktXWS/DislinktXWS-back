package persistence

import (
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
