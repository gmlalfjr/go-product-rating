package rating_service

import (
	"server/entity"
	"server/exception"
	"server/repository/rating_repository"
	"server/web/rating_web"
	"time"
)

type RatingServiceImpl struct {
	ratingRepository rating_repository.RatingRepository
}

func NewRatingService(r rating_repository.RatingRepository) RatingService {
	return RatingServiceImpl{ratingRepository: r}
}

func (r RatingServiceImpl) CreateRating(request *rating_web.RatingCreateRequest) (*rating_web.RatingCreateResponse, *exception.ErrorResponse) {
	rating := entity.Rating{
		Star:      request.Star,
		Comment:   request.Comment,
		ProductId: request.ProductId,
		CreatedAt: time.Now(),
	}
	data, err := r.ratingRepository.CreateRating(&rating)
	if err != nil {
		return nil, err
	}

	return &rating_web.RatingCreateResponse{
		Star:      data.Star,
		Comment:   data.Comment,
		ProductId: data.ProductId,
	}, nil
}

func (r RatingServiceImpl) GetAllProductRating(id int) ([]rating_web.RatingCreateResponse, *exception.ErrorResponse) {
	data, err := r.ratingRepository.GetAllRatingProduct(id)
	if err != nil {
		return nil, err
	}

	//var ratings []rating_web.RatingCreateResponse
	ratings := make([]rating_web.RatingCreateResponse, 0)
	for _, val := range data {
		ratings = append(ratings, rating_web.RatingCreateResponse{
			Star:      val.Star,
			Comment:   val.Comment,
			ProductId: val.ProductId,
		})
	}

	return ratings, nil
}
