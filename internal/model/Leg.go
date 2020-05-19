package model

type Leg struct {
	Distance          TextValue `json:"distance"`
	Duration          TextValue `json:"duration"`
	EndAddress        string    `json:"end_address"`
	EndLocation       LatLon    `json:"end_location"`
	StartAddress      string    `json:"start_address"`
	StartLocation     LatLon    `json:"start_location"`
	Steps             []Step    `json:"steps"`
	TrafficSpeedEntry []string  `json:"traffic_speed_entry"`
	ViaWayPoint       []string  `json:"via_way_point"`
}
