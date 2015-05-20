package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/png"
	_ "image/jpeg"
)

func Chop_img(img image.Image) {
	r_width := 16
	r_height := 16

	bounds := img.Bounds()
	fmt.Println(bounds.Min.X, bounds.Max.X, bounds.Min.Y, bounds.Max.Y)

	horizontal_blocks := bounds.Max.X / r_width
	vertical_blocks := bounds.Max.Y / r_height
	for y:= 0; y<vertical_blocks; y++ {
		for x :=0; x < horizontal_blocks; x++ {
			subimgrect := image.Rect(x*r_width, (y+1)*r_height, (x+1)*(r_width), y*r_height)
			r,g,b := Avg_color(img, subimgrect)
			fmt.Println(r,g,b)
		}
	}
}

func MakeMosaic(imgUrl string, imgDb *ImageDb) bool {
    //fetch img from web
    //read/decode img from file
    // Chop_img(
	return false
}
