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
;p4 speed factor (=transposition)

kenv	oscil		30000,1/p3,2
aindx	line		p2,p3,p2+p3*p4
asig	tablei	aindx*sr,1

	out		asig*kenv

endin`

var sco string = `
f1 0 131072 1 "../samples/schwermt.aif" 0 4 1  ;43520
f2 0 8193 8 0 4096 1 4096 0

f 0 2.2`

func events(cs csnd6.CSOUND) {
	f := gmask.NewField(0, 2.2)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmask.ConstGen(0.02)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmask.ConstGen(0.04)
	f.AddParam(p)

	g := gmask.RangeGen(0.5, 2)
	q := gmask.QuantGen(g, 0.5, 1)
	p.Num, p.Gen, p.Prec = 4, q, 3
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
