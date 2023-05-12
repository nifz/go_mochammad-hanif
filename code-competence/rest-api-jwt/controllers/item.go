package controllers

import (
	"code-competence/rest-api-jwt/dtos"
	"code-competence/rest-api-jwt/models"
	"code-competence/rest-api-jwt/usecases"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ItemController interface {
	GetAllItems(c echo.Context) error
	GetItemByID(c echo.Context) error
	CreateItem(c echo.Context) error
	UpdateItem(c echo.Context) error
	DeleteItem(c echo.Context) error
	GetItemsByCategoryID(c echo.Context) error
	GetItemsByName(c echo.Context) error
}

type itemController struct {
	itemUsecase usecases.ItemUsecase
}

func NewItemController(itemUsecase usecases.ItemUsecase) ItemController {
	return &itemController{itemUsecase}
}

func (c *itemController) GenerateUUID() string {
	uuidWithHyphen := uuid.New()
	uuid := uuidWithHyphen.String()
	return uuid
}

// Implementasi fungsi-fungsi dari interface ItemController

func (c *itemController) GetAllItems(ctx echo.Context) error {
	items, err := c.itemUsecase.GetAllItems()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	var itemDTOs []dtos.ItemDTO

	for _, item := range items {
		itemDTO := dtos.ItemDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			Category: dtos.CategoriesDTO{
				Name: item.Category.Name,
			},
		}
		itemDTOs = append(itemDTOs, itemDTO)
	}

	return ctx.JSON(http.StatusOK, dtos.ItemDTOsResponse{
		Message: "Successfully get all items.",
		Data:    itemDTOs,
	})
}

func (c *itemController) GetItemByID(ctx echo.Context) error {
	id := ctx.Param("id")
	item, err := c.itemUsecase.GetItemByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dtos.ItemDTOResponse{
		Message: "Successfully get item by id.",
		Data: dtos.ItemDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			Category: dtos.CategoriesDTO{
				Name: item.Category.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		},
	})
}

func (c *itemController) CreateItem(ctx echo.Context) error {
	var itemDTO dtos.ItemDTO
	if err := ctx.Bind(&itemDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	itemDTO.ID = c.GenerateUUID()

	item := models.Item{
		ID:          itemDTO.ID,
		Name:        itemDTO.Name,
		Description: itemDTO.Description,
		Stock:       itemDTO.Stock,
		Price:       itemDTO.Price,
		CategoryID:  itemDTO.CategoryID,
	}

	err := c.itemUsecase.CreateItem(&item)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, dtos.ItemDTOResponse{
		Message: "Successfully created item.",
		Data: dtos.ItemDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			Category: dtos.CategoriesDTO{
				Name: item.Category.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		},
	})
}

func (c *itemController) UpdateItem(ctx echo.Context) error {
	id := ctx.Param("id")

	var itemDTO dtos.ItemDTO
	if err := ctx.Bind(&itemDTO); err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	item, err := c.itemUsecase.GetItemByID(id)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	item.Name = itemDTO.Name
	item.Description = itemDTO.Description
	item.Stock = itemDTO.Stock
	item.Price = itemDTO.Price
	item.CategoryID = itemDTO.CategoryID

	err = c.itemUsecase.UpdateItem(&item)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dtos.ItemDTOResponse{
		Message: "Successfully updated item.",
		Data: dtos.ItemDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			Category: dtos.CategoriesDTO{
				Name: item.Category.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		},
	})
}

func (c *itemController) DeleteItem(ctx echo.Context) error {
	id := ctx.Param("id")

	err := c.itemUsecase.DeleteItem(id)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, dtos.ErrorDTO{
		Message: "Successfully deleted item.",
	})
}

func (c *itemController) GetItemsByCategoryID(ctx echo.Context) error {
	categoryID, _ := strconv.Atoi(ctx.Param("category_id"))

	items, err := c.itemUsecase.GetItemsByCategoryID(uint(categoryID))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	var itemDTOs []dtos.ItemDTO

	for _, item := range items {
		itemDTO := dtos.ItemDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			Category: dtos.CategoriesDTO{
				Name: item.Category.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		itemDTOs = append(itemDTOs, itemDTO)
	}

	return ctx.JSON(http.StatusOK, dtos.ItemDTOsResponse{
		Message: "Successfully get items by category id.",
		Data:    itemDTOs,
	})
}

func (c *itemController) GetItemsByName(ctx echo.Context) error {
	name := ctx.QueryParam("keyword")

	items, err := c.itemUsecase.GetItemsByName(name)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, dtos.ErrorDTO{
			Message: err.Error(),
		})
	}

	var itemDTOs []dtos.ItemDTO

	for _, item := range items {
		itemDTO := dtos.ItemDTO{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			Stock:       item.Stock,
			Price:       item.Price,
			CategoryID:  item.CategoryID,
			Category: dtos.CategoriesDTO{
				Name: item.Category.Name,
			},
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		}
		itemDTOs = append(itemDTOs, itemDTO)
	}

	return ctx.JSON(http.StatusOK, dtos.ItemDTOsResponse{
		Message: "Successfully get items by name query param.",
		Data:    itemDTOs,
	})
}
