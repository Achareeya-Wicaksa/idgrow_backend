package models

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/joho/godotenv"
    "encoding/json"
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
    Tanggal     time.Time `json:"tanggal"`
    JenisMutasi string    `json:"jenis_mutasi"`
    Jumlah      int       `json:"jumlah"`
    UserID      uint      `json:"user_id"`
    BarangID    uint      `json:"barang_id"`
    User        User      `json:"user" gorm:"foreignKey:UserID"`
    Barang      Barang    `json:"barang" gorm:"foreignKey:BarangID"`
}



const timeLayout = "2006-01-02T15:04:05Z07:00"

// ParseJSON menyesuaikan input waktu dengan format yang diinginkan
func (m *Mutasi) UnmarshalJSON(data []byte) error {
    var raw map[string]interface{}
    if err := json.Unmarshal(data, &raw); err != nil {
        return err
    }
    
    if t, ok := raw["tanggal"].(string); ok {
        parsedTime, err := time.Parse(timeLayout, t)
        if err != nil {
            return err
        }
        m.Tanggal = parsedTime
    }
    
    // parsing field lainnya
    
    return nil
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
