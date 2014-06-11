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

k1	oscil	8000*p5,1/p3,1
a1	oscil	k1,p4,2
	outs	a1*(1-p6),a1*p6

endin`

var sco string = `
f1 0 8193 8 0 4096 1 4096 0
f2 0 8193 10 1 .5 .3 .2 .1

f 0 10`

func events1(cs csnd6.CSOUND) {
	f := gmask.NewField(0, 10)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmask.ConstGen(0.003) //RangeGen(0.001, 0.005)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmask.ConstGen(0.025) //RangeGen(0.02, 0.03)
	f.AddParam(p)

	g := gmask.RndGen(gmask.TRI)
	m := gmask.MaskGen(g, 200, 800)
	q := gmask.ConstGen(400) //QuantGen(m, gmask.BpfGen([]float64{200, 50}, nil), 0.95,
		//gmask.BpfGen([]float64{0, 150}, nil))
	p.Num, p.Gen = 4, q
	f.AddParam(p)

	p.Num, p.Gen = 5, gmask.ConstGen(0.55) //RangeGen(0.5, 0.6)
	f.AddParam(p)

	g = gmask.RndGen(gmask.UNI)
	m = gmask.ConstGen(0.9) //MaskGen(g, gmask.BpfGen([]float64{0, 0, 5, 0.8, 10, 0}, nil),
		//gmask.BpfGen([]float64{0, 0.2, 5, 1, 10, 0.2}, nil))
	p.Num, p.Gen = 6, m
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
}

func events2(cs csnd6.CSOUND) {
	f := gmask.NewField(4, 6)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmask.RangeGen(0.001, 0.005)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmask.RangeGen(0.04, 0.08)
	f.AddParam(p)

	g := gmask.RndGen(gmask.TRI)
	m := gmask.MaskGen(g, gmask.BpfGen([]float64{2000, 1000}, nil),
		gmask.BpfGen([]float64{2010, 3000}, nil))
	p.Num, p.Gen = 4, m
	f.AddParam(p)

	p.Num, p.Gen = 5, gmask.RangeGen(0.3, 0.4)
	f.AddParam(p)

	p.Num, p.Gen = 6, gmask.RangeGen(0, 0.2)
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
}

func events3(cs csnd6.CSOUND) {
	f := gmask.NewField(6.5, 9.5)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	g := gmask.RndGen(gmask.UNI)
	m := gmask.MaskGen(g, gmask.BpfGen([]float64{0.001, 0.1}, nil),
		gmask.BpfGen([]float64{0.005, 0.2}, nil), 1)
	p.Num, p.Gen = 2, m
	f.AddParam(p)

	p.Num, p.Gen = 3, gmask.RangeGen(0.04, 0.08)
	f.AddParam(p)

	g = gmask.RndGen(gmask.TRI)
	m = gmask.MaskGen(g, gmask.BpfGen([]float64{4000, 2000}, nil),
		gmask.BpfGen([]float64{8000, 3000}, nil), 1)
	p.Num, p.Gen = 4, m
	f.AddParam(p)

	g = gmask.RndGen(gmask.UNI)
	m = gmask.MaskGen(g, gmask.BpfGen([]float64{0.3, 0.5}, nil),
		gmask.BpfGen([]float64{0.4, 0.8}, nil))
	p.Num, p.Gen = 5, m
	f.AddParam(p)

	p.Num, p.Gen = 6, gmask.RangeGen(0.8, 1)
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
	go events1(cs)
	go events2(cs)
	go events3(cs)
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
