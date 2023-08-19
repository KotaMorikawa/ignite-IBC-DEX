package types

func NewBuyOrderBook(Amount string, PriceDenom string) BuyOrderBook {
	book := NewOrderBook()
	return BuyOrderBook{
		AmountDenom: Amount,
		PriceDenom:  PriceDenom,
		Book:        &book,
	}
}
