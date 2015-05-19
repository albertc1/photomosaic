package main

import (
    "image"
    // "image/draw"
)

func Draw(dst image.Image, dstRect image.Rectangle, srcPath string) {
    src := ReadImageFromFile(srcPath)
    draw.Draw(dst, dstRect, src, image.Image.ZP, draw.Src)
}
