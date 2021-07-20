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

instr 1

ipanl	table	1-p5 ,4,1
ipanr	table	p5 ,4,1

andx	line	p4,p3,p4+p3*p6
asig	tablei	andx*sr,1
kamp	oscil	8000,1/p3,2
		outs	asig*kamp*ipanl, asig*kamp*ipanr  

endin`

var sco string = `
f1 0 262144 1 "../samples/whisp.aif" 0 4 1
;= 5.94 sec
f2 0 8192 19 1 1 270 1
f4 0 8192 9 .25 1 0			; pan function

f 0 60
`

func events() string {
	f := gmasklib.NewField(0, 60)
	p := gmasklib.NewParam(1, gmasklib.ConstGen(1), 5)
	f.AddParam(p)

	g := gmasklib.RndGen(gmasklib.UNI)
	b1 := gmasklib.BpfGen([]float64{0, 0.0005, 37, 0.007, 60, 0.003}, nil)
	b2 := gmasklib.BpfGen([]float64{0, 0.003, 37, 0.15, 60, 0.005}, nil)
	m := gmasklib.MaskGen(g, b1, b2)
	p.Num, p.Gen = 2, m
	f.AddParam(p)

	b1 = gmasklib.BpfGen([]float64{0.3, 0.02}, nil)
	b2 = gmasklib.BpfGen([]float64{0.7, 0.04}, nil)
	m = gmasklib.MaskGen(g, b1, b2)
	p.Num, p.Gen = 3, m
	f.AddParam(p)

	p.Num, p.Gen = 4, gmasklib.BpfGen([]float64{0, 5.9}, nil)
	f.AddParam(p)

	p.Num, p.Gen = 5, gmasklib.RangeGen(0, 1)
	f.AddParam(p)

	b1 = gmasklib.BpfGen([]float64{0, 0.3, 25, 1, 40, 0.7}, nil)
	b2 = gmasklib.BpfGen([]float64{0, 2, 4, 1, 25, 1.2}, nil)
	m = gmasklib.MaskGen(g, b1, b2)
	b1 = gmasklib.BpfGen([]float64{0, 0, 25, 0.9, 30, 0, 45, 0.9, 55, 0}, nil)
	b2 = gmasklib.BpfGen([]float64{40, 0, 45, 1.5, 55, 0}, nil)
	q := gmasklib.QuantGen(m, 0.3, b1, b2)
	p.Num, p.Gen = 6, q
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
