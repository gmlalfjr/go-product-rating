package collections

import "gorm.io/gorm"

type Collection interface {
	ProductDB() *gorm.DB
}
