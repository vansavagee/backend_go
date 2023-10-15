package maps

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func FetchRoute(start Coordinates, end Coordinates, transportType string) (*Result, error) {
	url := fmt.Sprintf("https://router.project-osrm.org/route/v1/%s/%f,%f;%f,%f?overview=full&geometries=geojson",
		transportType,
		start.Longitude, start.Latitude,
		end.Longitude, end.Latitude)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var routeResponse RouteResponse
	err = json.Unmarshal(body, &routeResponse)
	if err != nil {
		return nil, err
	}

	if len(routeResponse.Routes) > 0 {
		coords := routeResponse.Routes[0].Geometry.Coordinates

		// Инвертирование координат для Яндекс Карт
		for i, coord := range coords {
			coords[i] = []float64{coord[1], coord[0]}
		}

		result := &Result{
			Path: GeoJSON{
				Type:        "LineString",
				Coordinates: coords,
			},
			Duration: routeResponse.Routes[0].Duration,
		}

		return result, nil
	}

	return nil, fmt.Errorf("No route found or unexpected format")
}
