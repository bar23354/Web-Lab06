package models

// Serie representa el modelo de una serie de televisión o contenido similar.
type Serie struct {
	ID                 uint   `gorm:"primary_key" json:"id"`
	Title              string `json:"title"`
	Status             string `json:"status"`
	LastEpisodeWatched int    `json:"lastEpisodeWatched"`
	TotalEpisodes      int    `json:"totalEpisodes"`
	Ranking            int    `json:"ranking"`
}
