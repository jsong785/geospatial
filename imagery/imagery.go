package imagery

import (
	"geospatial"
        "osm"
        "tile"
)

const OSM_ROOT_DIR = "/home/.GMapCatcher/OSM_tiles"

type ImageryData struct {
    tile.TileData
    aoi geospatial.AOI
    LoadPath[] string
}

type ImageryLogic struct {
    tile.TileLogic
}

func (t ImageryData) GetAOI() geospatial.AOI {
    return t.aoi
}

func createLoadPathStrings(data[] tile.TileData, level int) []string {
    paths := make([]string, 0)

    for _, d := range data {
        osmTile:= osm.GetTile(d.GetAOI().Lower, level)
        paths = append(paths, osm.GetDirectory(osmTile, level, OSM_ROOT_DIR))
    }

    return paths
}

func GetImages(t tile.Tile, aoi geospatial.AOI) []string {
    tile := t.FindArea(aoi)
    if tile != nil {
        return createLoadPathStrings(tile.Data, tile.CalculateLevel())
    }
    return nil
}
