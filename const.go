package gmask

// Constant generator factory. The argument is the constant value to be returned
// by the generator.
func ConstGen(v float64) Generator {
	return func(t ...float64) float64 {
		return v
	}
}
