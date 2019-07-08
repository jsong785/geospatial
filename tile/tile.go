package tile

import (
	"errors"
	"geospatial"
)

// TileData - interface for data, an AOI of the same spot indicates a point.
type TileData interface {
	GetAOI() geospatial.AOI
}

// Tile - tile data.
type Tile struct {
	Parent   *Tile
	Children []*Tile
	AOI      geospatial.AOI
	Data     []TileData
}

// FindPoint - Returns a tile containing the given point.
func (t *Tile) FindPoint(point geospatial.Point) *Tile {
	aoi := geospatial.CreateAOI(point, point)
	return t.FindArea(aoi)
}

// FindArea - Returns a tile containing the given area.
func (t *Tile) FindArea(aoi geospatial.AOI) *Tile {
	var foundTile *Tile
	if t.AOI.IsAOIInside(aoi) {
		foundTile = t
		for _, child := range t.Children {
			if found := child.FindArea(aoi); found != nil {
				foundTile = found
				break
			}
		}
	}
	return foundTile
}

// TileLogic - interface for Tile operations.
type TileLogic interface {
	ShouldSplit(tile *Tile) bool
}

// AddTileData - adds data to the quadtree
func (t *Tile) AddTileData(logic TileLogic, data TileData) error {
	found := t.FindArea(data.GetAOI())
	if found == nil {
		return errors.New("error adding tile")
	}

	if logic.ShouldSplit(found) {
		found.split()
		return found.AddTileData(logic, data)
	}

	if found.Data == nil {
		found.Data = make([]TileData, 0)
	}
	found.Data = append(found.Data, data)
	return nil
}

func (t *Tile) split() {
	if t.Children == nil {
		t.Children = make([]*Tile, 4)
		t.Children[0] = new(Tile)
		t.Children[1] = new(Tile)
		t.Children[2] = new(Tile)
		t.Children[3] = new(Tile)

		t.Children[0].Parent = t
		t.Children[1].Parent = t
		t.Children[2].Parent = t
		t.Children[3].Parent = t

		t.Children[0].AOI = getQuadrantAOIUpperLeft(t.AOI)
		t.Children[1].AOI = getQuadrantAOIUpperRight(t.AOI)
		t.Children[2].AOI = getQuadrantAOILowerRight(t.AOI)
		t.Children[3].AOI = getQuadrantAOILowerLeft(t.AOI)
	}
}

func getQuadrantAOIUpperLeft(tileArea geospatial.AOI) geospatial.AOI {
	tileCenter := calculateQuadrantCenter(tileArea)

	lower := geospatial.Point{Latitude: tileCenter.Latitude,
		Longitude: tileArea.Lower.Longitude}
	upper := geospatial.Point{Latitude: tileArea.Upper.Latitude,
		Longitude: tileCenter.Longitude}

	return geospatial.AOI{Lower: lower, Upper: upper}
}

func getQuadrantAOIUpperRight(tileArea geospatial.AOI) geospatial.AOI {
	tileCenter := calculateQuadrantCenter(tileArea)

	lower := tileCenter
	upper := tileArea.Upper

	return geospatial.AOI{Lower: lower, Upper: upper}
}

func getQuadrantAOILowerRight(tileArea geospatial.AOI) geospatial.AOI {
	tileCenter := calculateQuadrantCenter(tileArea)

	lower := geospatial.Point{Latitude: tileArea.Lower.Latitude,
		Longitude: tileCenter.Longitude,
	}
	upper := geospatial.Point{Latitude: tileCenter.Longitude,
		Longitude: tileArea.Upper.Longitude,
	}

	return geospatial.AOI{Lower: lower, Upper: upper}
}

func getQuadrantAOILowerLeft(tileArea geospatial.AOI) geospatial.AOI {
	tileCenter := calculateQuadrantCenter(tileArea)

	lower := tileArea.Lower
	upper := tileCenter

	return geospatial.AOI{Lower: lower, Upper: upper}
}

func calculateQuadrantCenter(area geospatial.AOI) geospatial.Point {
	// need eps?
	latSpan := area.Upper.Latitude - area.Lower.Latitude
	lonSpan := area.Upper.Longitude - area.Lower.Longitude
	quadrantCenterPoint := geospatial.Point{
		Latitude:  area.Lower.Latitude + (latSpan * 0.5),
		Longitude: area.Lower.Longitude + (lonSpan * 0.5),
	}
	return quadrantCenterPoint
}
