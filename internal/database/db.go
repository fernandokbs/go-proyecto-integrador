package database

import (
    "gorm.io/gorm"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()

	if err != nil {
		log.Println("⚠️ No se encontró archivo .env, usando variables del sistema")
	}

	dsn := os.Getenv("DATABASE_URL")
	
	if dsn == "" {
		log.Fatal("❌ No se encontró la variable DATABASE_URL en el entorno")
	}

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("❌ Error al conectar a MySQL:", err)
	}

	// Aqui va la automigracion
	
	fmt.Println("✅ Conectado a MySQL")
}