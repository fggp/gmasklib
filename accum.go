package gmask

// Accumulator factory. The first argument is the generator which will send
// values to the accumulator. The second argument is the accumulator mode. The
// ON mode has no limits arguments. The other modes have a low and a high limit
// arg which can be each a single value or a generator. Finally
// an optional argument can fix the initial value of the accumulator.
func AccumGen(gen Generator, params ...interface{}) Generator {
	var sum float64
	var mode = params[0].(AccumMode)
	if mode == ON {
		if len(params) == 2 {
			sum = params[1].(float64)
		}
		return func(t ...float64) float64 {
			sum += gen(t...)
			return sum
		}
	}
	var low Generator
	g := params[1]
	switch g.(type) {
	case int:
		low = ConstGen(float64(g.(int)))
	case float64:
		low = ConstGen(g.(float64))
	case Generator:
		low = g.(Generator)
	}
	var high Generator
	g = params[2]
	switch g.(type) {
	case int:
		high = ConstGen(float64(g.(int)))
	case float64:
		high = ConstGen(g.(float64))
	case Generator:
		high = g.(Generator)
	}
	if len(params) == 4 {
		sum = params[3].(float64)
	}
	switch mode {
	case LIMIT:
		return func(t ...float64) float64 {
			min, max := low(t...), high(t...)
			sum += gen(t...)
			if sum > max {
				sum = max
			} else if sum < min {
				sum = min
			}
			return sum
		}
	case MIRROR:
		return func(t ...float64) float64 {
			min, max := low(t...), high(t...)
			sum += gen(t...)
			if sum > max {
				sum = max - sum + max
			} else if sum < min {
				sum = min + min - sum
			}
			return sum
		}
	case WRAP:
		return func(t ...float64) float64 {
			min, max := low(t...), high(t...)
			sum += gen(t...)
			if sum > max {
				sum = min + sum - max
			} else if sum < min {
				sum = max - min + sum
			}
			return sum
		}
	}
	return nil
}
