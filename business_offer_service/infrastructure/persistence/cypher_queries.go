package persistence

import (
	"fmt"
	"github.com/dislinktxws-back/business_offer_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func addOfferNode(authorId, name, position, description, industry string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run(
			"CREATE (businessOffer:BUSINESSOFFER {authorId: $authorId, name: $name, position: $position, description: $description, industry: $industry})",
			map[string]interface{}{"authorId": authorId,
				"name": name, "position": position,
				"description": description, "industry": industry})

		if err != nil {
			return nil, err
		}
		fmt.Println("NEMA GRESKE")
		return result.Consume()
	}
}

/*func addOffer(session neo4j.Session, authorId, name, position, description, industry string) (domain.BusinessOffer, error) {
	offer, _ := session.WriteTransaction(func(transaction neo4j.Transaction) (interface{}, error) {
		fmt.Println("OK1")
		result, err := transaction.Run("CREATE (businessOffer:BUSINESSOFFER {authorId: $authorId, name: $name, position: $position, description: $description, industry: $industry})"+
			"RETURN businessOffer.id, businessOffer.author, businessOffer.name, businessOffer.position, businessOffer.description, businessOffer.industry",
			map[string]interface{}{"authorId": authorId,
				"name": name, "position": position,
				"description": description, "industry": industry})
		fmt.Println("OK2")

		if err != nil {
			fmt.Println("NEOK1")
			return nil, err
		}

		if result.Next() {
			offerOk := domain.BusinessOffer{
				Id:          result.Record().Values[0].(int),
				AuthorId:    result.Record().Values[1].(string),
				Name:        result.Record().Values[2].(string),
				Position:    result.Record().Values[3].(string),
				Description: result.Record().Values[4].(string),
				Industry:    result.Record().Values[5].(string),
			}
			fmt.Println("OK3")
			return offerOk, nil
		}
		fmt.Println("OK4")
		return nil, result.Err()
	})

	return offerOk, nil
}*/

func addSkillNode(offerId int, name, proficiency string) neo4j.TransactionWork {
	return func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run("CREATE (skill:SKILL {name: $name, proficiency: $proficiency}) <-[:HAS_SKILL]- (:BUSINESSOFFER {id: $offerId})",
			map[string]interface{}{"offerId": offerId, "name": name, "proficiency": proficiency})
		if err != nil {
			return nil, err
		}
		fmt.Println("NEMA GRESKE SKILL")
		return result.Consume()
	}
}

func getAllOffers(session neo4j.Session) (offers []*domain.BusinessOffer, err1 error) {
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (offer:BUSINESSOFFER) RETURN offer.authorId, offer.name, offer.position, offer.description, offer.industry", map[string]interface{}{})

		for records.Next() {
			offer := domain.BusinessOffer{
				Id:          0,
				AuthorId:    records.Record().Values[0].(string),
				Name:        records.Record().Values[1].(string),
				Position:    records.Record().Values[2].(string),
				Description: records.Record().Values[3].(string),
				Industry:    records.Record().Values[4].(string),
			}
			offers = append(offers, &offer)
		}

		if err != nil {
			return nil, err
		}

		return offers, nil
	})
	if err != nil {
		return nil, err
	}
	return offers, nil

}
