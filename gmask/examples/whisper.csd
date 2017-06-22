<CsoundSynthesizer>
; adapted from Andre Bartetzki's original cmask example
; see http://www.bartetzki.de/en/index.html
<CsOptions>
  -d -o dac
</CsOptions>

<CsInstruments>
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

endin	
</CsInstruments>

<CsScore bin="gmask">
{
f1 0 262144 1 "whisp.aif" 0 4 1
;= 5.94 sec
f2 0 8192 19 1 1 270 1
f4 0 8192 9 .25 1 0			; pan function
}

f 0 60

p1 const 1

p2 rnd uni
mask (0 .0005 37 .007 60 .003) (0 .003 37 .15 60 .005) 

p3 rnd uni
mask [.3 .02] [.7 .04]

p4
seg [0 5.9]

p5 range 0 1

p6 rnd uni
mask (0 .3 25 1 40 .7) (0 2 4 1 25 1.2)
quant .3 (0 0 25 .9 30 0 45 .9 55 0) (40 0 45 1.5 55 0)
</CsScore>
</CsoundSynthesizer>

