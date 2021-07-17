Adaptation of Cmask to the Go programming language
========

This package is an adaptation of Andre Bartetzki's Cmask software for the Go
programming language. It provides a library that can be imported in any Go
program. The library is called gmasklib.

You can install this package with `go get`:

  `go get github.com/fggp/gmask`

The original Cmask software was written by Andre Bartetzki. Andre has stopped Cmask development but the software is still present on his site: http://www.bartetzki.de/en/software.html.

Cmask was published under GPL. Thanks to Andre who kindly allowed me to publish gmasklib under LGPL:  

"Dear François,

thanks for bringing Cmask to a new life!  
Yes, you may publish Gmask under LGPL.  
best

Andre"

The gmask library can be imported in any go program. See the gmasklib/docs directory for the library documentation. The library is more permissive than the gmask program. For example, in the gmask program some generators can take args that are a value or a breakpoint function. In the library, those generators can accept for the same args a value or any generator. This means that you could for example create a tendency mask with a rnd generator as low boundarie and an osc-mask-quantizer daisy chain as high boundarie: in the gmask library, generators have type Generator and modifiers as well.
