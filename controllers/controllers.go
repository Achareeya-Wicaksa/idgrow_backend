// controllers/controllers.go
package controllers

import (
    "backend/models"
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "net/http"
    "time"
    "golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("secret_key")

type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}
func GetAllMutasi(c *gin.Context) {
    var mutasi []models.Mutasi
    err := models.DB.Preload("User").Preload("Barang").Find(&mutasi).Error
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "details": err.Error()})
        return
    }

    if len(mutasi) == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "No mutations found"})
        return
    }

    c.JSON(http.StatusOK, mutasi)
}


func Register(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    user.Password = HashPassword(user.Password)
    if err := models.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
    var input models.User
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    var user models.User
    if err := models.DB.Where("email = ?", input.Email).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    if !CheckPasswordHash(input.Password, user.Password) {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
        return
    }

    token, err := GenerateToken(user.Email)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"token": token})
}

func GetAllBarang(c *gin.Context) {
    var barang []models.Barang
    models.DB.Find(&barang)
    c.JSON(http.StatusOK, barang)
}

func CreateBarang(c *gin.Context) {
    var barang models.Barang
    if err := c.ShouldBindJSON(&barang); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Create(&barang)
    c.JSON(http.StatusOK, barang)
}

func GetBarangByID(c *gin.Context) {
    var barang models.Barang
    if err := models.DB.Where("id = ?", c.Param("id")).First(&barang).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Barang not found"})
        return
    }
    c.JSON(http.StatusOK, barang)
}

func UpdateBarang(c *gin.Context) {
    var barang models.Barang
    if err := models.DB.Where("id = ?", c.Param("id")).First(&barang).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Barang not found"})
        return
    }

    if err := c.ShouldBindJSON(&barang); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Save(&barang)
    c.JSON(http.StatusOK, barang)
}

func DeleteBarang(c *gin.Context) {
    var barang models.Barang
    if err := models.DB.Where("id = ?", c.Param("id")).First(&barang).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Barang not found"})
        return
    }

    models.DB.Delete(&barang)
    c.JSON(http.StatusOK, gin.H{"message": "Barang deleted successfully"})
}

func CreateMutasi(c *gin.Context) {
    var mutasi models.Mutasi
    if err := c.ShouldBindJSON(&mutasi); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    models.DB.Create(&mutasi)
    c.JSON(http.StatusOK, mutasi)
}

func GetMutasiByBarang(c *gin.Context) {
    var mutasi []models.Mutasi
    if err := models.DB.Where("barang_id = ?", c.Param("barang_id")).Find(&mutasi).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "No mutations found for this item"})
        return
    }

    c.JSON(http.StatusOK, mutasi)
}

func HashPassword(password string) string {
    bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes)
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func GenerateToken(email string) (string, error) {
    expirationTime := time.Now().Add(24 * time.Hour)
    claims := &Claims{
        Email: email,
        StandardClaims: jwt.StandardClaims{
            ExpiresAt: expirationTime.Unix(),
        },
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}
