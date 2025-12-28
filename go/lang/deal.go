package main

func deal(hand []int ,cards []int,play*Player)Deal{
	
	
     myhand:=set(cards,hand,play)
     
	 number,color:=bitwise(myhand)
	 colordeal:=dealcolor(color,number)
	 numberdeal:=dealnumber(number)
	 
	 points:=compt(colordeal,numberdeal)
     
	 
   
	for i:=0;i<play.Person;i++{
		 opphands:=[]int{} 
		 opphands=append(opphands,play.Hand[i]...)//把二维结构体每个对手抽出来
	     oppnumber,oppcolor:=bitwise(opphands)
		 play.color[i]=oppcolor
		 play.number[i]=oppnumber
		 dealcolor:=dealcolor(oppcolor,oppnumber)
		 dealnumber:=dealnumber(oppnumber)
		 point:=compt(dealcolor,dealnumber)
		 
          play.Point[i] = []int{point.Points}
          play.max[i]=[]int {point.Maxs}
        
	}
    
   return points
}