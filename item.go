package gmask

import "math/rand"

// Item generator factory. The first arg is the item mode. The second
// arg is the list of values from which the generator will pick its return
// value.
func ItemGen(mode ItemMode, list []float64) Generator {
	var i int
	var inc int = 1
	n := len(list)
	switch mode {
	case CYCLE:
		return func(t ...float64) float64 {
			x := list[i]
			if i += inc; i >= n {
				i = 0
			}
			return x
		}
	case SWING:
		return func(t ...float64) float64 {
			x := list[i]
			if i += inc; i >= n || i < 0 {
				inc = -inc
				i += inc
				i += inc
			}
			return x
		}
	case HEAP:
		perm := rand.Perm(n)
		return func(t ...float64) float64 {
			x := list[perm[i]]
			if i += inc; i >= n {
				i = 0
				perm = rand.Perm(n)
			}
			return x
		}
	case RANDOM:
		return func(t ...float64) float64 {
			return list[rand.Intn(n)]
		}
	}
	return nil
}
