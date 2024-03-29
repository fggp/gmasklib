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
;p4 sound file pointer

kenv  = oscil(0.61, 1/p3, 2)
aindx = line:a(p4, p3, p3+p4)
asig  = tablei:a(aindx*sr, 1)

     out asig*kenv

  endin
`

var sco string = `
f1 0 131072 1 "../samples/schwermt.aif" 0 4 1  ;43520
f2 0 8193 8 0 4096 1 4096 0

f 0 11
`

func events() string {
	f := gmasklib.NewField(0, 11)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmasklib.ConstGen(0.01)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.ConstGen(0.02)
	f.AddParam(p)

	g := gmasklib.BpfGen([]float64{0.01, 0.0004},
		gmasklib.NewInterpolation(4, gmasklib.IPLNUM))
	p.Num, p.Gen, p.Prec = 4, gmasklib.AccumGen(g, gmasklib.ON), 4
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
