package main

import (
	"series-tracker-backend/database"
	"series-tracker-backend/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "series-tracker-backend/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Inicializa la conexión a la base de datos.
	database.InitDB()

	// Crea una nueva instancia del router de Gin.
	r := gin.Default()

	// Configura la ruta para la documentación Swagger.
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configura CORS para permitir solicitudes desde el cliente.
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost"}, // Orígenes permitidos.
		AllowMethods: []string{"*"},                // Métodos HTTP permitidos.
		AllowHeaders: []string{"*"},                // Encabezados permitidos.
	}))

	// Define el grupo de rutas para la API.
	api := r.Group("/api")
	{
		// Rutas para manejar las series.
		api.GET("/series", handlers.GetSeries)          // Obtiene todas las series.
		api.GET("/series/:id", handlers.GetSerieByID)   // Obtiene una serie por ID.
		api.POST("/series", handlers.CreateSerie)       // Crea una nueva serie.
		api.PUT("/series/:id", handlers.UpdateSerie)    // Actualiza una serie existente.
		api.DELETE("/series/:id", handlers.DeleteSerie) // Elimina una serie por ID.

		// Rutas para operaciones específicas en las series.
		api.PATCH("/series/:id/status", handlers.UpdateStatus)      // Actualiza el estado de una serie.
		api.PATCH("/series/:id/episode", handlers.IncrementEpisode) // Incrementa el episodio actual.
		api.PATCH("/series/:id/upvote", handlers.Upvote)            // Incrementa los votos positivos.
		api.PATCH("/series/:id/downvote", handlers.Downvote)        // Incrementa los votos negativos.
	}

	// Inicia el servidor en el puerto 8080.
	r.Run(":8080")
}
