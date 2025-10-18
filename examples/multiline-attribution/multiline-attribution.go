// This is an example on how to use a multiline attribution string.

package main

import (
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
	sm "github.com/shafiee-ali/go-staticmaps"
)

func main() {
	ctx := sm.NewContext()
	ctx.SetSize(400, 300)
	ctx.OverrideAttribution("This is a\nmulti-line\nattribution string.")
	ctx.SetCenter(s2.LatLngFromDegrees(48, 7.9))
	ctx.SetZoom(13)

	img, err := ctx.Render()
	if err != nil {
		panic(err)
	}

	if err := gg.SavePNG("multiline-attribution.png", img); err != nil {
		panic(err)
	}
}
