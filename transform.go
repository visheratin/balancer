package balancer

import "github.com/visheratin/balancer/curve"

type TransformFunc func(values []interface{}, sfc curve.Curve) ([]uint64, error)
