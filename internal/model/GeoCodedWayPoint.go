package model

type GeoCodedWayPoint struct {
	GeocoderStatus string   `json:"geocoder_status"`
	PlaceId        string   `json:"place_id"`
	Types          []string `json:"types"`
}
