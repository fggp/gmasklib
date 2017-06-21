Gmask standalone program
========

The lexer (gmlex.go) uses the Go text/scanner package which is enough for the
tokens used in the Cmask grammar
(see http://www2.ak.tu-berlin.de/~abartetzki/CMaskMan/CMask-Reference.htm).

The parser is defined in the gmask.y file. It is generated using the command
`go tool yacc -o gmask.go gmask.y`.

Then one can build and install the program with the command `go install`.

Once the program is built, you can run the examples in the 'examples' directory with the command 'csound example-name.csd', provided that the gmask program is in your path.
