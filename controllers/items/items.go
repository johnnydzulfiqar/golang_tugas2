package items

import (
	"net/http"
	"strconv"
	"tugas-dua/models"

	"github.com/gin-gonic/gin"
)

func (h handler) Create(c *gin.Context) {
	var (
		item   models.Item
		result gin.H
	)

	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	quantity := c.PostForm("quantity")
	order_refer := c.PostForm("order_refer")

	quantityInt, err := strconv.Atoi(quantity)
	order_refer_int, err := strconv.Atoi(order_refer)

	item.ItemCode = item_code
	item.Description = description
	item.Quantity = quantityInt
	item.OrderRefer = order_refer_int

	err = h.DB.Create(&item).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		result = gin.H{
			"result": item,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (h handler) Update(c *gin.Context) {
	var (
		item    models.Item
		newItem models.Item
		result  gin.H
	)
	id := c.Param("itemId")

	err := h.DB.Model(&item).Where("id=?", id).First(&item).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	item_code := c.PostForm("item_code")
	description := c.PostForm("description")
	quantity := c.PostForm("quantity")
	order_refer := c.PostForm("order_refer")

	quantityInt, err := strconv.Atoi(quantity)
	order_refer_int, err := strconv.Atoi(order_refer)

	newItem.ItemCode = item_code
	newItem.Description = description
	newItem.Quantity = quantityInt
	newItem.OrderRefer = order_refer_int

	err = h.DB.Model(&item).Updates(&newItem).Error
	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "updated failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}
	result = gin.H{
		"result": "sucessfully updated data",
	}
	c.JSON(http.StatusOK, result)
}

func (h handler) GetId(c *gin.Context) {
	var (
		item   models.Item
		result gin.H
	)
	id := c.Param("itemId")

	err := h.DB.Model(&item).Where("id=?", id).First(&item).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	} else {
		result = gin.H{
			"result": item,
			"count":  1,
		}
	}
	c.JSON(http.StatusNotFound, result)
}

func (h handler) Get(c *gin.Context) {
	var (
		item   []models.Item
		result gin.H
	)

	err := h.DB.Model(&item).Find(&item).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	} else {
		result = gin.H{
			"result": item,
			"count":  len(item),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (h handler) Delete(c *gin.Context) {
	id := c.Param("itemId")
	var (
		item   models.Item
		result gin.H
	)

	err := h.DB.First(&item, "id=?", id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	err = h.DB.Delete(&item).Error
	if err != nil {
		result = gin.H{
			"result": "delete failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		result = gin.H{
			"result": "sucessfully deleted data",
		}
	}
	c.JSON(http.StatusOK, result)
}
