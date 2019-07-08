package imagery

import (
    "geospatial"
    "gotest"
    "testing"
    "tile"
)

func TestGetImages(t *testing.T) {
    gotest.Initialize(gotest.TestImplement{Test: t})
    defer gotest.Cleanup()

    root := tile.Tile{AOI: geospatial.AOI{
            Lower: geospatial.Point{Latitude: -90.0, Longitude: -180.0},
            Upper: geospatial.Point{Latitude: 90.0, Longitude: 180.0}}}

    path := GetImages(root, geospatial.AOI{
            Lower: geospatial.Point{Latitude: -90.0, Longitude: -180.0},
            Upper: geospatial.Point{Latitude: 90.0, Longitude: 180.0}})

    //checkPath := OSM_ROOT_DIR + "/"
    gotest.ASSERT_EQ(len(path), 4)
}
