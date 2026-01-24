package main

import (
	"poker/controller"
	"poker/utils"
	
)

func main(){
    
	pool:=[]string{"CA","HA"}
	utils.DealMap(pool)
	
	controller.Route()
} 