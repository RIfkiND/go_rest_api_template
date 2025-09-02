package repositories

import (
    "gorm.io/gorm"
    "gorestapi/internals/models"
)

type CategoryRepository interface {
    Create(category *models.Category) error
    GetAll() ([]models.Category, error)
    Update(category *models.Category) error
    Delete(id uint) error
}
type CategoriesRepository struct {
	DB *gorm.DB
}

func NewCategoriesRepository(db *gorm.DB) *CategoriesRepository {
	return &CategoriesRepository{
		DB: db,
	}
}

func (r *CategoriesRepository) Create(category *models.Category) error {
	return r.DB.Create(category).Error
}

func (r *CategoriesRepository) GetAll() ([]models.Category, error) {
	var categories []models.Category
	err := r.DB.Find(&categories).Error
	return categories, err
}

func (r *CategoriesRepository) Update(category *models.Category) error {
	return r.DB.Save(category).Error
}

func (r *CategoriesRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Category{}, id).Error
}
