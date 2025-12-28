package main

import "math/rand"
//给每个玩家发牌 洗牌 和补公共牌
func set(cards []int, hand []int ,play*Player) []int {
	need := 7 - len(hand) //需要补几张牌
	cardscopy := make([]int, len(cards))
	copy(cardscopy, cards)
	handcopy := make([]int, len(hand))
	copy(handcopy, hand)
	rand.Shuffle(len(cardscopy), func(i, j int) {
		cardscopy[i], cardscopy[j] = cardscopy[j], cardscopy[i]
	})
    for i:=0;i<play.Person;i++{//给每个对手发牌
		start:=i*2
        end:=start+2
		play.Hand[i]=append(play.Hand[i],cardscopy[start:end]...)
	}

	
	if need > 0 {
		handstart:=play.Person*2
		handend:=handstart+need
		handcopy = append(handcopy, cardscopy[handstart:handend]...)
	     
		for i:=0;i<play.Person;i++{
		  play.Hand[i]=append(play.Hand[i],cardscopy[handstart:handend]...)
		}
	    
	}
	if need==0{
		for z:=0;z<play.Person;z++{
			
			play.Hand[z]=append(play.Hand[z],cardscopy[2:]...)
		}
      return handcopy
	}
     return handcopy
}