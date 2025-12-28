package main;

func pk (my Deal, play*Player) int{
     // fmt.Printf("%#v\n",my)
     // fmt.Printf("%#v\n",opp)
     z:=0
   for i:=0;i<play.Person;i++{
       
     playpoint:=play.Point[i][0]
     playmax:=play.max[i][0]
   if my.Points>playpoint{
	   z++
     
   } else if my.Points==playpoint{
	 if my.Maxs>playmax{
		z++
	 }else if my.Maxs<playmax{
     break
   }
   }else if my.Points<playpoint {
    break
   }
  
    if z==play.Person{
     return 1
    } 
}
 return 0
}