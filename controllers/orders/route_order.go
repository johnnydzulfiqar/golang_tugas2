package orders

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := &handler{
		DB: db,
	}
	router := r.Group("/orders")
	router.POST("/", h.Create)
	router.PUT("/:orderId", h.Update)
	router.GET("/:orderId", h.GetId)
	router.GET("/", h.Get)
	router.DELETE("/:orderId", h.Delete)
}
