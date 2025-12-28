package main;
func dealcolor(color[]int,number[]int) Deal{
count:=make([][]int,4)
 
for i,v:=range color{
	
count[v]=append(count[v],number[i])
}

 for i:=0;i<4;i++{
	if len(count[i])>=5{
		point:=5;
        max:=0;
     for _,v:=range count[i]{
		if v>max{
			max=v
		}
	 }
	 return Deal{
		Colorpoint:point,
		ColorMax:max,
	 }
	}
 }
 return Deal{}
} 