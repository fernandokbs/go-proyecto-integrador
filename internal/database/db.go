package database

import (
	"os"
	"fmt"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"github.com/fernandokbs/goimage/internal/models"
	log "github.com/fernandokbs/goimage/internal/logger"
)

var DB *gorm.DB

func GetConnection() (*gorm.DB, error) {
	err := godotenv.Load()

	if err != nil {
		log.LogInfo("⚠️ No se encontró archivo .env, usando variables del sistema", nil)
	}

	dsn := os.Getenv("DATABASE_URL")
	
	if dsn == "" {
		log.LogError("❌ No se encontró la variable DATABASE_URL en el entorno", nil)
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return db, nil
}

func Connect() {
	db, err := GetConnection()

	if err != nil {
		log.LogError("❌ Error al conectar a MySQL:", nil)
	}

	db.AutoMigrate(&models.Image{})
	
	fmt.Println("✅ Conectado a MySQL", nil)
}