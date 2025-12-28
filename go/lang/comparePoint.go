package main;

func compt(deal1 Deal ,deal2 Deal)Deal{
   
  if deal1.Colorpoint==5{
    if deal1.Colorpoint>deal2.Points{//判断是同花还是比同花大的牌
       deal1.Points=5
       deal1.Maxs=deal1.ColorMax
       return deal1
    } 
   return deal2
    }else if deal1.Colorpoint==5&&deal2.Points==4{//同花顺
      if deal2.Maxs==14{//皇家同花顺
        deal2.Points=9
      }else{deal2.Points=8}
      return deal2
    }else{ 
    return deal2
  }
}