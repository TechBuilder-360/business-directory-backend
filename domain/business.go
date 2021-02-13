package domain

type Business struct {
	BusId		 int	`db:"business_id"`
	BusName      string
	Services	 []string
	GeoLat       float32
	GeoLong      float32
	ShowComments bool
	active       bool
	verified     bool
}