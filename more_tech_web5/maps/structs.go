package maps

type Coordinates struct {
	Latitude  float64
	Longitude float64
}

type RouteResponse struct {
	Routes []struct {
		Geometry GeoJSON `json:"geometry"`
		Duration float64 `json:"duration"`
	} `json:"routes"`
}

type GeoJSON struct {
	Type        string      `json:"type"`
	Coordinates [][]float64 `json:"coordinates"`
}

type Result struct {
	Path     GeoJSON `json:"path"`
	Duration float64 `json:"duration"`
}

type MapRoute struct {
	Start         Coordinates `json:"start"`
	End           Coordinates `json:"end"`
	TransportType string      `json:"transportType"`
}
