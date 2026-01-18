package controller

import (
	"poker/model"
	"poker/service"
	"fmt"
)
func Begining(){
	win:=0
	beginner:=&model.Begin{}
	beginner.Hand=[]int{30,59}
	beginner.PublicCard=[]int{18,15,58}
	beginner.Person=6
    beginner.Frequency=10000
	result:=func(s int){
		win+=s
	}
	for i:=0;i<beginner.Frequency;i++{
		service.GameMain(beginner,result)
	}
	fmt.Println((float32(win)/float32(beginner.Frequency))*100)
}