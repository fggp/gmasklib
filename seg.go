package gmask

import "math"

// Segment function factory. Returns a generator.
//
// The float64 slice describes a sequence of time-value
// pairs if its length is greater than two, or one segment lasting from start
// to end of field. If no interpolation is specified, the interp pointer
// should be nil.
//
// If the float64 slice represents a single
// segment, the generator has to be called with three float64 params: current
// time, start time and end time of the segment. Otherwise, only the current
// time is needed.
func BpfGen(points []float64, interp *Interpolation) Generator {
	var interpolate Interpolator
	if interp == nil || interp.off {
		interpolate = func(t, t0, v0, t1, v1 float64) float64 {
			return v0
		}
	} else if interp.cos {
		interpolate = func(t, t0, v0, t1, v1 float64) float64 {
			x := ((t-t0)/(t1-t0))*math.Pi + math.Pi
			return v0 + ((v1 - v0) * (1.0 + math.Cos(x)) / 2.0)
		}
	} else if interp.ipl == 0 {
		interpolate = func(t, t0, v0, t1, v1 float64) float64 {
			return v0 + ((t-t0)/(t1-t0))*(v1-v0)
		}
	} else {
		exp := interp.ipl
		interpolate = func(t, t0, v0, t1, v1 float64) float64 {
			r := (t - t0) / (t1 - t0)
			if v0 == v1 {
				return v0
			} else if exp > 0 {
				if v1 >= v0 {
					return v0 + math.Pow(r, 1.0+exp)*(v1-v0)
				} else {
					return v1 + math.Pow(1.0-r, 1.0+exp)*(v0-v1)
				}
			} else {
				if v1 >= v0 {
					return v1 + math.Pow(1.0-r, 1.0-exp)*(v0-v1)
				} else {
					return v0 + math.Pow(r, 1.0-exp)*(v1-v0)
				}
			}
		}
	}
	if len(points) == 2 {
		v1, v2 := points[0], points[1]
		return func(t ...float64) float64 {
			return interpolate(t[0], t[1], v1, t[2], v2)
		}
	} else if len(points) > 2 && len(points)%2 == 0 {
		return func(t ...float64) float64 {
			if t[0] < points[0] {
				return points[1]
			}
			t0, v0 := points[0], points[1]
			var i int
			var t1, v1 float64
			for i = 2; i < len(points); i += 2 {
				t1, v1 = points[i], points[i+1]
				if t[0] < t1 {
					break
				}
				t0, v0 = t1, v1
			}
			if i == len(points) {
				return v0
			}
			return interpolate(t[0], t0, v0, t1, v1)
		}
	} else {
		return func(t ...float64) float64 {
			return -1.0
		}
	}
}
