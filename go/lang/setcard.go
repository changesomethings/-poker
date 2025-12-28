package main
import "fmt"
type Player struct{
	   Person int
	     Hand  [][]int
		   Point [][]int
		     number [][]int
			   color [][]int
			     max [][]int
 }

func calculator(data DataCard) float32{
   cards:=Card();
   	cards = know (cards,data.Card1,data.Card2);//删掉已知牌
	hands:=[]int{data.Card1,data.Card2}//自己加已知公牌
     fmt.Println(hands)
	win:=0
    fre:=100000//循环次数
	fre32:=float32(fre)//变浮点型
	for i:=0;i<fre;i++{
	 opp:=1 //对手数量
	play:=&Player{Person:opp,Hand:make([][]int,opp),
	number:make([][]int,opp),color:make([][]int,opp),
	Point:make([][]int,opp),max:make([][]int,opp)}
   mypoint:=deal(hands,cards,play)
   win+=pk(mypoint,play)
 }
 win32:=float32(win)
 cal:=win32/fre32
 return cal 
}