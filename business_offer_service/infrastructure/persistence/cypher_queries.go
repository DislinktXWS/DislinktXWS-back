package persistence

import (
	"fmt"
	"github.com/dislinktxws-back/business_offer_service/domain"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
)

func addOffer(session neo4j.Session, authorId, name, position, description, industry string) (int64, error) {
	var offerId int64
	session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		var result, err = tx.Run(
			"CREATE (businessOffer:BUSINESSOFFER {authorId: $authorId, name: $name, position: $position, description: $description, industry: $industry})"+
				"RETURN ID(businessOffer), businessOffer.name",
			map[string]interface{}{"authorId": authorId,
				"name": name, "position": position,
				"description": description, "industry": industry})

		if err != nil {
			return nil, err
		}
		for result.Next() {
			offerId = result.Record().Values[0].(int64)
			fmt.Println(offerId)
		}
		return offerId, nil
	})
	return offerId, nil
}

func addSkillNode(offerId int64, name, proficiency string) neo4j.TransactionWork {
	fmt.Println("ADD SKILL")
	return func(tx neo4j.Transaction) (interface{}, error) {
		if checkIfSkillExists(name, proficiency, tx) {
			fmt.Println("ADD SKILL EXISTS")
			var result, err = tx.Run("MATCH (skill:SKILL {name:$name}) "+
				"MATCH (offer:BUSINESSOFFER) where ID(offer) = $offerId "+
				"CREATE (offer) -[:HAS_SKILL] -> (skill)",
				map[string]interface{}{"offerId": offerId, "name": name, "proficiency": proficiency})
			if err != nil {
				fmt.Println("ADD SKILL ERROR")
				fmt.Println(err)
				return nil, err
			}
			fmt.Println("NEMA GRESKE SKILL postojeci")
			return result.Consume()
		} else {
			var result, err = tx.Run("MATCH(offer:BUSINESSOFFER) WHERE ID(offer) = $offerId CREATE (skill:SKILL {name: $name, proficiency: $proficiency}) <-[:HAS_SKILL]- (offer)",
				map[string]interface{}{"offerId": offerId, "name": name, "proficiency": proficiency})
			if err != nil {
				fmt.Println("OVDE PUKNE")
				fmt.Println(err)
				return nil, err
			}
			fmt.Println("NEMA GRESKE SKILL")
			return result.Consume()
		}
	}
}

func getAllOffers(session neo4j.Session) (offers []*domain.BusinessOffer, err1 error) {
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (offer:BUSINESSOFFER) RETURN ID(offer),offer.authorId, offer.name, offer.position, offer.description, offer.industry", map[string]interface{}{})

		for records.Next() {
			offer := domain.BusinessOffer{
				Id:          records.Record().Values[0].(int64),
				AuthorId:    records.Record().Values[1].(string),
				Name:        records.Record().Values[2].(string),
				Position:    records.Record().Values[3].(string),
				Description: records.Record().Values[4].(string),
				Industry:    records.Record().Values[5].(string),
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

func getAuthorOffers(session neo4j.Session, authorId string) (offers []*domain.BusinessOffer, err1 error) {
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (offer:BUSINESSOFFER) "+
			"WHERE offer.authorId = $authorId"+
			"RETURN offer.authorId, offer.name, offer.position, offer.description, offer.industry", map[string]interface{}{
			"authorId": authorId,
		})

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

func getOffersByIndustry(session neo4j.Session, industry string) (offers []*domain.BusinessOffer) {
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (offer: BUSINESSOFFER) "+
			"WHERE offer.industry = $industry "+
			"RETURN ID(offer), offer.authorId, offer.name, offer.position, offer.description, offer.industry", map[string]interface{}{
			"industry": industry,
		})

		for records.Next() {
			offer := domain.BusinessOffer{
				Id:          records.Record().Values[0].(int64),
				AuthorId:    records.Record().Values[1].(string),
				Name:        records.Record().Values[2].(string),
				Position:    records.Record().Values[3].(string),
				Description: records.Record().Values[4].(string),
				Industry:    records.Record().Values[5].(string),
			}
			offers = append(offers, &offer)
		}
		if err != nil {
			fmt.Sprintf(err.Error())
			return nil, err
		}

		return offers, nil
	})

	if err != nil {
		return nil
	}
	return offers

}

func getOffersBySkill(session neo4j.Session, skillName string) (offers []*domain.BusinessOffer) {
	fmt.Println("USLO U GET OFFERS BY SKILL, SKILLNAME:")
	fmt.Println(skillName)
	fmt.Println(&skillName)
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (offer: BUSINESSOFFER) -[HAS_SKILL]-> (skill:SKILL) "+
			"WHERE skill.name = $skillName "+
			"RETURN ID(offer), offer.authorId, offer.name, offer.position, offer.description, offer.industry", map[string]interface{}{
			"skillName": skillName,
		})

		fmt.Println(records)

		for records.Next() {
			offer := domain.BusinessOffer{
				Id:          records.Record().Values[0].(int64),
				AuthorId:    records.Record().Values[1].(string),
				Name:        records.Record().Values[2].(string),
				Position:    records.Record().Values[3].(string),
				Description: records.Record().Values[4].(string),
				Industry:    records.Record().Values[5].(string),
			}
			offers = append(offers, &offer)
		}
		if err != nil {
			fmt.Sprintf(err.Error())
			return nil, err
		}

		return offers, nil
	})

	if err != nil {
		return nil
	}
	return offers

}

func getOfferSkills(session neo4j.Session, offerId int64) (skills []*domain.Skill, err1 error) {
	_, err := session.WriteTransaction(func(tx neo4j.Transaction) (interface{}, error) {
		records, err := tx.Run("MATCH (offer:BUSINESSOFFER) -[:HAS_SKILL] -> (skill:SKILL) "+
			"WHERE ID(offer) = $offerId "+
			"RETURN ID(skill), skill.name, skill.proficiency", map[string]interface{}{
			"offerId": offerId,
		})

		for records.Next() {
			skill := domain.Skill{
				Id:          records.Record().Values[0].(int64),
				Name:        records.Record().Values[1].(string),
				Proficiency: setProficiency(records.Record().Values[2].(string)),
			}
			skills = append(skills, &skill)
		}

		if err != nil {
			return nil, err
		}

		return skills, nil
	})
	if err != nil {
		return nil, err
	}
	return skills, nil

}

func checkIfSkillExists(name, proficiency string, transaction neo4j.Transaction) bool {
	result, _ := transaction.Run("MATCH (skill:SKILL {name: $name, proficiency: $proficiency}) RETURN skill.name, skill.proficiency", map[string]interface{}{"name": name, "proficiency": proficiency})
	if result != nil && result.Next() && result.Record().Values[0].(string) == name && result.Record().Values[1].(string) == proficiency {
		return true
	}
	return false
}

func setProficiency(proficiency string) domain.SkillProficiency {
	if proficiency == "novice" {
		return 0
	} else if proficiency == "advanced beginner" {
		return 1
	} else if proficiency == "proficient" {
		return 2
	} else if proficiency == "expert" {
		return 3
	} else if proficiency == "master" {
		return 4
	}
	return 0
}
