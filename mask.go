package gmasklib

import "math"

// Tendency mask factory. The first arg is the generator which will send
// values to the mask. The second and the third arg are the low and high
// boundaries of the tendency mask. They can each be a single value or a
// generator. Finally an optional argument can fix a map
// exponent to apply a non-linear function to the mask output.
func MaskGen(gen Generator, params ...interface{}) Generator {
	var low Generator
	g1 := params[0]
	switch g1.(type) {
	case int:
		low = ConstGen(float64(g1.(int)))
	case float64:
		low = ConstGen(g1.(float64))
	case Generator:
		low = g1.(Generator)
	}
	var high Generator
	g2 := params[1]
	switch g2.(type) {
	case int:
		high = ConstGen(float64(g2.(int)))
	case float64:
		high = ConstGen(g2.(float64))
	case Generator:
		high = g2.(Generator)
	}
	var exp float64
	if len(params) > 2 {
		g3 := params[2]
		if _, ok := g3.(int); ok {
			exp = float64(g3.(int))
		} else {
			exp = g3.(float64)
		}
	}
	if exp != 0 {
		exp = math.Pow(2.0, exp)
		return func(t ...float64) float64 {
			min, max := low(t...), high(t...)
			value := gen(t...)
			return min + (max-min)*math.Pow(value, exp)
		}
	}
	return func(t ...float64) float64 {
		min, max := low(t...), high(t...)
		return min + (max-min)*gen(t...)
	}
}
