<CsoundSynthesizer>
; adapted from Andre Bartetzki's original cmask example
; see http://www.bartetzki.de/en/index.html
<CsOptions>
  -d -o dac
</CsOptions>

<CsInstruments>
sr     = 44100
ksmps  = 10
nchnls = 1

instr 1

;p2 onset
;p3 duration
;p4 sound file pointer
;p5 speed factor (=transposition)

kenv	oscil		20000,1/p3,4
aindx	line		p4,p3,p4+p3*p5
asig	tablei	aindx*sr,1

	out		asig*kenv
endin	
</CsInstruments>

<CsScore bin="gmask">
{
f1 0 131072 1 "schwermt.aif" 0 4 1  
f4 0 8193 8 0 4096 1 4096 0
}

;for use with pitch+time.orc !

f 0 22  

p1 const 1

p2 const 0.02 	;constant grain interonset 20 ms

p3 const 0.04 	;constant grain duration 40 ms

p4 const .002 	;1/10 tempo
accum on
prec 3

p5 range .5 2.5
quant .5 (0 0 5 1 17 1 22 0)


f 5.5 16.5 		;a second field   

p1 const 1

p2 const 0.01 	;constant grain interonset 10 ms

p3 const 0.02 	;constant grain duration 20 ms

p4 const .002
accum on 
prec 3

p5 const 3.0 
</CsScore>
</CsoundSynthesizer>

