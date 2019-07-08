package tile

import (
	"geospatial"
	"gotest"
	"testing"
)

type MockTileLogic struct {
	TileLogic
}

func (l MockTileLogic) ShouldSplit(tile *Tile) bool {
	return len(tile.Data) >= 2
}

type MockTileData struct {
	TileData
	number int
	aoi    geospatial.AOI
}

func (t MockTileData) GetAOI() geospatial.AOI {
	return t.aoi
}

func GetMockTileData(num int, lower, upper geospatial.Point) *MockTileData {
	data := new(MockTileData)
	data.number = num
	data.aoi.Lower = lower
	data.aoi.Upper = upper
	data.aoi.Normalise() // just in case
	return data
}
func TestTileSplit(t *testing.T) {
	gotest.Initialize(gotest.TestImplement{Test: t})
	defer gotest.Cleanup()

	root := Tile{AOI: geospatial.AOI{
		Lower: geospatial.Point{Latitude: 0.0, Longitude: 0.0},
		Upper: geospatial.Point{Latitude: 100.0, Longitude: 100.0}}}

	data := (*MockTileData)(nil)
	notSplitYetTest := func(t *testing.T) {
		err := root.AddTileData(MockTileLogic{}, data)
		gotest.ASSERT_EQ(err, nil)
		gotest.ASSERT_EQ(root.Children, ([]*Tile)(nil))

		found := root.FindArea(data.GetAOI())
		gotest.ASSERT_NE(found, nil)
		gotest.ASSERT_EQ(found, &root)
		gotest.ASSERT_EQ(found.Parent, (*Tile)(nil))
	}

	hasSplitTest := func(t *testing.T) {
		root.AddTileData(MockTileLogic{}, data)
		gotest.ASSERT_EQ(len(root.Children), 4)

		found := root.FindArea(data.GetAOI())
		gotest.ASSERT_NE(found, nil)
		gotest.ASSERT_NE(found, &root)
		gotest.ASSERT_EQ(found.Parent, &root)
	}

	data = GetMockTileData(0,
		geospatial.Point{Latitude: 1.0, Longitude: 1.0},
		geospatial.Point{Latitude: 2.0, Longitude: 2.0})
	t.Run("First insert", notSplitYetTest)

	data = GetMockTileData(10,
		geospatial.Point{Latitude: 41.0, Longitude: 1.0},
		geospatial.Point{Latitude: 42.0, Longitude: 2.0})
	t.Run("First insert", notSplitYetTest)

	data = GetMockTileData(10,
		geospatial.Point{Latitude: 41.0, Longitude: 1.0},
		geospatial.Point{Latitude: 42.0, Longitude: 2.0})
	t.Run("Third insert", hasSplitTest)

	data = GetMockTileData(20,
		geospatial.Point{Latitude: 42.0, Longitude: 42.0},
		geospatial.Point{Latitude: 43.0, Longitude: 43.0})
	t.Run("Third insert", hasSplitTest)

	found := root.FindArea(data.aoi)
	gotest.ASSERT_NE(found, nil)
	gotest.ASSERT_EQ(found.Parent, &root)

	// look for area that is only in the root
	found = root.FindArea(geospatial.AOI{
		Upper: geospatial.Point{Latitude: 49, Longitude: 49},
		Lower: geospatial.Point{Latitude: 51, Longitude: 51},
	})
	gotest.ASSERT_NE(found, nil)
	gotest.ASSERT_EQ(found.Parent, (*Tile)(nil))
}
