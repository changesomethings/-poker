package main;

type Deal struct{
	Colorpoint int
	ColorMax int
	Points int 
	Maxs int 
	fullhouse []int 
	threekind int
	Straight int
	straightMax int 
}
func dealnumber(number[]int)(Deal){
	
     d:=Deal{};
    straights:=straight(number,d)//去判断顺子 逻辑在下面那个函数
 point:=0;
 max:=0;
 three:=false;
 threeValue:=0;
 two:=false;
 twoValue:=0
 pairs:=0;
 fullhouse:=make([]int,2);
 count:=[15]int{};
 for _,v:=range number{
	count[v]++;
 }
 for i:=14;i>=2;i--{
    if (count[i]==4){	
     point =7;
	 max=i;
   return Deal{
	Points:point,
	Maxs:max,
   } 
 } else if(count[i]==3){
	  three=true;
	  max=i;
	  threeValue=i;
 }else if(count[i]==2){
    two=true; 
	max=i
	if max>twoValue{
		twoValue=max
	}
	pairs++
}  else if count[i]==1{
	point=0
	if i>max{
		max=i
	}
	   } 	
}
  if(three==true&&two==true){
	 point=6;
	 fullhouse[0]=threeValue;
	 fullhouse[1]=twoValue;
	return Deal{
	  Points:point,
	  fullhouse:fullhouse,
	}
  } else if straights.Straight ==4{//判断顺子
	 return Deal{
		Points:straights.Straight,
		Maxs:straights.straightMax,
	 }
  }else if three==true&&two==false{
     point=3
	 return Deal{
		Points:point,
		Maxs:max,
	 }
  } else if three==false&&two==true{
	 if pairs>=2{
       max=twoValue
	   point=2
	   return Deal{
		Points:point,
		Maxs:max,
	   }
}else  if pairs==1{
		point=1
		max=twoValue
		return Deal{
			Points:point,
			Maxs:max,
		}
	 }
  } else {
	point=0
	return Deal{
		Points:point,
		Maxs:max,
	}
  }

  return Deal{
	 Points:point,
	 Maxs:max,
  }
}
func straight(number[]int,deal Deal)Deal{
	 var  mask uint32=0
  for _,v:=range number{
     value:=v-1
	 mask=mask|(1<<uint32(value))
  }
  if(mask>>13)&1==1{
	mask|=(1<<0)
  }
   s:=mask
   s&=(s<<1)
   s&=(s<<1)
   s&=(s<<1)
   s&=(s<<1)
   if s!=0{
	 deal.Straight=4
     var minkind int
	 for i:=13;i>0;i--{
		if (s>>uint(i))&1==1{
			minkind =i
			break
		}
       deal.straightMax=minkind+5
	 }
  return deal
   }
  return deal
}