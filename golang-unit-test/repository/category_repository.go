package repository

import "github.com/abdultalif/golang-unit-test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}