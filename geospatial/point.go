package geospatial

import (
	"math"
)

// Point - latittude/longitude degrees
type Point struct {
	Latitude  float64
	Longitude float64
}

// ConvertToRadians - convert from degrees to radians.
func ConvertToRadians(p *Point) Point {
	var radians Point
	radians.Latitude = p.Latitude * math.Pi / 180.0
	radians.Longitude = p.Longitude * math.Pi / 180.0
	return radians
}

// ConvertToDegrees - convert from radians to degrees.
func ConvertToDegrees(p *Point) Point {
	var degrees Point
	degrees.Latitude = p.Latitude * 180.0 / math.Pi
	degrees.Longitude = p.Longitude * 180.0 / math.Pi
	return degrees
}
