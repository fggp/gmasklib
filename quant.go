package gmask

import "math"

// Quantizer factory. The first argument is the generator which will send
// values to the quantizer. The second arg fixes the quantization grid interval.
// The optional third arg fixes the quantization strength (0..1) (defaults to 1),
// and an optional fourth arg fixes the offset (defaults to 0). Those three args
// can be each a single value or a generator.
func QuantGen(gen Generator, params ...interface{}) Generator {
	var q Generator
	g1 := params[0]
	switch g1.(type) {
	case int:
		q = ConstGen(float64(g1.(int)))
	case float64:
		q = ConstGen(g1.(float64))
	case Generator:
		q = g1.(Generator)
	}
	s := ConstGen(1.0)
	if len(params) > 1 {
		g2 := params[1]
		switch g2.(type) {
		case int:
			s = ConstGen(float64(g2.(int)))
		case float64:
			s = ConstGen(g2.(float64))
		case Generator:
			s = g2.(Generator)
		}
	}
	o := ConstGen(0.0)
	if len(params) == 3 {
		g3 := params[2]
		switch g3.(type) {
		case int:
			o = ConstGen(float64(g3.(int)))
		case float64:
			o = ConstGen(g3.(float64))
		case Generator:
			o = g3.(Generator)
		}
	}
	return func(t ...float64) float64 {
		val := gen(t...)
		delta := q(t...)
		strength := s(t...)
		offset := o(t...)
		var factor float64
		if strength >= 1.0 {
			factor = 0.0
		} else if strength <= 0.0 {
			factor = 1.0
		} else {
			factor = 1.0 - strength
		}
		val -= offset
		diff := math.Mod(val, delta)
		qval := val - diff
		if diff > delta/2.0 {
			qval += delta
		}
		if diff < -delta/2.0 {
			qval -= delta
		}
		diff = val - qval
		qval += diff * factor
		return qval + offset
	}
}
