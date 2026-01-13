package model 
type Dealer interface{
	NewDeck()(*Deck,Decker)
}
type Decker interface{
	ShuffleCard(d*Deck)
    Dealing(need int,d*Deck)[]int
}
