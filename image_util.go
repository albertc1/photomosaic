package main

import (
    "image"
    "image/draw"
)

func Draw(dst draw.Image, dstRect image.Rectangle, srcPath string) {
    src := ReadImageFromFile(srcPath)
    draw.Draw(dst, dstRect, src, image.ZP, draw.Src)
}
