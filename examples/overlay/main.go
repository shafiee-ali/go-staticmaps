// This is an example on how to use a map overlay.

package main

import (
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
	sm "github.com/shafiee-ali/go-staticmaps"
)

func main() {
	ctx := sm.NewContext()
	ctx.SetSize(1600, 1200)

	ctx.SetCenter(s2.LatLngFromDegrees(48.78110, -3.59638))
	ctx.SetZoom(15)

	// base map
	ctx.SetTileProvider(sm.NewTileProviderOpenStreetMaps())
	// OpenSeaMap as a overlay to the base map
	ctx.AddOverlay(sm.NewTileProviderOpenSeaMap())

	img, err := ctx.Render()
	if err != nil {
		panic(err)
	}

	if err := gg.SavePNG("overlay.png", img); err != nil {
		panic(err)
	}
}
