package optimizer

import (
	"math"
	"sort"

	"github.com/pkg/errors"
	balancer "github.com/visheratin/balancer"
)

func RangeOptimizer(s *balancer.Space) (res []*balancer.CellGroup, err error) {
	totalPower := s.TotalPower()
	cgs := s.CellGroups()
	if len(cgs) == 0 {
		return res, nil
	}
	var max, min uint64
	sort.Slice(cgs, func(i, j int) bool { return cgs[i].Node().ID() < cgs[j].Node().ID() })
	for iter := 0; iter < len(cgs); iter++ {
		min = max
		p := cgs[iter].Node().Power().Get() / totalPower
		max = min + uint64(math.Round(float64(s.TotalCells())*p))
		if err := cgs[iter].SetRange(min, max); err != nil {
			return nil, errors.Wrap(err, "range optimizer error")
		}
	}
	if max < s.TotalCells() {
		if err := cgs[len(cgs)-1].SetRange(min, s.TotalCells()+1); err != nil {
			return nil, errors.Wrap(err, "range optimizer error")
		}
	}
	cells := s.Cells()
	for iter := range cells {
		for cgiter := range cgs {
			if cgs[cgiter].InRange(cells[iter].ID()) {
				cgs[cgiter].AddCell(cells[iter], true)
				break
			}
		}
	}
	return cgs, nil
}

func PowerRangeOptimizer(s *balancer.Space) (res []*balancer.CellGroup, err error) {
	cells := s.Cells()
	totalPower := s.TotalPower()
	//totalFree := s.TotalCapacity() - float64(s.TotalLoad())
	cgs := s.CellGroups()
	if len(cgs) == 0 {
		return res, nil
	}
	var max, min uint64

	sort.Slice(cgs, func(i, j int) bool {
		return (cgs[i].Node().Capacity().Get() - float64(cgs[i].TotalLoad())) < (cgs[j].Node().Capacity().Get() - float64(cgs[j].TotalLoad()))
	})

	for iter := 0; iter < len(cgs); iter++ {
		min = max
		p := cgs[iter].Node().Power().Get() / totalPower
		f := cgs[iter].Node().Capacity().Get()
		max = min + uint64(math.Round(float64(s.TotalCells())*p))

		for citer := 0; citer < len(cells); citer++ {
			if cells[citer].ID() > max {
				break
			}
			if cells[citer].ID() >= min {
				f -= float64(cells[citer].Load())
				if f <= 0 {
					c := citer - 1
					if c < 0 {
						c = 0
					}
					max = cells[citer].ID()
					break
				}
				cgs[iter].AddCell(cells[citer], true)
			}
		}
		if err := cgs[iter].SetRange(min, max); err != nil {
			return nil, errors.Wrap(err, "power range optimizer error")
		}
	}

	if max < s.TotalCells() {
		if err := cgs[len(cgs)-1].SetRange(min, s.TotalCells()+1); err != nil {
			return nil, errors.Wrap(err, "range optimizer error")
		}
		for citer := range cells {
			if cells[citer].ID() >= max {
				cgs[len(cgs)-1].AddCell(cells[citer], true)
			}
		}
	}

	return cgs, nil
}
