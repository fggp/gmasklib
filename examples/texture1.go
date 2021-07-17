package main

import (
	"github.com/fggp/gmasklib"
	"github.com/fggp/go-csnd"
)

var orc string = `
sr     = 44100
ksmps  = 10
nchnls = 2

instr 1

;p4 frequency
;p5 pan (0...1) 

ipanl	table	1-p5 ,1,1
ipanr	table	p5 ,1,1

k1	expon	1,p3,.01
a1	foscil	k1*4200,p4,1,2.41,k1*6,2
	
	outs	a1*ipanl, a1*ipanr

endin`

var sco string = `
f1 0 8192 9 .25 1 0
f2 0 8193 10 1 

f 0 33`

func events1(cs csnd.CSOUND, ready chan bool) {
	f := gmasklib.NewField(0, 30)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.UNI)
	m := gmasklib.MaskGen(g, gmasklib.BpfGen([]float64{0.01, 0.002}, nil),
		gmasklib.BpfGen([]float64{0.1, 0.01}, nil))
	p.Num, p.Gen = 2, m
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.RangeGen(0.5, 1)
	f.AddParam(p)

	b1 := gmasklib.BpfGen([]float64{860, 80}, gmasklib.NewInterpolation(-1.2, false, false))
	b2 := gmasklib.BpfGen([]float64{940, 2000}, gmasklib.NewInterpolation(1, false, false))
	m = gmasklib.MaskGen(g, b1, b2, 1)
	q := gmasklib.QuantGen(m, 100, 0.9, 0)
	p.Num, p.Gen = 4, q
	f.AddParam(p)

	b1 = gmasklib.BpfGen([]float64{0.4, 0}, nil)
	b2 = gmasklib.BpfGen([]float64{0.6, 1}, nil)
	p.Num, p.Gen = 5, gmasklib.MaskGen(g, b1, b2)
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
	ready <- true
}

func events2(cs csnd.CSOUND, ready chan bool) {
	f := gmasklib.NewField(31, 33)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	g := gmasklib.BpfGen([]float64{0.08, 0.8}, gmasklib.NewInterpolation(2, false, false))
	p.Num, p.Gen = 2, g
	f.AddParam(p)

	g = gmasklib.BpfGen([]float64{0.1, 2}, nil)
	p.Num, p.Gen = 3, g
	f.AddParam(p)

	p.Num, p.Gen = 4, gmasklib.RangeGen(300, 400)
	f.AddParam(p)

	g = gmasklib.BpfGen([]float64{0, 1}, nil)
	p.Num, p.Gen = 5, g
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
	ready <- true
}

func perform(cs csnd.CSOUND, done chan bool) {
	cs.Perform()
	done <- true
}

func main() {
	cs := csnd.Create(nil)
	cs.SetOption("-odac")
	cs.CompileOrc(orc)
	cs.ReadScore(sco)
	cs.Start()
	ready := make(chan bool, 3)
	go events1(cs, ready)
	go events2(cs, ready)
	for i := 1; i <= 2; i++ {
		<-ready
	}
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
