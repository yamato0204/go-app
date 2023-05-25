package main

import (
	//"fmt"

	
	"go-app/controllers"
	"go-app/models"
	//"log"
	//	"go-app/models"
)



func main(){

	Db := models.Init()

	defer Db.Close()

	

    controllers.StartMainServer()

}