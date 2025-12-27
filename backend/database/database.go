package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Connect establece la conexión con la base de datos MySQL
func Connect() error {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Construir DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	// Configurar logger de GORM
	logLevel := logger.Silent
	if os.Getenv("ENVIRONMENT") == "development" {
		logLevel = logger.Info
	}

	// Conectar a la base de datos
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
	})

	if err != nil {
		return fmt.Errorf("error al conectar a la base de datos: %w", err)
	}

	// Configurar pool de conexiones
	sqlDB, err := DB.DB()
	if err != nil {
		return fmt.Errorf("error al obtener DB instance: %w", err)
	}

	// Configuración del pool
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	log.Println("✅ Conexión a la base de datos establecida exitosamente")
	return nil
}

// Close cierra la conexión a la base de datos
func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
