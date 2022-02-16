package main

import (
	"github.com/gin-gonic/gin"
	"server/collections"
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
	newCollection := collections.NewCollection(conn, conf)
	newRepo := repository.NewProductRepository(newCollection)
	newService := services.NewProductService(newRepo)
	newController := controller.NewProductController(newService)

	r.GET("/ping", newController.CreateProduct)
	err = r.Run()
	if err != nil {
		panic(err)
	}
}
