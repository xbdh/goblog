package main

import (
	"goblog/model"
	"goblog/router"
)

func main (){
	model.InitDb()
	router.InitRouter()
}