package main

import ( 
	"golang.org/x/tour/pic"
	"image/color"
	"image"
)

type Image struct{
	w,h int
	cdet [][]uint8
}

// ColorModel returns the Image's color model.
func (im Image) ColorModel() color.Model{
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (im Image) Bounds() image.Rectangle{
	//def: func Rect(x0, y0, x1, y1 int) Rectangle
	return image.Rect(0,0,im.w,im.h)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (im Image) At(x, y int) color.Color{
	v := im.cdet[x][y]
	return color.RGBA{v, v, 255, 255}
}

func Pic(dx, dy int) [][]uint8 {
	ans := make([][]uint8, dy)
	for y:=0; y < dy; y++{
		ans[y] = make([]uint8, dx)
		for x:=0; x < dx; x++ {
			ans[y][x] = (uint8(x)+uint8(y))/uint8(2)
			ans[y][x] = uint8(x)*uint8(y)
			ans[y][x] = (uint8(x) ^ uint8(y)) * (uint8(x) ^ uint8(y))
		}
	}
	return ans
}


func main() {
	m := Image{256,256,nil}
	m.cdet = Pic(256,256)
	pic.ShowImage(m)
}
