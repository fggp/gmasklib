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

  instr 1    ;mallet ?

;p2 onset
;p3 duration
;p4 pitch (0-4)
;p5 octav (7-9)

kenv  = oscil(1, 1/p3, 2)
kindx = pow(kenv, 6, 0.5)
iton  = table(p4, 5)
a1    = foscil(kenv*0.25, cpspch(p5+iton), 1, 4, kindx, 1)
     outs a1*(1-p4/4), a1*p4/4
  endin  

  instr 2    ;metal plate

;p2 onset
;p3 duration
;p4 pitch (0/1)

kindx = expon(1, p3, 0.001)
a1    = rand:a(100)
a2    = oscil:a(0.31*kindx, 3000 + 1500*p4 + a1*(1+kindx), 1)
     outs a2*p4, a2*(1-p4)  
  endin  

  instr 3

;p2 onset
;p3 duration
;p4 pitch (0-3)

kenv  = oscil(1, 1/p3, 3)
kindx = oscil(2, 1/p3, 4)
a1    = foscil(kenv*0.34, 100+p4*20, 1, 1.4, kindx, 1)
     outs a1,a1
  endin  
`

var sco string = `
f1 0 8193 10 1
f2 0 8193 5 1 8193 .003
f3 0 8193 8 .8 1000 1 2192  .3 5000 0
f4 0 8193 5 1 1193 0.02 7000 .01
f5 0 8 -2 0 .02 .04 .07 .09 

f 0 20
`

func events1(ret chan string) {
	f := gmasklib.NewField(0, 20)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.EXP, 2)
	m := gmasklib.MaskGen(g, 0.1, 0.5, 1)
	q := gmasklib.QuantGen(m, 0.1, 0.96)
	p.Num, p.Gen, p.Prec = 2, q, 4
	f.AddParam(p)

	g = gmasklib.RangeGen(0.4, 0.5)
	p.Num, p.Gen, p.Prec = 3, g, 2
	f.AddParam(p)

	g = gmasklib.RangeGen(0, 4)
	p.Num, p.Gen, p.Prec = 4, g, 5
	f.AddParam(p)

	g = gmasklib.RangeGen(7, 9)
	p.Num, p.Gen = 5, g
	f.AddParam(p)

	var buf bytes.Buffer
	f.EvalToScore(&buf, 1)
	ret <- buf.String()
}

func events2(ret chan string) {
	f := gmasklib.NewField(0, 20)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(2), 5)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.RLIN)
	m := gmasklib.MaskGen(g, 0.3, 1, 1)
	q := gmasklib.QuantGen(m, 0.3, 0.96)
	p.Num, p.Gen, p.Prec = 2, q, 2
	f.AddParam(p)

	g = gmasklib.RangeGen(0.4, 0.5)
	p.Num, p.Gen = 3, g
	f.AddParam(p)

	g = gmasklib.RangeGen(0, 1)
	p.Num, p.Gen, p.Prec = 4, g, 5
	f.AddParam(p)

	var buf bytes.Buffer
	f.EvalToScore(&buf, 2)
	ret <- buf.String()
}

func events3(ret chan string) {
	f := gmasklib.NewField(0, 20)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(3), 5)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.BETA, 0.2, 0.5)
	m := gmasklib.MaskGen(g, 0.1, 1, 1)
	q := gmasklib.QuantGen(m, 0.2, 0.9)
	p.Num, p.Gen, p.Prec = 2, q, 2
	f.AddParam(p)

	g = gmasklib.RangeGen(0.8, 1.5)
	p.Num, p.Gen = 3, g
	f.AddParam(p)

	g = gmasklib.RangeGen(0, 3)
	p.Num, p.Gen, p.Prec = 4, g, 5
	f.AddParam(p)

	var buf bytes.Buffer
	f.EvalToScore(&buf, 3)
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
	s := make(chan string, 3)
	go events1(s)
	go events2(s)
	go events3(s)
	for i := 1; i <= 3; i++ {
		sco += <-s
	}
	cs.ReadScore(sco)
	cs.Start()
	done := make(chan bool)
	go perform(cs, done)
	<-done
	cs.Stop()
}
