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
;p4 speed factor (=transposition)

kenv	oscil		30000,1/p3,2
aindx	line		p2,p3,p2+p3*p4
asig	tablei	aindx*sr,1

	out		asig*kenv

endin	
</CsInstruments>

<CsScore bin="gmask">
{
f1 0 131072 1 "schwermt.aif" 0 4 1  ;43520
f2 0 8193 8 0 4096 1 4096 0
}

; for use with pitchshift.orc !

f 0 2.2  

p1 const 1

p2 const .02 	;constant grain interonset 20 ms

p3 const .04 	;constant grain duration 40 ms

p4 range .75 1.5 	;random intervall
prec 2
</CsScore>
</CsoundSynthesizer>

