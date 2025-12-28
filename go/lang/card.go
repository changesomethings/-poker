package main;

func Card() []int{
	colors:=[]int{0,1,2,3};
	numbers:=make([]int,13);
	for i:=2;i<15;i++{
		numbers[i-2]=i;
	}

	cards:=make([]int,52);
	index:=0;
	for i:=0;i<len(numbers);i++{
		for j:=0;j<len(colors);j++{
			cards[index]=(numbers[i]<<2)|colors[j];
			index++;
		}
	}
	return cards;
	
}

func know( cards []int, card ...int)[] int{
   for _,card :=range card{
   for i,v:=range cards{
	 if v==card{
       cards[i]=cards[len(cards)-1]
	   cards=cards[:len(cards)-1]
	   break;
	 }  
   }
}
   return cards;
   
}

