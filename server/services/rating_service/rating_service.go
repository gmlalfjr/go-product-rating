package rating_service

import (
	"server/exception"
	"server/web/rating_web"
)

type RatingService interface {
	CreateRating(request *rating_web.RatingCreateRequest) (*rating_web.RatingCreateResponse, *exception.ErrorResponse)
	GetAllProductRating(id int) ([]rating_web.RatingCreateResponse, *exception.ErrorResponse)
}
