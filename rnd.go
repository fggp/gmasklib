package gmasklib

import (
	"math"
	"math/rand"
)

// Random number generator factory.
func RangeGen(min, max float64) Generator {
	delta := max - min
	return func(t ...float64) float64 {
		return rand.Float64()*delta + min
	}
}

// Probability distribution generator factory. The first arg specifies the
// kind of distribution used. The params args depend on the distribution:
// UNI, LIN, RLIN, and TRI have no param.
//
// EXP, REXP, and BEXP can have one param that is
// a single value or a generator. This param has to be
// greater than 0, it defaults to 1.0.
//
// GAUSS, CAUCHY, BETA, and WEI can have to params that are each a single
// value or a generator. For GAUSS, those params are standard
// deviation and mean (0..1), and they default to 0.1 and 0.5. For CAUCHY, those
// params are spread and mean (0..1) and they default to 0.1 and 0.5. For BETA,
// those params are a and b (0..1), and they both default to 0.1. For WEI, those
// params are s (0..1) and t (>0), and they default to 0.5 and 2.
func RndGen(mode RndMode, params ...interface{}) Generator {
	switch mode {
	case UNI:
		return func(t ...float64) float64 {
			return rand.Float64()
		}
	case LIN:
		return func(t ...float64) float64 {
			a := rand.Float64()
			b := rand.Float64()
			if a < b {
				return a
			} else {
				return b
			}
		}
	case RLIN:
		return func(t ...float64) float64 {
			a := rand.Float64()
			b := rand.Float64()
			if a > b {
				return a
			} else {
				return b
			}
		}
	case TRI:
		return func(t ...float64) float64 {
			a := rand.Float64()
			b := rand.Float64()
			return 0.5 * (a + b)
		}
	}
	var g1 Generator
	if len(params) == 0 {
		switch mode {
		case EXP, REXP, BEXP:
			g1 = ConstGen(1.0)
		case GAUSS, CAUCHY, BETA:
			g1 = ConstGen(0.1)
		case WEI:
			g1 = ConstGen(0.5)
		}
	} else {
		g := params[0]
		switch g.(type) {
		case int:
			g1 = ConstGen(float64(g.(int)))
		case float64:
			g1 = ConstGen(g.(float64))
		case Generator:
			g1 = g.(Generator)
		}
	}
	switch mode {
	case EXP:
		return func(t ...float64) float64 {
			lambda := g1(t...)
			for {
				x := rand.ExpFloat64() / lambda
				if x < 1.0 {
					return x
				}
			}
		}
	case REXP:
		return func(t ...float64) float64 {
			lambda := g1(t...)
			for {
				x := 1.0 - rand.ExpFloat64()/lambda
				if x > 0.0 {
					return x
				}
			}
		}
	case BEXP:
		return func(t ...float64) float64 {
			lambda := g1(t...)
			var x float64
			for {
				x = rand.ExpFloat64() / lambda
				if x < 1.0 {
					break
				}
			}
			if rand.Float64() > 0.5 {
				return 0.5 + x/2.0
			} else {
				return 0.5 - x/2.0
			}
		}
	}
	var g2 Generator
	if len(params) < 2 {
		switch mode {
		case GAUSS, CAUCHY:
			g2 = ConstGen(0.5)
		case BETA:
			g2 = ConstGen(0.1)
		case WEI:
			g2 = ConstGen(2.0)
		}
	} else {
		g := params[1]
		switch g.(type) {
		case int:
			g2 = ConstGen(float64(g.(int)))
		case float64:
			g2 = ConstGen(g.(float64))
		case Generator:
			g2 = g.(Generator)
		}
	}
	switch mode {
	case GAUSS:
		return func(t ...float64) float64 {
			sigma, mu := g1(t...), g2(t...)
			for {
				x := rand.NormFloat64()*sigma + mu
				if x >= 0.0 && x <= 1.0 {
					return x
				}
			}
		}
	case CAUCHY:
		return func(t ...float64) float64 {
			var x float64
			alpha, mu := g1(t...), g2(t...)
			for {
				for {
					x = rand.Float64()
					if x != 0.5 {
						break
					}
				}
				x = alpha*math.Tan(x*math.Pi) + mu
				if x >= 0.0 && x <= 1 {
					return x
				}
			}
		}
	case BETA:
		return func(t ...float64) float64 {
			a, b := g1(t...), g2(t...)
			for {
				x1 := math.Pow(rand.Float64(), 1.0/a)
				x2 := math.Pow(rand.Float64(), 1.0/b)
				sum := x1 + x2
				if sum > 1.0 {
					return x1 / sum
				}
			}
		}
	case WEI:
		return func(t ...float64) float64 {
			k, lambda := g1(t...), g2(t...)
			for {
				a := 1.0 / (1.0 - rand.Float64())
				x := k * math.Pow(math.Log(a), 1.0/lambda)
				if x <= 1.0 {
					return x
				}
			}
		}
	}
	return nil
}
