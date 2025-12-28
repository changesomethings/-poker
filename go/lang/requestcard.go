package main

import "github.com/gin-gonic/gin"
import "fmt"

type DataCard struct {
	Card1 int `json:"hand1"`
	Card2 int `json:"hand2"`
}
func requestcard (){
 gins:=gin.Default()
 gins.POST("/message",func(c*gin.Context){
    var data DataCard
	
    err:= c.ShouldBindJSON(&data)
	if err!=nil{
		fmt.Println("wrong",err)
		return 
	}
	fmt.Println("收到",data.Card1,data.Card2)
	cal:=calculator(data)//计算胜率
	fmt.Println("胜率",cal)

  c.JSON(200,gin.H{
    "message":"success",
	"cal":cal,
  })

 })
  gins.Run(":8080")
  
}