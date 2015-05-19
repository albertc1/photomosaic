package main

import (
	"image"
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"
)

func Avg_color(img image.Image, r image.Rectangle) (int,int,int) {
	

	pixels := 0
	green := 0
	red := 0 
	blue := 0

	for y:= r.Min.Y; y<r.Max.Y; y++ {
		for x :=r.Min.X; x < r.Max.X; x++ {
			pixels += 1 
			r, g, b, _ := img.At(x, y).RGBA()
			r = r >> 8
			g = g >> 8
			b = b >> 8
			red += int(r)
			green += int(g)
			blue += int(b)
		}
	}

	return red/pixels, green/pixels, blue/pixels
	
}