package rating_repository

import (
	"errors"
	"gorm.io/gorm"
	"server/collections"
	"server/entity"
	"server/exception"
)

type RatingRepositoryImpl struct {
	collections.Collection
}

func NewRatingRepository(c collections.Collection) RatingRepository {
	return RatingRepositoryImpl{RatingRepositoryImpl{c}}
}

func (r RatingRepositoryImpl) CreateRating(rating *entity.Rating) (*entity.Rating, *exception.ErrorResponse) {
	if err := r.RatingDB().Create(&rating).Error; err != nil {
		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}
	return rating, nil
}

func (r RatingRepositoryImpl) GetAllRatingProduct(id int) ([]entity.Rating, *exception.ErrorResponse) {
	var rat []entity.Rating
	if err := r.RatingDB().Where("product_id = ?", id).Find(&rat).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &exception.ErrorResponse{
				Code:   500,
				Status: "Data Not Found",
				Data:   err.Error(),
			}
		}
		return nil, &exception.ErrorResponse{
			Code:   500,
			Status: "Internal Server Error",
			Data:   err.Error(),
		}
	}
	return rat, nil
}
