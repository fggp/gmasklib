package gmasklib

import (
	"bufio"
	"fmt"
	//"github.com/fggp/go-csnd"
	"io"
	"math"
	"strconv"
	"strings"
)

const SLICEINC = 10

// Create a new Field
func NewField(start, end float64) *Field {
	p := make([]Param, SLICEINC)
	f := Field{start, end, p}
	return &f
}

// Add a param to the field. This param corresponds to a pfield in Csound. If
// a param for that particular pfield existed already in the field, it will be
// overwritten.
func (f *Field) AddParam(p Param) {
	k := p.Num - 1
	if k >= len(f.Params) {
		n := len(f.Params)
		for n <= k {
			n += SLICEINC
		}
		n -= len(f.Params)
		f.Params = append(f.Params, make([]Param, n)...)
	}
	f.Params[k] = p
}

// Create a new param corresponding to a pfield in Csound. The first arg is the
// pfield number (1..). If this param uses a generator daisy chain, the gen arg
// has to be the last generator in the chain. The prec arg is only used when
// evaluating the field to a Csound score.
func NewParam(num int, gen Generator, prec int) Param {
	if prec < 0 || prec > 5 {
		prec = 5
	}
	p := Param{Num: num, Gen: gen, Prec: prec}
	return p
}

// Returns the pfield value at time t, between t0 and t1.
func (p Param) Value(t, t0, t1 float64) float64 {
	return p.Gen(t, t0, t1)
}

// Create a new interpolation value.
func NewInterpolation(val float64, cos, off bool) *Interpolation {
	ipl := Interpolation{val, cos, off}
	return &ipl
}

func pVal(p Param, t, start, end float64) float64 {
	v := p.Value(t, start, end)
	prec := p.Prec
	if prec < 0 || prec > 5 {
		return v
	}
	m := math.Pow(10.0, float64(prec))
	return math.Floor(v*m+0.5) / m
}

/*
// Evaluate a field generating score events sent to Csound via the API
// scoreEvent or scoreEventAbsolute functions.
func (f *Field) EvalToScoreEvents(cs csnd.CSOUND, absolute bool, timeOfs float64) {
	t := f.Start
	pFields := make([]csnd.MYFLT, len(f.Params))
	for {
		pFields[0] = csnd.MYFLT(pVal(f.Params[0], t, f.Start, f.End))
		pFields[1] = csnd.MYFLT(t)
		for i := 2; i < len(f.Params); i++ {
			if f.Params[i].Num == i+1 {
				pFields[i] = csnd.MYFLT(pVal(f.Params[i], t, f.Start, f.End))
			} else {
				break
			}
		}
		if absolute {
			cs.ScoreEventAbsolute('i', pFields, timeOfs)
		} else {
			cs.ScoreEvent('i', pFields)
		}
		if t += pVal(f.Params[1], t, f.Start, f.End); t > f.End {
			break
		}
	}
}
*/

func pFmt(format string, prec int) string {
	return strings.Replace(format, "p", strconv.Itoa(prec), -1)
}

// Evaluate a field as a score section. The result is written into an io.Writer.
// This procedure is generally invoked from the parser in the gmask program.
// But one can use it from any go program, if the Field receiver pointer points
// to a valid Field structure.
func (f *Field) EvalToScore(dest io.Writer, fieldNum int) {
	var nEvents int
	w := bufio.NewWriter(dest)
	fmt.Fprintf(w, "\n; ------- begin of field %d --- beats: %.2f - %.2f --------\n\n",
		fieldNum, f.Start, f.End)
	t := f.Start
	for {
		fmt.Fprintf(w, pFmt("i%.pf ", f.Params[0].Prec), f.Params[0].Value(t, f.Start, f.End))
		fmt.Fprintf(w, pFmt("%.pf ", f.Params[1].Prec), t)
		for i := 2; i < len(f.Params); i++ {
			if f.Params[i].Num == i+1 {
				fmt.Fprintf(w, pFmt("%.pf ", f.Params[i].Prec), f.Params[i].Value(t, f.Start, f.End))
			} else {
				break
			}
		}
		fmt.Fprintf(w, "\n")
		w.Flush()
		nEvents++
		if t += f.Params[1].Value(t, f.Start, f.End); t > f.End {
			break
		}
	}
	fmt.Fprintf(w, "\n; ------- end of field %d --- number of events: %d -------\n",
		fieldNum, nEvents)
	w.Flush()
}
