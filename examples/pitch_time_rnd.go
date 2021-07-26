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
;p5 speed factor (=transposition)

kenv  = oscil(0.61, 1/p3, 4)
aindx = line:a(p4, p3, p4 + p3*p5)
asig  = tablei:a(aindx*sr, 1)
     out asig*kenv
endin  
`

var sco string = `
f1 0 131072 1 "../samples/schwermt.aif" 0 4 1  
f4 0 8193 8 0 4096 1 4096 0

f 0 22
`

func events1(ret chan string) {
	f := gmasklib.NewField(0, 20)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	p.Num, p.Gen = 2, gmasklib.ConstGen(0.02)
	f.AddParam(p)

	p.Num, p.Gen = 3, gmasklib.ConstGen(0.04)
	f.AddParam(p)

	g := gmasklib.ConstGen(0.002)
	a := gmasklib.AccumGen(g, gmasklib.ON)
	p.Num, p.Gen, p.Prec = 4, a, 3
	f.AddParam(p)

	g = gmasklib.RangeGen(0.5, 2.5)
	b := gmasklib.BpfGen([]float64{0, 0, 5, 1, 17, 1, 22, 0}, nil)
	q := gmasklib.QuantGen(g, 0.5, b)
	p.Num, p.Gen, p.Prec = 5, q, 5
	f.AddParam(p)

	var buf bytes.Buffer
	f.EvalToScore(&buf, 1)
	ret <- buf.String()
}

func events2(ret chan string) {
	f := gmasklib.NewField(5.5, 16.5)
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

	p.Num, p.Gen, p.Prec = 5, gmasklib.ConstGen(3.0), 5
	f.AddParam(p)

	var buf bytes.Buffer
	f.EvalToScore(&buf, 1)
	ret <- buf.String()
}

func perform(cs csnd.CSOUND, done chan bool) {
	cs.Perform()
	done <- true
}

func main() {
	cs := csnd.Create(nil)
	cs.SetOption("-odac")
	cs.CompileOrc(orc)
	s := make(chan string, 2)
	go events1(s)
	go events2(s)
	for i := 1; i <= 2; i++ {
		sco += <-s
	}
	cs.ReadScore(sco)
	cs.Start()
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
