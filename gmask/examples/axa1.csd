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
;p4 grain pointer (in seconds)
;p5 pan (0...1)


ipanl	table	1-p5 ,4,1
ipanr	table	p5 ,4,1

andx	line	p4,p3,p4+p3
asig	tablei	andx*sr,1

k1	oscil	30000,1/p3,2	
asig	=	asig*k1
	
	outs	asig*ipanl, asig*ipanr
	
          endin
</CsInstruments>

<CsScore bin="gmask">
{
f1 0 65536 1 "axaxaxas.aif" 0 4 1
;= 1.4861 sec

f2 0 8193 19 1 1 270 1
f4 0 8192 9 .25 1 0
}

f 0 5

p1 const 1

p2 const .02
;p2 rnd uni mask  .005 .1 map 1

p3 const .04
;p3 range .04 .2	

p4 seg [0 1.44 ipl 0]
;p4 seg [0 1.44 ipl -2]
;p4 range 0 1.44
;p4 rnd uni mask .002 .05 map 1
;accum wrap 0 1.4

p5 const .5
</CsScore>
</CsoundSynthesizer>

