package items

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

	router := r.Group("/items")
	router.POST("/", h.Create)
	router.PUT("/:itemId", h.Update)
	router.GET("/:itemId", h.GetId)
	router.GET("/", h.Get)
	router.DELETE("/:itemId", h.Delete)
}
