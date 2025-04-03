package database

import (
	"series-tracker-backend/models"

	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB es una variable global que representa la conexión a la base de datos.
var DB *gorm.DB

// InitDB inicializa la conexión a la base de datos PostgreSQL utilizando GORM.
func InitDB() {
	// Construye el Data Source Name (DSN) utilizando variables de entorno.
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),     // Dirección del host de la base de datos.
		os.Getenv("DB_USER"),     // Usuario de la base de datos.
		os.Getenv("DB_PASSWORD"), // Contraseña del usuario.
		os.Getenv("DB_NAME"),     // Nombre de la base de datos.
		os.Getenv("DB_PORT"),     // Puerto de conexión.
	)

	var err error
	// Abre la conexión a la base de datos utilizando GORM y el driver de PostgreSQL.
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	// Si ocurre un error al conectar, se detiene la ejecución y se muestra un mensaje.
	if err != nil {
		panic("Error al conectar a la base de datos: " + err.Error())
	}

	// Realiza la migración automática de la estructura del modelo Serie.
	// Esto asegura que la tabla correspondiente esté creada y actualizada.
	DB.AutoMigrate(&models.Serie{})
}
