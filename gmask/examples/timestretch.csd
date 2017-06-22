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

kenv	oscil		20000,1/p3,2
aindx	line		p4,p3,p3+p4
asig	tablei	aindx*sr,1

	out		asig*kenv

endin	
</CsInstruments>

<CsScore bin="gmask">
{
f1 0 131072 1 "schwermt.aif" 0 4 1  ;43520
f2 0 8193 8 0 4096 1 4096 0
}


;original sound length = 2.2 sec
;after timestretching = 11 sec
f 0 11  

p1 const 1

p2 const .01 	;constant grain interonset 10 ms

p3 const .02 	;constant grain duration 20 ms

p4 const .002	;constant walk through the soundfile table
accum on		;with 1/5 of interonset = timestretch factor 5
prec 3
</CsScore>
</CsoundSynthesizer>

