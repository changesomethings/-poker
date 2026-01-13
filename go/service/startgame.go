package service
import"fmt"
import "poker/model"
func GameStart(){
	game:=GameFactory("texas")
	_=game.MakeJudge()
    dealer:=game.MakeDealer()
     _,_=dealer.NewDeck()//T是发牌员 d 是deck的属性
	g:=model.NewGameData(6)
	g.Add(0,[]int{15,10})
	player:=g.GetAllPlayer()
	fmt.Println(player)
}