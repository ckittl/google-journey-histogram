package model

type Bounds struct {
	NorthEast LatLon `json:"north_east"`
	SouthWest LatLon `json:"south_west"`
}
