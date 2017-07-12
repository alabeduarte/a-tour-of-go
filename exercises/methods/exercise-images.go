// https://tour.golang.org/methods/25

package main

import(
  "golang.org/x/tour/pic"
  "image"
  "image/color"
)

type Image struct{
  Width, Height int
}

func (myImage *Image) ColorModel() color.Model {
  return color.RGBAModel
}

func (myImage *Image) Bounds() image.Rectangle {
  return image.Rect(0, 0, myImage.Width, myImage.Height);
}

func (myImage *Image) At(x, y int) color.Color {
  return color.RGBA{uint8(x), uint8(y), 255, 255};
}

func main() {
  m := Image{100, 200}
  pic.ShowImage(&m)
}
