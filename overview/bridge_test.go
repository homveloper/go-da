package overview_test

import (
	"log"
	"testing"
)

type IDrawShape interface {
	drawShape(x [5]float32, y [5]float32)
}

type DrawShape struct {
}

func (d *DrawShape) drawShape(x [5]float32, y [5]float32) {
	log.Println("Drawing shape")
}

type IContour interface {
	drawContour(x [5]float32, y [5]float32)
	resizeByFactor(factor int)
}

type DrawContour struct {
	x      [5]float32
	y      [5]float32
	shape  DrawShape
	factor int
}

func (d *DrawContour) drawContour(x [5]float32, y [5]float32) {
	log.Println("Drawing contour")
	d.shape.drawShape(x, y)
}

func (d *DrawContour) resizeByFactor(factor int) {
	log.Println("Resizing contour")
	d.factor = factor
}

func TestBridge(t *testing.T) {
	var xs = [5]float32{1, 2, 3, 4, 5}
	var ys = [5]float32{1, 2, 3, 4, 5}

	var contour IContour = &DrawContour{xs, ys, DrawShape{}, 2}
	contour.drawContour(xs, ys)
	contour.resizeByFactor(3)
}
