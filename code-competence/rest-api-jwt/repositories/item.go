package repositories

import (
	"code-competence/rest-api-jwt/models"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(id string) (models.Item, error)
	CreateItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(id string) error
	GetItemsByCategoryID(categoryID uint) ([]models.Item, error)
	GetItemsByName(name string) ([]models.Item, error)
}

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	return &itemRepository{db}
}

// Implementasi fungsi-fungsi dari interface ItemRepository

func (r *itemRepository) GetAllItems() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Joins("JOIN categories ON items.category_id = categories.id").Where("categories.deleted_at IS NULL").Preload("Category").Find(&items).Error
	return items, err
}

func (r *itemRepository) GetItemByID(id string) (models.Item, error) {
	var item models.Item
	err := r.db.Preload("Category").Where("id = ?", id).First(&item).Error
	return item, err
}

func (r *itemRepository) CreateItem(item *models.Item) error {
	var _ = r.db.Preload("Category").Create(item).Error
	return r.db.Preload("Category").Where("id = ?", item.ID).First(&item).Error
}

func (r *itemRepository) UpdateItem(item *models.Item) error {
	return r.db.Preload("Category").Save(item).Error
}

func (r *itemRepository) DeleteItem(id string) error {
	return r.db.Delete(&models.Item{ID: id}).Error
}

func (r *itemRepository) GetItemsByCategoryID(categoryID uint) ([]models.Item, error) {
	var items []models.Item
	err := r.db.Where("category_id = ?", categoryID).Preload("Category").Find(&items).Error
	return items, err
}

func (r *itemRepository) GetItemsByName(name string) ([]models.Item, error) {
	var items []models.Item
	err := r.db.Where("name LIKE ?", "%"+name+"%").Preload("Category").Find(&items).Error
	return items, err
}
