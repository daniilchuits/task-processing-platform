package jpg

import (
	"fmt"
	"image"
	"sort"
	"strings"
)

type rgb struct {
	r uint8
	g uint8
	b uint8
}

type colorCount struct {
	rgb   rgb
	count int
}

func process(image image.Image) string {

	colors := make(map[rgb]int)

	for x := image.Bounds().Min.X; x < image.Bounds().Max.X; x++ {
		for y := image.Bounds().Min.Y; y < image.Bounds().Max.Y; y++ {

			r32, g32, b32, _ := image.At(x, y).RGBA()

			r8 := uint8(r32 >> 8)
			g8 := uint8(g32 >> 8)
			b8 := uint8(b32 >> 8)

			bucket := rgb{
				r: r8 / 16,
				g: g8 / 16,
				b: b8 / 16,
			}

			colors[bucket]++

		}
	}

	var ccs []colorCount

	for rgb, integ := range colors {
		ccs = append(ccs, colorCount{
			rgb:   rgb,
			count: integ,
		})
	}

	sort.Slice(ccs, func(i, j int) bool {
		return ccs[i].count > ccs[j].count
	})

	var sb strings.Builder

	for i := 0; i < 3 && i < len(ccs); i++ {

		r := ccs[i].rgb.r*16 + 8
		g := ccs[i].rgb.g*16 + 8
		b := ccs[i].rgb.b*16 + 8

		fmt.Fprintf(&sb,
			"|#%02X%02X%02X|",
			r, g, b,
		)
	}
	return sb.String()
}
