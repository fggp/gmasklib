Go examples
========

The go files in this directory are a transcription of Andre Bartetzki's original
examples into go programs. See in the main func of the programs how concurrency
managing is easy with go. One or more
go routines generate events, then another goroutine performs the orchestra. No
threads, no locks, no mutexes, just a boolean channel to keep the main program
alive while there is something to play!

One can run directly a program (`go run axa1.go` for example) or one can compile
the program (`go build axa1.go`) and then run it (`./axa1`).
