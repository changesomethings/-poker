package utils
import "poker/model"
func(T*TexasDealer) Dealing(need int,d*model.Deck)[]int{
    card:=d.AllCards[:need]
	d.SendCard=append(d.SendCard,card...)
	return card
}