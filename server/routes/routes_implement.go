package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"server/collections"
	"server/controller/product_controller"
	"server/controller/rating_controller"
	"server/helpers"
	"server/repository/product_repository"
	"server/repository/rating_repository"
	"server/services/product_service"
	"server/services/rating_service"
)

type Engine struct {
	gin  *gin.Engine
	db   *gorm.DB
	conf helpers.Configuration
}

func NewRoutes(gin *gin.Engine, db *gorm.DB, conf helpers.Configuration) Routes {
	return Engine{
		gin:  gin,
		db:   db,
		conf: conf,
	}
}

func (p Engine) ProductRoutes() {
	newCollection := collections.NewCollection(p.db, p.conf)
	newRepo := product_repository.NewProductRepository(newCollection)
	newService := product_service.NewProductService(newRepo)
	newController := product_controller.NewProductController(newService)
	p.gin.POST("/create-product", newController.CreateProduct)
	p.gin.GET("/get-product", newController.GetProduct)
	p.gin.GET("/get-product/:id", newController.GetProductDetail)
}

func (p Engine) RatingRoutes() {
	newCollection := collections.NewCollection(p.db, p.conf)
	newRepo := rating_repository.NewRatingRepository(newCollection)
	newService := rating_service.NewRatingService(newRepo)
	newController := rating_controller.NewRatingController(newService)
	p.gin.POST("/create-rating/:productId", newController.CreateRating)
	p.gin.GET("/get-rating/:productId", newController.GetAllRatingByProduct)
}
