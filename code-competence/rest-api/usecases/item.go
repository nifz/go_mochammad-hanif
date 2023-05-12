package usecases

import (
	"code-competence/rest-api/models"
	"code-competence/rest-api/repositories"
)

type ItemUsecase interface {
	GetAllItems() ([]models.Item, error)
	GetItemByID(id string) (models.Item, error)
	CreateItem(item *models.Item) error
	UpdateItem(item *models.Item) error
	DeleteItem(id string) error
	GetItemsByCategoryID(categoryID uint) ([]models.Item, error)
	GetItemsByName(name string) ([]models.Item, error)
}

type itemUsecase struct {
	itemRepo repositories.ItemRepository
}

func NewItemUsecase(itemRepo repositories.ItemRepository) ItemUsecase {
	return &itemUsecase{itemRepo}
}

func (u *itemUsecase) GetAllItems() ([]models.Item, error) {
	return u.itemRepo.GetAllItems()
}

func (u *itemUsecase) GetItemByID(id string) (models.Item, error) {
	return u.itemRepo.GetItemByID(id)
}

func (u *itemUsecase) CreateItem(item *models.Item) error {
	return u.itemRepo.CreateItem(item)
}

func (u *itemUsecase) UpdateItem(item *models.Item) error {
	return u.itemRepo.UpdateItem(item)
}

func (u *itemUsecase) DeleteItem(id string) error {
	return u.itemRepo.DeleteItem(id)
}

func (u *itemUsecase) GetItemsByCategoryID(categoryID uint) ([]models.Item, error) {
	return u.itemRepo.GetItemsByCategoryID(categoryID)
}

func (u *itemUsecase) GetItemsByName(name string) ([]models.Item, error) {
	return u.itemRepo.GetItemsByName(name)
}
