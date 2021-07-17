package gmasklib

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func Test(t *testing.T) {
	rand.Seed(int64(time.Now().Nanosecond()))
	display("Constant", ConstGen(2705.1956), 10, 4, "\n")

	list := []float64{1, 2, 3, 4}
	display("Item (cycle)", ItemGen(CYCLE, list), 14, 0, " ")
	display("Item (swing)", ItemGen(SWING, list), 14, 0, " ")
	display("Item (heap)", ItemGen(HEAP, list), 14, 0, " ")
	display("Item (random)", ItemGen(RANDOM, list), 14, 0, " ")

	points := []float64{2, 1, 4, 0.5, 5, 3.0, 6, 0.27}
	fmt.Println(points)
	ipl := Interpolation{off: true}
	displayt("Bpf", BpfGen(points, &ipl), 0.0, 8.0, 0.3, 2, "\n")

	display("Range", RangeGen(-27.19, +5.56), 20, 2, " ")
	display("Rnd (uni)", RndGen(UNI), 20, 2, " ")
	display("Rnd (lin)", RndGen(LIN), 20, 2, " ")
	display("Rnd (rlin)", RndGen(RLIN), 20, 2, " ")
	display("Rnd (tri)", RndGen(TRI), 20, 2, " ")
	display("Rnd (exp[2])", RndGen(EXP, 2), 20, 2, " ")
	display("Rnd (exp[gen()])", RndGen(EXP, ItemGen(HEAP, list)), 20, 2, " ")
	displayt("Rnd (exp[gen(t)])", RndGen(EXP, BpfGen(points, &ipl)),
		0.0, 8.0, 0.3, 2, "\n")
}

func display(msg string, gen Generator, n, dec int, sep string) {
	fmt.Printf("\n%s\n", msg)
	s := fmt.Sprintf("%%.%df%%s", dec)
	for i := 0; i < n; i++ {
		fmt.Printf(s, gen(), sep)
	}
	fmt.Println()
}

func displayt(msg string, gen Generator, t1, t2, dt float64, dec int, sep string) {
	fmt.Printf("\n%s\n", msg)
	s := fmt.Sprintf("%%.%df -> %%2.%df%%s", dec, dec)
	for t := t1; t <= t2; t += dt {
		fmt.Printf(s, t, gen(t), sep)
	}
}
