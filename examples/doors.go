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

garev   init   0

  instr  1
;p4 transposition (1=normal)
;p5 table number (1...6)
;p6 pan (0...1)
;p7 dry/wet (0...1)

ipanl = table(1-p6, 10, 1)
ipanr = table(p6, 10, 1)

k1 = expon(0.5, p3, 0.01)
a1   loscil k1, p4, p5, 1, 0, 0, 2
a1 = linen:a(a1, 0, p3, 0.05)

garev = garev + a1*p7  
a2    = a1*ipanr
a1    = a1*ipanl
     outs  a1*(1-p7*p7), a2*(1-p7*p7)
  endin

  instr 99
krev = expseg(0.03, p3-5, 4, 5, 4) 
kral = linseg(0, p3*0.3, 1.1, p3*0.3, 0, p3*0.4, 0)
kral = kral*kral  

a1  = alpass(garev, kral, 0.05)
a2  = alpass(garev, kral, 0.06)
a1  = a1 * kral
a2  = a2 * kral
a1r = reverb2(garev+a1, krev, 0.3)
a2r = reverb2(garev+a2, krev*1.2, 0.33)
     outs  a1r+a1/2, a2r+a2/2

garev  =   0
  endin
`

var sco string = `
f1 0 0 -1 "../samples/door1.aif" 0 4 1
f2 0 0 -1 "../samples/door2.aif" 0 4 1
f3 0 0 -1 "../samples/door3.aif" 0 4 1
f4 0 0 -1 "../samples/door4.aif" 0 4 1
f5 0 0 -1 "../samples/door5.aif" 0 4 1
f6 0 0 -1 "../samples/door6.aif" 0 4 1

f10 0 8192 9 .25 1 0
i99 0 27
`

func events() string {
	f := gmasklib.NewField(0, 20)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.BETA, 0.05, 0.1)
	m := gmasklib.MaskGen(g, gmasklib.BpfGen([]float64{12, 0.01, 18, 0.2}, nil),
		gmasklib.BpfGen([]float64{12, 0.1, 18, 1}, nil))
	p.Num, p.Gen = 2, m
	f.AddParam(p)

	i := gmasklib.NewInterpolation(0.4, gmasklib.IPLNUM)
	g = gmasklib.BpfGen([]float64{0.3, 1.2}, i)
	p.Num, p.Gen = 3, g
	f.AddParam(p)

	g = gmasklib.RndGen(gmasklib.UNI)
	m = gmasklib.MaskGen(g, gmasklib.BpfGen([]float64{3, 0.8}, i),
		gmasklib.BpfGen([]float64{5, 1.2}, i))
	p.Num, p.Gen = 4, m
	f.AddParam(p)

	p.Num, p.Gen, p.Prec = 5, gmasklib.RangeGen(1, 6), 0
	f.AddParam(p)

	p.Num, p.Gen, p.Prec = 6, gmasklib.RangeGen(0, 1), 5
	f.AddParam(p)

	i = gmasklib.NewInterpolation(1, gmasklib.IPLNUM)
	g = gmasklib.BpfGen([]float64{2, 0, 18, 0.5}, i)
	p.Num, p.Gen = 7, g
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
