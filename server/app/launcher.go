package app

import (
	"github.com/gin-gonic/gin"
	"server/db"
	"server/entity"
	"server/helpers"
	"server/routes"
)

func AppLancher() *gin.Engine {
	r := gin.Default()
	conf := helpers.NewConfiguration()
	conn := db.Connection(conf)
	err := conn.AutoMigrate(&entity.Product{})
	err = conn.AutoMigrate(&entity.Rating{})
	if err != nil {
		panic(err)
	}

	routes.NewRoutes(r, conn, conf).ProductRoutes()
	routes.NewRoutes(r, conn, conf).RatingRoutes()

	return r
}
