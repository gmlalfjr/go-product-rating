package rating_repository

import (
	"server/entity"
	"server/exception"
)

type RatingRepository interface {
	CreateRating(rating *entity.Rating) (*entity.Rating, *exception.ErrorResponse)
	GetAllRatingProduct(id int) ([]entity.Rating, *exception.ErrorResponse)
}
