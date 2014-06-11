package main

import (
	"github.com/fggp/gmask"
	csnd6 "github.com/fggp/go-csnd6"
)

var orc string = `
sr     = 44100
ksmps  = 10
nchnls = 2

          instr 1
;p2 onset
;p3 duration
;p4 base frequency
;p5 fm index
;p6 pan (L=0, R=1)

kenv	expon		1,p3,0.01
kindx	expon		p5,p3,.4
a1	foscil	kenv*10000,p4,1,1.143,kindx,1
	outs		a1*(1-p6),a1*p6
	
          endin`

var sco string = `
f1 0 8193 10 1

f 0 20`

func events(cs csnd6.CSOUND) {
	f := gmask.NewField(0, 20)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	g := gmask.RndGen(gmask.UNI)
	i1 := gmask.NewInterpolation(1, false, false)
	i3 := gmask.NewInterpolation(3, false, false)

	seg1 := []float64{0.03, 0.5}
	seg2 := []float64{0.08, 1}
	m := gmask.MaskGen(g, gmask.BpfGen(seg1, i3), gmask.BpfGen(seg2, i3), 1)
	p.Num, p.Gen = 2, m
	f.AddParam(p)

	seg1[0], seg1[1] = 0.2, 3
	seg2[0], seg2[1] = 0.4, 5
	m = gmask.MaskGen(g, gmask.BpfGen(seg1, i1), gmask.BpfGen(seg2, i1))
	p.Num, p.Gen = 3, m
	f.AddParam(p)

	seg1[0], seg1[1] = 3000, 90
	seg2[0], seg2[1] = 5000, 150
	m = gmask.MaskGen(g, gmask.BpfGen(seg1, i1), gmask.BpfGen(seg2, i1), 1)
	q := gmask.QuantGen(m, gmask.BpfGen([]float64{400, 50}, nil), 0.95)
	p.Num, p.Gen = 4, q
	f.AddParam(p)

	seg1[0], seg1[1] = 2, 4
	seg2[0], seg2[1] = 4, 7
	m = gmask.MaskGen(g, gmask.BpfGen(seg1, nil), gmask.BpfGen(seg2, nil))
	p.Num, p.Gen = 5, m
	f.AddParam(p)

	p.Num, p.Gen = 6, gmask.RangeGen(0, 1)
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
}

func perform(cs csnd6.CSOUND, done chan bool) {
	cs.Perform()
	done <- true
}

func main() {
	cs := csnd6.Create(nil)
	cs.SetOption("-odac")
	cs.CompileOrc(orc)
	cs.ReadScore(sco)
	cs.Start()
	go events(cs)
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
