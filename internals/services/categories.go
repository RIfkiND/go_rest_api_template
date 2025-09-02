package services

import (
	"gorestapi/internals/models"
	"gorestapi/internals/repositories""
)


type CategoriesService struct {
	repo repositories.CategoryRepository
}

func NewCategoryServices(repo *repositories.CategoryRepository) *CategoriesService {
	return &CategoriesService{
		repo: repo,
	}
}