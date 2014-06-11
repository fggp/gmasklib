package main

import (
	"github.com/fggp/gmask"
	csnd6 "github.com/fggp/go-csnd6"
)

var orc string = `
sr     = 44100
ksmps  = 10
nchnls = 1

instr 1

;p2 onset
;p3 duration
;p4 sound file pointer
;p5 speed factor (=transposition)

kenv	oscil		20000,1/p3,4
aindx	line		p4,p3,p4+p3*p5
asig	tablei	aindx*sr,1

	out		asig*kenv
endin`

var sco string = `
f1 0 131072 1 "../samples/schwermt.aif" 0 4 1  
f4 0 8193 8 0 4096 1 4096 0

f 0 22`

func events1(cs csnd6.CSOUND) {
	f := gmask.NewField(0, 20)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmask.ConstGen(0.02)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmask.ConstGen(0.04)
	f.AddParam(p)

	g := gmask.ConstGen(0.002)
	a := gmask.AccumGen(g, gmask.ON)
	p.Num, p.Gen = 4, a
	f.AddParam(p)

	g = gmask.RangeGen(0.5, 2.5)
	b := gmask.BpfGen([]float64{0, 0, 5, 1, 17, 1, 22, 0}, nil)
	q := gmask.QuantGen(g, 0.5, b)
	p.Num, p.Gen = 5, q
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
}

func events2(cs csnd6.CSOUND) {
	f := gmask.NewField(5.5, 16.5)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmask.ConstGen(0.01)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmask.ConstGen(0.02)
	f.AddParam(p)

	g := gmask.ConstGen(0.002)
	a := gmask.AccumGen(g, gmask.ON)
	p.Num, p.Gen = 4, a
	f.AddParam(p)

	p.Num, p.Gen = 5, gmask.ConstGen(3.0)
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
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
