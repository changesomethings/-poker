package main;
func bitwise(handcopy[] int) ([]int,[]int){
    //取回数字和花色 
      bitnumber:=make([]int,7);
	  bitcolor:=make([]int,7);
	for i:=0;i<len(handcopy);i++{
       bitnumber[i]=handcopy[i]>>2;
       bitcolor[i]=handcopy[i]&3; 
	} 
   return bitnumber,bitcolor
	}