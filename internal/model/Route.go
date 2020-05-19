package model

type Route struct {
	Bounds           Bounds   `json:"bounds"`
	CopyRights       string   `json:"copyrights"`
	Legs             []Leg    `json:"legs"`
	OverViewPolyLine PolyLine `json:"overview_poly_line"`
	Summary          string   `json:"summary"`
	Warnings         []string `json:"warnings"`
	WayPointOrder    []string `json:"way_point_order"`
}
