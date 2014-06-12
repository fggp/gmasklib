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

func events1(cs csnd6.CSOUND) {
	f := gmask.NewField(0, 30)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	g := gmask.RndGen(gmask.UNI)
	m := gmask.MaskGen(g, gmask.BpfGen([]float64{0.01, 0.002}, nil),
		gmask.BpfGen([]float64{0.1, 0.01}, nil))
	p.Num, p.Gen = 2, m
	f.AddParam(p)

	p.Num, p.Gen = 3, gmask.RangeGen(0.5, 1)
	f.AddParam(p)

	b1 := gmask.BpfGen([]float64{860, 80}, gmask.NewInterpolation(-1.2, false, false))
	b2 := gmask.BpfGen([]float64{940, 2000}, gmask.NewInterpolation(1, false, false))
	m = gmask.MaskGen(g, b1, b2, 1)
	q := gmask.QuantGen(m, 100, 0.9, 0)
	p.Num, p.Gen = 4, q
	f.AddParam(p)

	b1 = gmask.BpfGen([]float64{0.4, 0}, nil)
	b2 = gmask.BpfGen([]float64{0.6, 1}, nil)
	p.Num, p.Gen = 5, gmask.MaskGen(g, b1, b2)
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
}

func events2(cs csnd6.CSOUND) {
	f := gmask.NewField(31, 33)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	g := gmask.BpfGen([]float64{0.08, 0.8}, gmask.NewInterpolation(2, false, false))
	p.Num, p.Gen = 2, g
	f.AddParam(p)

	g = gmask.BpfGen([]float64{0.1, 2}, nil)
	p.Num, p.Gen = 3, g
	f.AddParam(p)

	p.Num, p.Gen = 4, gmask.RangeGen(300, 400)
	f.AddParam(p)

	g = gmask.BpfGen([]float64{0, 1}, nil)
	p.Num, p.Gen = 5, g
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
	// Here we have a subtle issue. As the first field will generate more than
	// 3300 notes, there is a risk that the system can not compute all the notes
	// in real-time. So calling events1 as a normal func will block the system
	// until the field is computed. Then we can launch go routines as usual...
	events1(cs)
	go events2(cs)
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
