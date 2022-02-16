package collections

import (
	"gorm.io/gorm"
	"server/helpers"
)

type CollectionImpl struct {
	*gorm.DB
	config helpers.Configuration
}

func NewCollection(db *gorm.DB, conf helpers.Configuration) Collection {
	return &CollectionImpl{
		DB:   db,
		config: conf,
	}
}

func (c CollectionImpl) ProductDB() *gorm.DB {
	 data := c.Table(c.config.Get("PRODUCT_TABLE"))
	 return data
}
