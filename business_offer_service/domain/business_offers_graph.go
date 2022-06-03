package domain

type BusinessOffersGraph interface {
	InsertBusinessOffer(offer *BusinessOffer) error
}
