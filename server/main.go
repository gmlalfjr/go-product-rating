package main

import (
	"github.com/gin-gonic/gin"
	"server/controller"
	"server/db"
	"server/entity"
	"server/helpers"
	"server/repository"
	"server/services"
)

func main(){
	r := gin.Default()

	conf := helpers.NewConfiguration()
	conn := db.Connection(conf)
	err := conn.AutoMigrate(&entity.Product{})
	newRepo := repository.NewProductRepository()
	newService := services.NewProductService(newRepo, conn)
	newController := controller.NewProductController(newService)

	r.GET("/ping", newController.CreateProduct)
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
