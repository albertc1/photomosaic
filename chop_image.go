package main

import (
	"fmt"
	"image"
	"os"
	_ "image/gif"
	_ "image/png"
	 "image/jpeg"
	 "time"
)

func Chop_img(img image.Image, imgDb *ImageDb) image.Image {
	r_width := 10
	r_height := 10

	bounds := img.Bounds()

	horizontal_blocks := bounds.Max.X / r_width
	vertical_blocks := bounds.Max.Y / r_height

	dest_img := image.NewRGBA(image.Rect(0,0, bounds.Max.X, bounds.Max.Y))

	for y:= 0; y<vertical_blocks; y++ {
		for x :=0; x < horizontal_blocks; x++ {
			subimgrect := image.Rect(x*r_width, (y+1)*r_height, (x+1)*(r_width), y*r_height)
			r,g,b := Avg_color(img, subimgrect)
			fmt.Println("Input Rect rgb", r,g,b)
			ret := imgDb.Find(r,g,b)
			fmt.Println("Matched img", ret.Path, ret.R, ret.G, ret.B)
			Draw(dest_img, subimgrect, ret.Path)
		}
	}

	return dest_img
}

func MakeMosaic(imgUrl string, imgDb *ImageDb) bool {
    //fetch img from web
	img := fetch_img()
	dest_img:= Chop_img(img, imgDb)

	dstName := fmt.Sprintf("static/%d.jpg", time.Now().Unix())
	out, _ := os.Create(dstName)
	jpeg.Encode(out, dest_img, nil)
	return false
}

func fetch_img() image.Image{
	filepath := "/Users/udaysaraf/Downloads/IMAG0142.jpg"
    reader, err := os.Open(filepath)
    if err != nil {
        panic(err)
    }
    // lots of files have errors for some reason, so instead of killing the program, just ignore those files
    img, _, err := image.Decode(reader)
    return img
}
