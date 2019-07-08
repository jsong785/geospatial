package geospatial

import (
	"gotest"
	"testing"
)

func TestPointConversions(t *testing.T) {
	gotest.Initialize(gotest.TestImplement{Test: t})
	defer gotest.Cleanup()

	degrees := Point{10.0, 20.0}

	radians := ConvertToRadians(&degrees)
	fromRadians := ConvertToDegrees(&radians)
	gotest.ASSERT_EQ(degrees, fromRadians)
}

func TestAOI_IsValid(t *testing.T) {
	gotest.Initialize(gotest.TestImplement{Test: t})
	defer gotest.Cleanup()

	gotest.ASSERT_TRUE(CreateAOI(Point{0.0, 0.0}, Point{0.0, 0.0}).IsValid())
	gotest.ASSERT_TRUE(CreateAOI(Point{0.0, 0.0}, Point{-1.0, -1.0}).IsValid())
	gotest.ASSERT_TRUE(CreateAOI(Point{-1.0, -2.0}, Point{1.0, 10.0}).IsValid())
	gotest.ASSERT_TRUE(CreateAOI(Point{-1.0, -2.0}, Point{-20.0, -10.0}).IsValid())
}

func TestAOI_AddPoint(t *testing.T) {
	gotest.Initialize(gotest.TestImplement{Test: t})
	defer gotest.Cleanup()

	aoi := CreateAOI(Point{-20.0, -40.0}, Point{-19.0, -39.0})
	gotest.ASSERT_TRUE(aoi.IsValid())

	gotest.ASSERT_TRUE(aoi.IsPointInside(Point{-19.5, -39.5}))

	aoi.AddPoint(Point{
		Latitude:  1.0,
		Longitude: 2.0,
	})

	gotest.ASSERT_TRUE(aoi.IsPointInside(Point{
		Latitude:  0.5,
		Longitude: 1.5,
	}))
}

func TestAOI_IsPointInside(t *testing.T) {
	gotest.Initialize(gotest.TestImplement{Test: t})
	defer gotest.Cleanup()

	aoi := CreateAOI(Point{-50.0, -50.0}, Point{50.0, 50.0})
	gotest.ASSERT_FALSE(aoi.IsPointInside(aoi.Lower))
	gotest.ASSERT_TRUE(aoi.IsPointInside(Point{aoi.Lower.Latitude + 1.0,
		aoi.Lower.Longitude + 1.0}))
}

func TestAOI_IsPointInsideIncludingEdges(t *testing.T) {
	gotest.Initialize(gotest.TestImplement{Test: t})
	defer gotest.Cleanup()

	aoi := CreateAOI(Point{-50.0, -50.0}, Point{50.0, 50.0})
	gotest.ASSERT_TRUE(aoi.IsValid())

	gotest.ASSERT_FALSE(aoi.IsPointInside(aoi.Lower))
	gotest.ASSERT_TRUE(aoi.IsPointInsideIncludeEdges(aoi.Lower))
	gotest.ASSERT_FALSE(aoi.IsPointInside(aoi.Upper))
	gotest.ASSERT_TRUE(aoi.IsPointInsideIncludeEdges(aoi.Upper))
}

func TestAOI_IsAOIInside(t *testing.T) {
	gotest.Initialize(gotest.TestImplement{Test: t})
	defer gotest.Cleanup()

	aoi := CreateAOI(Point{-50.0, -50.0}, Point{50.0, 50.0})
	otherAOI := CreateAOI(Point{aoi.Lower.Latitude + 1.0, aoi.Lower.Longitude + 1.0},
		Point{aoi.Upper.Latitude - 1.0, aoi.Upper.Longitude - 1.0})
	gotest.ASSERT_FALSE(aoi.IsAOIInside(aoi))
	gotest.ASSERT_TRUE(aoi.IsAOIInside(otherAOI))
}

func TestAOI_IsAOIInsideIncludingEdges(t *testing.T) {
	gotest.Initialize(gotest.TestImplement{Test: t})
	defer gotest.Cleanup()

	aoi := CreateAOI(Point{-50.0, -50.0}, Point{50.0, 50.0})
	gotest.ASSERT_TRUE(aoi.IsAOIInsideIncludeEdges(aoi))
	gotest.ASSERT_FALSE(aoi.IsAOIInside(aoi))
}
