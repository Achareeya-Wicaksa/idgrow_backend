package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/joho/godotenv"
    "time"
    "os"
    "log"
    "fmt"
)

var DB *gorm.DB

type User struct {
    gorm.Model
    Name     string `json:"name"`
    Email    string `json:"email" gorm:"unique"`
    Password string `json:"-"`
}

type Barang struct {
    gorm.Model
    NamaBarang string `json:"nama_barang"`
    Kategori   string `json:"kategori"`
    Lokasi     string `json:"lokasi"`
}

type Mutasi struct {
    gorm.Model
    Tanggal    time.Time `json:"tanggal"`
    JenisMutasi string   `json:"jenis_mutasi"`
    Jumlah     int       `json:"jumlah"`
    UserID     uint      `json:"user_id"`
    BarangID   uint      `json:"barang_id"`
    User       User      `json:"user" gorm:"foreignkey:UserID"`
    Barang     Barang    `json:"barang" gorm:"foreignkey:BarangID"`
}


func InitDB() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
    
    dsn := os.Getenv("DSN")
    
    DB, err = gorm.Open("mysql", dsn)
    
    if err != nil {
        panic("Failed to connect to database!")
    }
    
    fmt.Println("Database connected!")
    DB.AutoMigrate(&User{}, &Barang{}, &Mutasi{})
}
