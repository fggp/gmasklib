package gmasklib

type Interpolator func(t, t0, v0, t1, v1 float64) float64

// Type of the function returned by the different generator factories. A
// generator can be called with one or three float64 parameters. If the
// generator is a segment function defined with one single segment, the
// generator needs three float64 params: the current time, the start time,
// and the end time of the segment. If one of the generators in a daisy
// chain is a single segment generator, then the first generator in the
// chain has to be called with those three parameters. The other generators
// can be called with an ellipsis arg: gen(t...).
// If the generator is not a single segment generator or if it is in a daisy
// chain without any single segment generator, then it needs only one
// float64 parameter, the current time.
type Generator func(t ...float64) float64

type ItemMode int

// Modes for the item generator
const (
	CYCLE ItemMode = iota
	SWING
	HEAP
	RANDOM
)

type Interpolation struct {
	ipl float64
	cos bool
	off bool
}

type RndMode int

// Modes for the rnd generator
const (
	UNI RndMode = iota
	LIN
	RLIN
	TRI
	EXP
	REXP
	BEXP
	GAUSS
	CAUCHY
	BETA
	WEI
)

type OscMode int

// Modes for the osc generator
const (
	SIN OscMode = iota
	COS
	SQUARE
	TRIANGLE
	SAWUP
	SAWDOWN
	POWUP
	POWDOWN
)

type AccumMode int

// Modes for the accumulator
const (
	ON AccumMode = iota
	LIMIT
	MIRROR
	WRAP
)

type Field struct {
	Start, End float64
	Params     []Param
}

type Param struct {
	Num  int
	Gen  Generator
	Prec int
}
