Adaptation of Cmask to the Go programming language
========

This package is an adaptation of Andre Bartetzki's Cmask software for the Go
programming language. It provides a library that can be imported in any Go
program and a standalone program which uses the library. Both the library and
the program are called gmask. This is not a problem, because Go uses different
places to store packages (GOPATH/pkg) and compiled programs (GOPATH/bin).

You can install this package with `go get`:

  `go get github.com/fggp/gmask`

Or you can download a zip archive of the project using the 'Download ZIP' button on the right.
You'll get a zip file named 'gmask-master.zip'. Decompressing it you'll get a directory named 'gmask-master'.
Rename this directory to 'gmask' and move it to '$GOPATH/src/github/fggp'. Enter into
the '$GOPATH/src/github/fggp/gmask' directory and install the package with `go install`.

Once the package is installed, you'll have to compile the gmask standalone program.
Enter into the '$GOPATH/src/github/fggp/gmask/gmask' directory. The parser gmask.go
has already been generated with the go yacc tool. Build and install the program
with `go install`.

The original Cmask software was written by Andre Bartetzki. Andre has stopped Cmask development but the software is still present on his site: http://www.bartetzki.de/en/software.html.

Cmask was published under GPL. Thanks to Andre who kindly allowed me to publish gmask under LGPL:  

"Dear François,

thanks for bringing Cmask to a new life!  
Yes, you may publish Gmask under LGPL.  
best

Andre"

The gmask program reflects exactly Cmask features. It has a parser that recognizes the grammar written by Andre: http://www2.ak.tu-berlin.de/~abartetzki/CMaskMan/CMask-Reference.htm

When the program is called on a parameter file respecting Cmask language, it will output a
Csound sco file on standard out. One can also write the attribute bin="gmask" in a CsScore tag of a csd file to get the score generated on the fly while playing the csd file with Csound. See the examples directory in the gmask/gmask directory.

The gmask library can be imported in any go program. See the gmask/docs directory for the library documentation. The library is more permissive than the gmask program. For example, in the gmask program some generators can take args that are a value or a breakpoint function. In the library, those generators can accept for the same args a value or any generator. This means that you could for example create a tendency mask with a rnd generator as low boundarie and an osc-mask-quantizer daisy chain as high boundarie: in the gmask library, generators have type Generator and modifiers as well.
