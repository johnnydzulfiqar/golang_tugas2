package main

import (
	"tugas-dua/controllers/items"
	"tugas-dua/controllers/orders"
	"tugas-dua/database"

	"github.com/gin-gonic/gin"
)

func main() {
	const PORT = ":8080"
	router := gin.Default()
	db := database.GetConnection()
	items.RegisterRoutes(router, db)
	orders.RegisterRoutes(router, db)
	router.Run(PORT)
}
