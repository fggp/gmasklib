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

k1	oscil	8000*p5,1/p3,1
a1	oscil	k1,p4,2
	outs	a1*(1-p6),a1*p6

endin`

var sco string = `
f1 0 8193 8 0 4096 1 4096 0
f2 0 8193 10 1 .5 .3 .2 .1

f 0 10`

func events1(cs csnd.CSOUND, ready chan bool) {
	f := gmasklib.NewField(0, 10)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmasklib.RangeGen(0.001, 0.005)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.RangeGen(0.02, 0.03)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.TRI)
	m := gmasklib.MaskGen(g, 200, 800)
	q := gmasklib.QuantGen(m, gmasklib.BpfGen([]float64{200, 50}, nil), 0.95,
		gmasklib.BpfGen([]float64{0, 150}, nil))
	p.Num, p.Gen = 4, q
	f.AddParam(p)

	p.Num, p.Gen = 5, gmasklib.RangeGen(0.5, 0.6)
	f.AddParam(p)

	g = gmasklib.RndGen(gmasklib.UNI)
	m = gmasklib.MaskGen(g, gmasklib.BpfGen([]float64{0, 0, 5, 0.8, 10, 0}, nil),
		gmasklib.BpfGen([]float64{0, 0.2, 5, 1, 10, 0.2}, nil))
	p.Num, p.Gen = 6, m
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
	ready <- true
}

func events2(cs csnd.CSOUND, ready chan bool) {
	f := gmasklib.NewField(4, 6)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmasklib.RangeGen(0.001, 0.005)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.RangeGen(0.04, 0.08)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.TRI)
	m := gmasklib.MaskGen(g, gmasklib.BpfGen([]float64{2000, 1000}, nil),
		gmasklib.BpfGen([]float64{2010, 3000}, nil))
	p.Num, p.Gen = 4, m
	f.AddParam(p)

	p.Num, p.Gen = 5, gmasklib.RangeGen(0.3, 0.4)
	f.AddParam(p)

	p.Num, p.Gen = 6, gmasklib.RangeGen(0, 0.2)
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
	ready <- true
}

func events3(cs csnd.CSOUND, ready chan bool) {
	f := gmasklib.NewField(6.5, 9.5)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.UNI)
	m := gmasklib.MaskGen(g, gmasklib.BpfGen([]float64{0.001, 0.1}, nil),
		gmasklib.BpfGen([]float64{0.005, 0.2}, nil), 1)
	p.Num, p.Gen = 2, m
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.RangeGen(0.04, 0.08)
	f.AddParam(p)

	g = gmasklib.RndGen(gmasklib.TRI)
	m = gmasklib.MaskGen(g, gmasklib.BpfGen([]float64{4000, 2000}, nil),
		gmasklib.BpfGen([]float64{8000, 3000}, nil), 1)
	p.Num, p.Gen = 4, m
	f.AddParam(p)

	g = gmasklib.RndGen(gmasklib.UNI)
	m = gmasklib.MaskGen(g, gmasklib.BpfGen([]float64{0.3, 0.5}, nil),
		gmasklib.BpfGen([]float64{0.4, 0.8}, nil))
	p.Num, p.Gen = 5, m
	f.AddParam(p)

	p.Num, p.Gen = 6, gmasklib.RangeGen(0.8, 1)
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
	go events3(cs, ready)
	for i := 1; i <= 3; i++ {
		<-ready
	}
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
