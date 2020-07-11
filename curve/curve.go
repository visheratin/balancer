package curve

import (
	"errors"
	"fmt"
	"image"
	"math/rand"

	"github.com/fogleman/gg"
	"github.com/visheratin/balancer/curve/hilbert"
	"github.com/visheratin/balancer/curve/morton"
)

//Curve is an interface of space filling curve realisation.
type Curve interface {
	Decode(code uint64) (coords []uint64, err error) //Decode returns coordinates for a given code(distance)
	DecodeWithBuffer(buf []uint64, code uint64) (coords []uint64, err error)
	Encode(coords []uint64) (code uint64, err error) //Encode returns code(distance) for a given set of coordinates
	DimensionSize() uint64                           // DimensionSize returns the maximum coordinate value in any dimension
	Length() uint64                                  // Length returns the maximum distance along curve
	Dimensions() uint64
	Bits() uint64
}

func NewCurve(cType CurveType, dims, bits uint64) (Curve, error) {
	switch cType {
	case Hilbert:
		return hilbert.New(dims, bits)
	case Morton:
		return morton.New(dims, bits)
	default:
		return nil, errors.New("unknown curve type")
	}
}

func DrawCurve(cType CurveType, bits uint64, op string) error {
	dims := uint64(2)
	c, err := NewCurve(cType, dims, bits)
	if err != nil {
		return err
	}
	dcSize := 2048
	dc := gg.NewContext(dcSize, dcSize)
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(10)
	maxSize := (1 << bits)
	cSize := float64(dcSize / maxSize)
	maxCode := uint64((1 << (dims * bits)) - 1)
	sx, sy := -1.0, -1.0
	for idx := uint64(0); idx <= maxCode; idx++ {
		cs, err := c.Decode(idx)
		if err != nil {
			return err
		}
		x := float64(cs[0])*cSize + cSize/2
		y := float64(cs[1])*cSize + cSize/2
		if sx != -1 {
			dc.DrawLine(sx, sy, x, y)
			dc.Stroke()
		}
		sx = x
		sy = y
	}
	dc.SavePNG(op)
	return nil
}

func DrawSplitCurve(c Curve, ranges []uint64) (image.Image, error) {
	dcSize := 2048
	dc := gg.NewContext(dcSize, dcSize)
	dc.SetRGB(1, 1, 1)
	dc.DrawRectangle(0, 0, float64(dcSize), float64(dcSize))
	dc.Fill()
	dc.SetLineWidth(7)
	maxSize := (1 << c.Bits())
	cSize := float64(dcSize / maxSize)
	maxCode := uint64((1 << (c.Dimensions() * c.Bits())) - 1)
	sx, sy := -1.0, -1.0
	si := 0
	r := rand.Float64()
	g := rand.Float64()
	b := rand.Float64()
	dc.SetRGB(r, g, b)
	fmt.Println(maxCode)
	for idx := uint64(0); idx <= maxCode; idx++ {
		if ranges != nil && idx > ranges[si] {
			r := rand.Float64()
			g := rand.Float64()
			b := rand.Float64()
			dc.SetRGB(r, g, b)
			si++
		}
		cs, err := c.Decode(idx)
		if err != nil {
			return nil, err
		}
		x := float64(cs[0])*cSize + cSize/2
		y := float64(cs[1])*cSize + cSize/2
		if sx != -1 {
			dc.DrawLine(sx, sy, x, y)
			dc.Stroke()
		}
		sx = x
		sy = y
	}
	return dc.Image(), nil
}
