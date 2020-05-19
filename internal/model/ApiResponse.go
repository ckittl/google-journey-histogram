package model

type ApiResponse struct {
	GeocodedWaypoints []GeoCodedWayPoint `json:"geocoded_waypoints"`
	Routes            []Route            `json:"routes"`
	Status            string             `json:"status"`
}
