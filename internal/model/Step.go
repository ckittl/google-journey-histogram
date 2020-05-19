package model

type Step struct {
	Distance         TextValue `json:"distance"`
	Duration         TextValue `json:"duration"`
	EndLocation      LatLon    `json:"end_location"`
	StartLocation    LatLon    `json:"start_location"`
	HtmlInstructions string    `json:"html_instructions"`
	PolyLine         PolyLine  `json:"polyline"`
	TravelMode       string    `json:"travel_mode"`
}
