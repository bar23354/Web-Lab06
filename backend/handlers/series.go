package handlers

// @title Series Tracker API
// Comentarios generados con ChatGPT-4.0
import (
	"series-tracker-backend/database"
	"series-tracker-backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// @Summary Obtener todas las series
// @Description Retorna una lista de series con filtros opcionales
// @Produce json
// @Param search query string false "Buscar por título"
// @Param status query string false "Filtrar por estado"
// @Param sort query string false "Ordenar por ranking (asc/desc)"
// @Success 200 {array} models.Serie
// @Router /api/series [get]

func UpdateSerie(c *gin.Context) {
	id := c.Param("id")
	var serie models.Serie

	// Buscar la serie existente
	if err := database.DB.First(&serie, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	// Vincular datos JSON
	if err := c.ShouldBindJSON(&serie); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// Guardar cambios
	if err := database.DB.Save(&serie).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar la serie"})
		return
	}

	c.JSON(200, serie)
}

func GetSeries(c *gin.Context) {
	var series []models.Serie
	db := database.DB

	// Filtros
	status := c.Query("status")
	search := c.Query("search")

	if status != "" {
		db = db.Where("status = ?", status)
	}

	if search != "" {
		db = db.Where("title ILIKE ?", "%"+search+"%")
	}

	// Ordenamiento
	sortOrder := c.Query("sort")
	if sortOrder == "asc" {
		db = db.Order("ranking asc")
	} else {
		db = db.Order("ranking desc")
	}

	if err := db.Find(&series).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al obtener series"})
		return
	}

	c.JSON(200, series)
}

func CreateSerie(c *gin.Context) {
	var newSerie models.Serie

	// Validar el JSON recibido
	if err := c.ShouldBindJSON(&newSerie); err != nil {
		c.JSON(400, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Validar campos requeridos
	if newSerie.Title == "" {
		c.JSON(400, gin.H{"error": "El campo 'title' es obligatorio"})
		return
	}

	// Crear la serie en la base de datos
	result := database.DB.Create(&newSerie)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al crear la serie: " + result.Error.Error()})
		return
	}

	c.JSON(201, newSerie)
}

func IncrementEpisode(c *gin.Context) {
	id := c.Param("id")
	var serie models.Serie

	// Buscar la serie
	if err := database.DB.First(&serie, id).Error; err != nil {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	// Validar episodio máximo
	if serie.LastEpisodeWatched >= serie.TotalEpisodes {
		c.JSON(400, gin.H{"error": "Ya has visto todos los episodios"})
		return
	}

	// Incrementar episodio
	serie.LastEpisodeWatched++

	// Guardar cambios
	if err := database.DB.Save(&serie).Error; err != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar el episodio: " + err.Error()})
		return
	}

	c.JSON(200, serie)
}

// Handler: Obtener serie por ID
func GetSerieByID(c *gin.Context) {
	id := c.Param("id")
	var serie models.Serie
	result := database.DB.First(&serie, id)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(200, serie)
}

func DeleteSerie(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Delete(&models.Serie{}, id)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al eliminar la serie"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(200, gin.H{"message": "Serie eliminada correctamente"})
}

func UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var request struct {
		Status string `json:"status"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	result := database.DB.Model(&models.Serie{}).
		Where("id = ?", id).
		Update("status", request.Status)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar el estado"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(200, gin.H{"message": "Estado actualizado correctamente"})
}

func Upvote(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Model(&models.Serie{}).
		Where("id = ?", id).
		Update("ranking", gorm.Expr("ranking + 1"))

	handleVoteResult(c, result)
}

func Downvote(c *gin.Context) {
	id := c.Param("id")
	result := database.DB.Model(&models.Serie{}).
		Where("id = ?", id).
		Update("ranking", gorm.Expr("ranking - 1"))

	handleVoteResult(c, result)
}

// Función auxiliar para manejar resultados de votación
func handleVoteResult(c *gin.Context, result *gorm.DB) {
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error al actualizar el ranking"})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Serie no encontrada"})
		return
	}

	c.JSON(200, gin.H{"message": "Ranking actualizado correctamente"})
}
