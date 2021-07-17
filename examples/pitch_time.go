package main

import (
	"github.com/fggp/gmasklib"
	"github.com/fggp/go-csnd"
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
f1 0 131072 1 "../samples/schwermt.aif" 0 4 1  ;95962
f4 0 8193 8 0 4096 1 4096 0

f 0 11`

func events(cs csnd.CSOUND) {
	f := gmasklib.NewField(0, 11)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmasklib.ConstGen(0.01)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.ConstGen(0.02)
	f.AddParam(p)

	g := gmasklib.ConstGen(0.002)
	a := gmasklib.AccumGen(g, gmasklib.ON)
	p.Num, p.Gen, p.Prec = 4, a, 3
	f.AddParam(p)

	p.Num, p.Gen, p.Prec = 5, gmasklib.ConstGen(1.5), 5
	f.AddParam(p)

	f.EvalToScoreEvents(cs, true, 0)
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
	events(cs)
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
