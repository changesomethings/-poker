package utils

import (
	"math/rand/v2"
	"poker/model"
)
func (T*TexasDealer)ShuffleCard(d*model.Deck){
	rand.Shuffle(len(d.AllCards),func(i,j int){
		d.AllCards[i],d.AllCards[j]=d.AllCards[j],d.AllCards[i]
	})
}