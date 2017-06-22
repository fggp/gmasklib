Csd examples
========

The csd files in this directory are a transcription of Andre Bartetzki's original
examples. The CsScore section in each csd contains the original input to the Cmask
program. The CsSCore tag has an attribute bin="gmask". This means that when the
csd is executed by Csound, Csound calls the standalone gmask program, giving to
it the CsScore text as input. gmask returns the generated score which is then read
by Csound and executed. So a simple command like `csound axa1.csd` launches
all the process.  

Note: some csd files use samples. These samples have to be reachable: you shoudl copy
them from the [https://github.com/fggp/gmask/samples](https://github.com/fggp/gmask/tree/master/samples) into your SSDIR.
