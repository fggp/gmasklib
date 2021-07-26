package main

import (
	"bytes"
	"github.com/fggp/gmasklib"
	"github.com/fggp/go-csnd"
)

var orc string = `
sr     = 44100
ksmps  = 10
nchnls = 2
0dbfs  = 1.0

  instr 1
;p2 onset
;p3 duration
;p4 base frequency
;p5 fm index
;p6 pan (L=0, R=1)

kenv  = expon(1, p3, 0.01)
kindx = expon(p5, p3, 0.4)
a1    = foscil(kenv*0.31, p4, 1, 1.143, kindx, 1)
        outs a1*(1-p6), a1*p6
  endin
`

var sco string = `
f1 0 8193 10 1

f 0 20
`

func events() string {
	f := gmasklib.NewField(0, 20)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 2)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.UNI)
	i1 := gmasklib.NewInterpolation(1, gmasklib.IPLNUM)
	i3 := gmasklib.NewInterpolation(3, gmasklib.IPLNUM)

	seg1 := []float64{0.03, 0.5}
	seg2 := []float64{0.08, 1}
	m := gmasklib.MaskGen(g, gmasklib.BpfGen(seg1, i3), gmasklib.BpfGen(seg2, i3), 1)
	p.Num, p.Gen = 2, m
	f.AddParam(p)

	seg1[0], seg1[1] = 0.2, 3
	seg2[0], seg2[1] = 0.4, 5
	m = gmasklib.MaskGen(g, gmasklib.BpfGen(seg1, i1), gmasklib.BpfGen(seg2, i1))
	p.Num, p.Gen = 3, m
	f.AddParam(p)

	seg1[0], seg1[1] = 3000, 90
	seg2[0], seg2[1] = 5000, 150
	m = gmasklib.MaskGen(g, gmasklib.BpfGen(seg1, i1), gmasklib.BpfGen(seg2, i1), 1)
	q := gmasklib.QuantGen(m, gmasklib.BpfGen([]float64{400, 50}, nil), 0.95)
	p.Num, p.Gen = 4, q
	f.AddParam(p)

	seg1[0], seg1[1] = 2, 4
	seg2[0], seg2[1] = 4, 7
	m = gmasklib.MaskGen(g, gmasklib.BpfGen(seg1, nil), gmasklib.BpfGen(seg2, nil))
	p.Num, p.Gen = 5, m
	f.AddParam(p)

	p.Num, p.Gen = 6, gmasklib.RangeGen(0, 1)
	f.AddParam(p)

	var buf bytes.Buffer
	f.EvalToScore(&buf, 1)
	return buf.String()
}

func perform(cs csnd.CSOUND, done chan bool) {
	cs.Perform()
	done <- true
}

func main() {
	cs := csnd.Create(nil)
	cs.SetOption("-odac")
	cs.CompileOrc(orc)
	cs.ReadScore(sco + events())
	cs.Start()
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
