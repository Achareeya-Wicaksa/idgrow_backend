// main.go
package main

import (
    "backend/models"
    "backend/controllers"
    "backend/middlewares"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()

    models.InitDB()

    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)

    authorized := r.Group("/")
    authorized.Use(middlewares.AuthMiddleware())
    {
        authorized.GET("/barang", controllers.GetAllBarang)
        authorized.POST("/barang", controllers.CreateBarang)
        authorized.GET("/barang/:id", controllers.GetBarangByID)
        authorized.PUT("/barang/:id", controllers.UpdateBarang)
        authorized.DELETE("/barang/:id", controllers.DeleteBarang)
        authorized.POST("/mutasi", controllers.CreateMutasi)
        authorized.GET("/mutasi/:barang_id", controllers.GetMutasiByBarang)
    }

    r.Run(":8080")
}
