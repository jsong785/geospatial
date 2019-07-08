package geospatial

import (
	"math"
)

// AOI - contains a lower and upper point.
type AOI struct {
	Lower Point
	Upper Point
}

// CreateAOI - CreateAOIs an AOI from two points; points do not have to be in a particular order.
func CreateAOI(p1, p2 Point) AOI {
	aoi := AOI{p1, p2}
	aoi.Normalise()
	return aoi
}

// Normalise - Sorts a given points into a lower/upper.
func (a *AOI) Normalise() {
	lower := a.Lower
	upper := a.Upper

	a.Lower = getLowerBoundsPoint(lower, upper)
	a.Upper = getUpperBoundsPoint(lower, upper)
}

// AddPoint - Adds a point to the AOI, expands it if necessary.
func (a *AOI) AddPoint(p Point) {
	lower := a.Lower
	upper := a.Upper

	a.Lower = getLowerBoundsPoint(lower, p)
	a.Upper = getUpperBoundsPoint(upper, p)
}

func getLowerBoundsPoint(p1, p2 Point) Point {
	return Point{
		math.Min(p1.Latitude, p2.Latitude),
		math.Min(p1.Longitude, p2.Longitude),
	}
}

func getUpperBoundsPoint(p1, p2 Point) Point {
	return Point{
		math.Max(p1.Latitude, p2.Latitude),
		math.Max(p1.Longitude, p2.Longitude),
	}
}

// IsValid - Returns true if the AOI is valid.
func (a AOI) IsValid() bool {
	Lower := &a.Lower
	Upper := &a.Upper
	return Lower.Latitude <= Upper.Latitude &&
		Lower.Longitude <= Upper.Longitude
}

// IsAOIInside - Returns true if the given aoi is inside this.
func (a AOI) IsAOIInside(other AOI) bool {
	return a.IsPointInside(other.Lower) &&
		a.IsPointInside(other.Upper)
}

// IsAOIInsideIncludeEdges - Returns true if the given aoi is inside this, including edges.
func (a AOI) IsAOIInsideIncludeEdges(other AOI) bool {
	return a.IsPointInsideIncludeEdges(other.Lower) &&
		a.IsPointInsideIncludeEdges(other.Upper)
}

// IsPointInside - Returns true if the given point is inside this.
func (a AOI) IsPointInside(p Point) bool {
	Lower := &a.Lower
	Upper := &a.Upper

	return Lower.Latitude < p.Latitude &&
		Lower.Longitude < p.Longitude &&
		p.Latitude < Upper.Latitude &&
		p.Longitude < Upper.Longitude
}

// IsPointInsideIncludeEdges - Returns true if the given point is inside this, including edges.
func (a AOI) IsPointInsideIncludeEdges(p Point) bool {
	Lower := &a.Lower
	Upper := &a.Upper

	return Lower.Latitude <= p.Latitude &&
		Lower.Longitude <= p.Longitude &&
		p.Latitude <= Upper.Latitude &&
		p.Longitude <= Upper.Longitude
}
