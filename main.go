// main.go
package main

import (
    "backend/models"
    "backend/controllers"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"

)

func main() {
    r := gin.Default()
    
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:8000"}, // Sesuaikan dengan URL frontend Anda
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))
    models.InitDB()
    r.GET("/barang", controllers.GetAllBarang)
    r.POST("/register", controllers.Register)
    r.POST("/login", controllers.Login)
    r.GET("/barang/:id", controllers.GetBarangByID)
    authorized := r.Group("/")
    {
        authorized.POST("/barang", controllers.CreateBarang)
        authorized.PUT("/barang/:id", controllers.UpdateBarang)
        authorized.DELETE("/barang/:id", controllers.DeleteBarang)
        authorized.GET("/mutasi", controllers.GetAllMutasi)
        authorized.POST("/mutasi", controllers.CreateMutasi)
        authorized.GET("/mutasi/:barang_id", controllers.GetMutasiByBarang)
    }

    r.Run(":8080")
}
