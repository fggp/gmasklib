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
;p4 grain pointer (in seconds)
;p5 pan (0...1)

ipanl = table(1-p5 ,4, 1)
ipanr = table(p5 , 4, 1)

andx  = line:a(p4, p3, p4+p3)
asig  = tablei:a(andx*sr, 1)

k1    = oscil(0.91, 1/p3, 2)
asig  = asig*k1
  
        outs asig*ipanl, asig*ipanr
  endin
`

var sco string = `
f1 0 65536 1 "../samples/axaxaxas.aif" 0 4 1
;= 1.4861 sec

f2 0 8193 19 1 1 270 1
f4 0 8192 9 .25 1 0

f 0 5
`

func events() string {
	f := gmasklib.NewField(0, 5)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmasklib.ConstGen(0.02)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.ConstGen(0.04)
	f.AddParam(p)

	g := gmasklib.BpfGen([]float64{0, 1.44}, gmasklib.NewInterpolation(0, false, true))
	p.Num, p.Gen = 4, g
	f.AddParam(p)

	//p.Num, p.Gen = 5, gmasklib.ConstGen(0.5)
	p.Num, p.Gen = 5, gmasklib.BpfGen([]float64{0, 0, 2.5, 1, 5, 0}, 
		gmasklib.NewInterpolation(0, false, true))
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
