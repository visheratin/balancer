package curve

import (
	"fmt"
	"github.com/struckoff/SFCFramework/curve/hilbert"
	"github.com/struckoff/SFCFramework/curve/morton"
)

type Curve interface {
	Decode(code uint64) (coords []uint64, err error) //Decode returns coordinates for a given code(distance)
	DecodeWithBuffer(buf []uint64, code uint64) (coords []uint64, err error)
	Encode(coords []uint64) (code uint64, err error) //Encode returns code(distance) for a given set of coordinates
	Size() uint                                      // Size returns the maximum coordinate value in any dimension
	MaxDistance() uint64                             // MaxDistance returns the maximum distance along curve
}

func NewCurve(cType CurveType, dims, bits uint64) (Curve, error) {
	switch cType {
	case Hilbert:
		return hilbert.New(dims, bits)
	case Morton:
		return morton.New(dims, bits)
	default:
		return nil, fmt.Errorf("curve type %v not supported", cType)
	}
}
