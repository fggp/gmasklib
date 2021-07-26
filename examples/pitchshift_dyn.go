package main

import (
	"bytes"
	"github.com/fggp/gmasklib"
	"github.com/fggp/go-csnd"
)

var orc string = `
sr     = 44100
ksmps  = 10
nchnls = 1
0dbfs  = 1.0

  instr 1

;p2 onset
;p3 duration
;p4 speed factor (=transposition)

kenv  = oscil(0.92, 1/p3, 2)
aindx = line:a(p2, p3, p2 + p3*p4)
asig  = tablei:a(aindx*sr, 1)
     out asig*kenv
  endin  
`

var sco string = `
f1 0 131072 1 "../samples/schwermt.aif" 0 4 1  
f2 0 8193 8 0 4096 1 4096 0

f 0 2.2
`

func events() string {
	f := gmasklib.NewField(0, 2.2)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmasklib.ConstGen(0.01)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.ConstGen(0.02)
	f.AddParam(p)

	b := gmasklib.BpfGen([]float64{1, 2.2},
		gmasklib.NewInterpolation(0.3, gmasklib.IPLNUM))
	p.Num, p.Gen = 4, b
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
