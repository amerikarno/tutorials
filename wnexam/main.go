package main

import (
	"wnexam/controller"
	"wnexam/repository"
	"wnexam/service"

	"github.com/gin-gonic/gin"
	// "github.com/rs/cors/wrapper/gin"
)

func main(){
	url := "https://static.wongnai.com/devinterview/covid-cases.json"

	repo := repository.NewDataRepository(url)
	svc := service.NewWnService(repo)

	// ds, _ := svc.GetProvince()

	// data, _ := json.Marshal(ds)
	// fmt.Println(string(data))
	ctrl := controller.NewWnController(svc)

	r := gin.Default()
	r.GET("/", ctrl.GetCovidSum)
	
	r.Run()
}

