package orders

import (
	"net/http"
	"time"
	"tugas-dua/models"

	"github.com/gin-gonic/gin"
)

func (h handler) Create(c *gin.Context) {
	var (
		order  models.Order
		result gin.H
	)

	customer_name := c.PostForm("customer_name")

	order.CustomerName = customer_name
	order.OrderedAt = time.Now()

	err := h.DB.Create(&order).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	} else {
		result = gin.H{
			"result": order,
		}
	}
	c.JSON(http.StatusOK, result)
}

func (h handler) Update(c *gin.Context) {
	var (
		order    models.Order
		newOrder models.Order
		result   gin.H
	)
	id := c.Param("orderId")

	err := h.DB.Model(&order).Where("id=?", id).First(&order).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	customer_name := c.PostForm("customer_name")
	newOrder.CustomerName = customer_name

	err = h.DB.Model(&order).Updates(&newOrder).Error
	if err != nil {
		result = gin.H{
			"result":  err.Error(),
			"message": "updated failed",
		}
		c.JSON(http.StatusInternalServerError, result)
		return
	}

	var items []customItem
	for _, itemOrders := range order.Items {
		detailStruct := customItem{
			ItemCode:    itemOrders.ItemCode,
			Description: itemOrders.Description,
			Quantity:    itemOrders.Quantity,
		}
		items = append(items, detailStruct)
	}
	result = gin.H{
		"customerName": order.CustomerName,
		"orderedAt":    order.OrderedAt,
		"items":        items,
	}
	c.JSON(http.StatusOK, result)
}

type customItem struct {
	ItemCode    string
	Description string
	Quantity    int
}

func (h handler) GetId(c *gin.Context) {
	id := c.Param("orderId")
	var (
		order  models.Order
		result gin.H
	)
	err := h.DB.Model(&models.Order{}).Preload("Items").Where("id=?", id).First(&order).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	}

	var items []customItem

	for _, itemOrders := range order.Items {
		detailStruct := customItem{
			ItemCode:    itemOrders.ItemCode,
			Description: itemOrders.Description,
			Quantity:    itemOrders.Quantity,
		}
		items = append(items, detailStruct)
	}

	result = gin.H{
		"customerName": order.CustomerName,
		"orderedAt":    order.OrderedAt,
		"items":        items,
	}

	c.JSON(http.StatusOK, result)
}

func (h handler) Get(c *gin.Context) {
	var result gin.H
	var orders []models.Order
	err := h.DB.Model(&models.Order{}).Preload("Items").Find(&orders).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
			"count":  0,
		}
		c.JSON(http.StatusNotFound, result)
		return
	} else {
		result = gin.H{
			"result": orders,
			"count":  len(orders),
		}
	}
	c.JSON(http.StatusOK, result)
}

func (h handler) Delete(c *gin.Context) {
	id := c.Param("orderId")
	var (
		order  models.Order
		result gin.H
	)

	err := h.DB.First(&order, "id=?", id).Error
	if err != nil {
		result = gin.H{
			"result": "data not found",
		}
		c.JSON(http.StatusNotFound, result)
		return
	}
	err = h.DB.Delete(&order).Error
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
