package osm

import (
	"geospatial"
	"math"
)

// OSMTile OSM type tile
type OSMTile struct {
	X, Y int
}

// GetTile ...
func GetTile(point geospatial.Point, lod int) OSMTile {
	return OSMTile{
		getTileX(point.Latitude, lod),
		getTileY(point.Longitude, lod)}
}

func getTileX(lon float64, lod int) int {
	x := int(math.Floor((lon + 180.0) / 360.0 * (math.Exp2(float64(lod)))))
	return x
}

func getTileY(lat float64, lod int) int {
	y := int(math.Floor((1.0 - math.Log(math.Tan(lat*math.Pi/180.0)+1.0/math.Cos(lat*math.Pi/180.0))/math.Pi) / 2.0 * (math.Exp2(float64(lod)))))
	return y
}
