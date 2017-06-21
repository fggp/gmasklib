package main

import (
	"github.com/fggp/gmask"
	"github.com/fggp/go-csnd6"
)

var orc string = `
sr     = 44100
ksmps  = 10
nchnls = 2

          instr 1
;p4 grain pointer (in seconds)
;p5 pan (0...1)


ipanl	table	1-p5 ,4,1
ipanr	table	p5 ,4,1

andx	line	p4,p3,p4+p3
asig	tablei	andx*sr,1

k1	oscil	30000,1/p3,2	
asig	=	asig*k1
	
	outs	asig*ipanl, asig*ipanr
	
          endin`

var sco string = `
f1 0 65536 1 "../samples/axaxaxas.aif" 0 4 1
;= 1.4861 sec

f2 0 8193 19 1 1 270 1
f4 0 8192 9 .25 1 0

f 0 5`

func events(cs csnd6.CSOUND) {
	f := gmask.NewField(0, 5)
	p := gmask.NewParam(1, gmask.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmask.ConstGen(0.02)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmask.ConstGen(0.04)
	f.AddParam(p)

	g := gmask.BpfGen([]float64{0, 1.44}, gmask.NewInterpolation(0, false, false))
	p.Num, p.Gen = 4, g
	f.AddParam(p)

	p.Num, p.Gen = 5, gmask.ConstGen(0.5)
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
	events(cs)
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
