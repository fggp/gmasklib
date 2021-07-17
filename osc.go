package gmasklib

import "math"

// Oscillator factory. The first arg is the oscillator mode. The second arg is
// the oscillator frequency. It can be a single value or a
// generator. An optional third argument can fix the initial phase of the
// oscillator. For the power function oscillators, an optional fourth argument
// can fix the exponent which defaults to 0.
func OscGen(mode OscMode, params ...interface{}) Generator {
	var freq Generator
	g1 := params[0]
	switch g1.(type) {
	case int:
		freq = ConstGen(float64(g1.(int)))
	case float64:
		freq = ConstGen(g1.(float64))
	case Generator:
		freq = g1.(Generator)
	}
	var phs float64
	if len(params) > 1 {
		g2 := params[1]
		if _, ok := g2.(int); ok {
			phs = float64(g2.(int))
		} else {
			phs = g2.(float64)
		}
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
	switch mode {
	case SIN:
		return func(t ...float64) float64 {
			x := math.Sin(2*math.Pi*freq(t...)*t[0] + phs)
			return x*0.5 + 0.5
		}
	case COS:
		return func(t ...float64) float64 {
			x := math.Cos(2*math.Pi*freq(t...)*t[0] + phs)
			return x*0.5 + 0.5
		}
	case SQUARE:
		phs /= (2 * math.Pi)
		return func(t ...float64) float64 {
			_, x := math.Modf(freq(t...)*t[0] + phs)
			if x = math.Abs(x); x < 0.5 {
				return 1.0
			}
			return 0.0
		}
	case TRIANGLE:
		phs /= (2 * math.Pi)
		return func(t ...float64) float64 {
			_, x := math.Modf(freq(t...)*t[0] + phs)
			if x = math.Abs(x); x < 0.5 {
				return 2.0 * x
			}
			return 2.0 * (1.0 - x)
		}
	case SAWUP:
		phs /= (2 * math.Pi)
		return func(t ...float64) float64 {
			_, x := math.Modf(freq(t...)*t[0] + phs)
			return math.Abs(x)
		}
	case SAWDOWN:
		phs /= (2 * math.Pi)
		return func(t ...float64) float64 {
			_, x := math.Modf(freq(t...)*t[0] + phs)
			return 1.0 - math.Abs(x)
		}
	case POWUP:
		phs /= (2 * math.Pi)
		exp = math.Pow(2, exp)
		return func(t ...float64) float64 {
			_, x := math.Modf(freq(t...)*t[0] + phs)
			return math.Pow(math.Abs(x), exp)
		}
	case POWDOWN:
		phs /= (2 * math.Pi)
		exp = math.Pow(2, exp)
		return func(t ...float64) float64 {
			_, x := math.Modf(freq(t...)*t[0] + phs)
			return math.Pow(1.0-math.Abs(x), exp)
		}
	}
	return nil
}
