package osm

import (
    "geospatial"
    "tile_imagery"
    "testing"
)

func TestTileGeneric(t *testing.T) {
    tile := GetTile(50.0, 100.0, 3)

    aoi := geospatial.CreateAOIFromLatLon(-90, -180, 90, 180)
    tree := tile_imagery.CreateTileTree(aoi,
                            10,
                            tile_imagery.CLOCKWISE,
                            tile_imagery.UPPER_LEFT)
}
