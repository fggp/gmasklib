//line gmask.y:4
package main

import __yyfmt__ "fmt"

//line gmask.y:5
import (
	"bufio"
	"fmt"
	"github.com/fggp/gmask"
	"io"
	"io/ioutil"
	"log"
	"os"
)

//line gmask.y:19
type yySymType struct {
	yys int
	val float64
	n   int
	sco string
	fie *gmask.Field
	par gmask.Param
	gen gmask.Generator
	itm gmask.ItemMode
	lst *List
	ipl *gmask.Interpolation
	rmd gmask.RndMode
	omd gmask.OscMode
	amd gmask.AccumMode
	sli []interface{}
}

const SCO = 57346
const NUMBER = 57347
const F = 57348
const PARAM = 57349
const IPL = 57350
const COS = 57351
const OFF = 57352
const PREC = 57353
const ITEM = 57354
const CYCLE = 57355
const SWING = 57356
const HEAP = 57357
const RANDOM = 57358
const CONST = 57359
const SEG = 57360
const RANGE = 57361
const RND = 57362
const UNI = 57363
const LIN = 57364
const RLIN = 57365
const TRI = 57366
const EXP = 57367
const REXP = 57368
const BEXP = 57369
const GAUSS = 57370
const CAUCHY = 57371
const BETA = 57372
const WEI = 57373
const OSC = 57374
const SIN = 57375
const SQUARE = 57376
const TRIANGLE = 57377
const SAWUP = 57378
const SAWDOWN = 57379
const POWUP = 57380
const POWDOWN = 57381
const MASK = 57382
const MAP = 57383
const QUANT = 57384
const ACCUM = 57385
const ON = 57386
const LIMIT = 57387
const MIRROR = 57388
const WRAP = 57389
const INIT = 57390

var yyToknames = []string{
	"SCO",
	"NUMBER",
	"F",
	"PARAM",
	"IPL",
	"COS",
	"OFF",
	"PREC",
	"ITEM",
	"CYCLE",
	"SWING",
	"HEAP",
	"RANDOM",
	"CONST",
	"SEG",
	"RANGE",
	"RND",
	"UNI",
	"LIN",
	"RLIN",
	"TRI",
	"EXP",
	"REXP",
	"BEXP",
	"GAUSS",
	"CAUCHY",
	"BETA",
	"WEI",
	"OSC",
	"SIN",
	"SQUARE",
	"TRIANGLE",
	"SAWUP",
	"SAWDOWN",
	"POWUP",
	"POWDOWN",
	"MASK",
	"MAP",
	"QUANT",
	"ACCUM",
	"ON",
	"LIMIT",
	"MIRROR",
	"WRAP",
	"INIT",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line gmask.y:193
const SLICEINC = 10

type List struct {
	pos int
	val []float64
}

func NewList(x float64) *List {
	v := make([]float64, SLICEINC)
	v[0] = x
	l := List{pos: 1, val: v}
	return &l
}

func NewBpList(t, x float64) *List {
	v := make([]float64, SLICEINC)
	v[0], v[1] = t, x
	l := List{pos: 2, val: v}
	return &l
}

func (l *List) AddVal(x float64) {
	if l.pos == len(l.val) {
		l.val = append(l.val, make([]float64, SLICEINC)...)
	}
	l.val[l.pos] = x
	l.pos++
}

func (l *List) AddBp(t, x float64) {
	if l.pos == len(l.val) {
		l.val = append(l.val, make([]float64, SLICEINC)...)
	}
	l.val[l.pos], l.val[l.pos+1] = t, x
	l.pos += 2
}

func (l *List) GetVal() []float64 {
	return l.val[:l.pos]
}

func NewInterfaceSlice(params ...interface{}) []interface{} {
	return params
}

var w io.Writer
var fieldNum int

func main() {
	input, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	if len(os.Args) <= 2 {
		w = os.Stdout
	} else {
		fo, err := os.Create(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
		defer fo.Close()
		w = bufio.NewWriter(fo)
	}
	yyParse(NewLexer(input))
}

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 93
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 163

var yyAct = []int{

	33, 106, 120, 118, 128, 122, 116, 114, 32, 90,
	102, 100, 69, 64, 98, 96, 108, 105, 42, 44,
	94, 92, 3, 5, 85, 83, 34, 123, 35, 72,
	74, 76, 124, 7, 75, 73, 125, 126, 71, 43,
	113, 112, 84, 86, 41, 111, 34, 34, 35, 35,
	34, 34, 35, 35, 34, 34, 35, 35, 34, 34,
	35, 35, 104, 110, 34, 34, 35, 35, 34, 34,
	35, 35, 93, 95, 97, 99, 101, 103, 34, 34,
	35, 35, 34, 34, 35, 35, 109, 91, 34, 57,
	35, 65, 66, 67, 68, 27, 89, 115, 117, 119,
	121, 16, 107, 88, 87, 108, 10, 12, 13, 19,
	82, 81, 127, 56, 58, 59, 60, 61, 62, 63,
	80, 20, 78, 70, 28, 36, 29, 25, 45, 46,
	47, 48, 49, 50, 51, 52, 53, 54, 55, 37,
	38, 39, 40, 31, 21, 8, 1, 30, 26, 24,
	23, 18, 17, 79, 77, 11, 15, 14, 9, 6,
	4, 22, 2,
}
var yyPact = []int{

	18, 17, -1000, -1000, 26, 140, -1000, 89, 139, 84,
	138, -41, -23, 120, -1000, -1000, 126, 39, 34, 107,
	80, -1000, -1000, -28, -1000, 47, -36, 118, 33, 30,
	29, -1000, 117, -1000, 115, 106, 105, -1000, -1000, -1000,
	-1000, 20, 19, 99, 98, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, 91, -39, -1000, -1000, -1000, 82,
	-1000, 16, 15, 10, 9, 6, 5, 12, -1000, 97,
	81, 58, -1000, -1000, -1000, -1000, -1000, 40, 36, -1000,
	35, -1000, -1000, -1000, -1000, -1000, 2, 1, -2, -3,
	-1000, -1000, -1000, -1000, -1000, -1000, -45, 22, 27, -1000,
	8, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000, -1000,
	-1000, -1000, -1000, -1000, -1000, -1000, -1000, -48, -1000,
}
var yyPgo = []int{

	0, 162, 161, 160, 159, 158, 0, 157, 156, 155,
	154, 153, 1, 152, 151, 150, 149, 148, 147, 146,
}
var yyR1 = []int{

	0, 19, 19, 1, 3, 3, 4, 5, 5, 5,
	5, 5, 5, 5, 5, 5, 5, 5, 9, 9,
	9, 9, 10, 10, 6, 6, 11, 11, 12, 12,
	12, 12, 13, 13, 13, 13, 13, 13, 13, 13,
	13, 13, 13, 7, 7, 7, 7, 7, 7, 7,
	14, 14, 14, 14, 14, 14, 14, 14, 8, 8,
	8, 8, 8, 8, 15, 15, 15, 15, 15, 16,
	16, 16, 16, 16, 16, 16, 16, 16, 16, 16,
	16, 16, 16, 18, 18, 18, 17, 17, 17, 17,
	17, 2, 2,
}
var yyR2 = []int{

	0, 1, 2, 1, 3, 2, 3, 2, 4, 2,
	3, 1, 1, 2, 2, 3, 5, 2, 2, 2,
	2, 2, 1, 2, 4, 5, 2, 3, 0, 2,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	2, 2, 2, 1, 2, 2, 3, 3, 3, 3,
	2, 2, 2, 2, 2, 2, 2, 2, 2, 2,
	3, 3, 4, 4, 3, 3, 3, 3, 3, 2,
	2, 3, 3, 3, 3, 4, 4, 4, 4, 4,
	4, 4, 4, 2, 2, 2, 3, 3, 3, 3,
	3, 0, 2,
}
var yyChk = []int{

	-1000, -19, -1, 4, -3, 6, -4, 7, 5, -5,
	17, -9, 18, 19, -7, -8, 12, -13, -14, 20,
	32, 5, -2, -15, -16, 43, -17, 11, 40, 42,
	-18, 5, 49, -6, 49, 51, 5, 13, 14, 15,
	16, 5, -6, 5, -6, 21, 22, 23, 24, 25,
	26, 27, 28, 29, 30, 31, 33, 9, 34, 35,
	36, 37, 38, 39, 41, 44, 45, 46, 47, 48,
	5, 5, -6, 5, -6, 5, -6, -10, 5, -11,
	5, 5, 5, 5, -6, 5, -6, 5, 5, 5,
	48, 5, 5, -6, 5, -6, 5, -6, 5, -6,
	5, -6, 5, -6, 50, 5, -12, 5, 8, 5,
	5, 5, 5, 5, 5, -6, 5, -6, 5, -6,
	5, -6, 50, 5, 5, 9, 10, -12, 52,
}
var yyDef = []int{

	0, -2, 1, 3, 2, 0, 5, 0, 0, 91,
	0, 0, 0, 0, 11, 12, 0, 43, 0, 0,
	0, 4, 6, 13, 14, 0, 17, 0, 0, 0,
	0, 7, 0, 9, 0, 0, 0, 18, 19, 20,
	21, 44, 45, 58, 59, 32, 33, 34, 35, 36,
	37, 38, 39, 40, 41, 42, 50, 51, 52, 53,
	54, 55, 56, 57, 0, 15, 83, 84, 85, 0,
	92, 0, 0, 69, 70, 0, 0, 0, 22, 28,
	0, 0, 10, 46, 48, 47, 49, 60, 61, 68,
	0, 90, 64, 66, 65, 67, 71, 72, 73, 74,
	86, 88, 87, 89, 8, 23, 0, 0, 0, 26,
	28, 62, 63, 16, 75, 76, 77, 78, 79, 80,
	81, 82, 24, 27, 29, 30, 31, 0, 25,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	49, 50, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 51, 3, 52,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 20, 21,
	22, 23, 24, 25, 26, 27, 28, 29, 30, 31,
	32, 33, 34, 35, 36, 37, 38, 39, 40, 41,
	42, 43, 44, 45, 46, 47, 48,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line gmask.y:63
		{
			fmt.Fprintf(w, "%s\n", yyS[yypt-0].sco)
		}
	case 2:
		//line gmask.y:64
		{
			fieldNum++
			yyS[yypt-0].fie.EvalToScore(w, fieldNum)
		}
	case 3:
		//line gmask.y:67
		{
			yyVAL.sco = yyS[yypt-0].sco
		}
	case 4:
		//line gmask.y:70
		{
			yyVAL.fie = gmask.NewField(yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 5:
		//line gmask.y:71
		{
			yyS[yypt-1].fie.AddParam(yyS[yypt-0].par)
		}
	case 6:
		//line gmask.y:74
		{
			yyVAL.par = gmask.NewParam(yyS[yypt-2].n, yyS[yypt-1].gen, yyS[yypt-0].n)
		}
	case 7:
		//line gmask.y:77
		{
			yyVAL.gen = gmask.ConstGen(yyS[yypt-0].val)
		}
	case 8:
		//line gmask.y:78
		{
			yyVAL.gen = gmask.ItemGen(yyS[yypt-3].itm, yyS[yypt-1].lst.GetVal())
		}
	case 9:
		//line gmask.y:79
		{
			yyVAL.gen = yyS[yypt-0].gen
		}
	case 10:
		//line gmask.y:80
		{
			yyVAL.gen = gmask.RangeGen(yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 11:
		yyVAL.gen = yyS[yypt-0].gen
	case 12:
		yyVAL.gen = yyS[yypt-0].gen
	case 13:
		//line gmask.y:83
		{
			yyVAL.gen = gmask.MaskGen(yyS[yypt-1].gen, yyS[yypt-0].sli...)
		}
	case 14:
		//line gmask.y:84
		{
			yyVAL.gen = gmask.QuantGen(yyS[yypt-1].gen, yyS[yypt-0].sli...)
		}
	case 15:
		//line gmask.y:85
		{
			yyVAL.gen = gmask.AccumGen(yyS[yypt-2].gen, gmask.ON)
		}
	case 16:
		//line gmask.y:86
		{
			yyVAL.gen = gmask.AccumGen(yyS[yypt-4].gen, gmask.ON, yyS[yypt-0].val)
		}
	case 17:
		//line gmask.y:87
		{
			yyVAL.gen = gmask.AccumGen(yyS[yypt-1].gen, yyS[yypt-0].sli...)
		}
	case 18:
		//line gmask.y:90
		{
			yyVAL.itm = gmask.CYCLE
		}
	case 19:
		//line gmask.y:91
		{
			yyVAL.itm = gmask.SWING
		}
	case 20:
		//line gmask.y:92
		{
			yyVAL.itm = gmask.HEAP
		}
	case 21:
		//line gmask.y:93
		{
			yyVAL.itm = gmask.RANDOM
		}
	case 22:
		//line gmask.y:96
		{
			yyVAL.lst = NewList(yyS[yypt-0].val)
		}
	case 23:
		//line gmask.y:97
		{
			yyS[yypt-1].lst.AddVal(yyS[yypt-0].val)
		}
	case 24:
		//line gmask.y:100
		{
			yyVAL.gen = gmask.BpfGen(yyS[yypt-2].lst.GetVal(), yyS[yypt-1].ipl)
		}
	case 25:
		//line gmask.y:101
		{
			yyVAL.gen = gmask.BpfGen([]float64{yyS[yypt-3].val, yyS[yypt-2].val}, yyS[yypt-1].ipl)
		}
	case 26:
		//line gmask.y:104
		{
			yyVAL.lst = NewBpList(yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 27:
		//line gmask.y:105
		{
			yyS[yypt-2].lst.AddBp(yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 28:
		//line gmask.y:108
		{
			yyVAL.ipl = gmask.NewInterpolation(0, false, false)
		}
	case 29:
		//line gmask.y:109
		{
			yyVAL.ipl = gmask.NewInterpolation(yyS[yypt-0].val, false, false)
		}
	case 30:
		//line gmask.y:110
		{
			yyVAL.ipl = gmask.NewInterpolation(0, true, false)
		}
	case 31:
		//line gmask.y:111
		{
			yyVAL.ipl = gmask.NewInterpolation(0, false, true)
		}
	case 32:
		//line gmask.y:114
		{
			yyVAL.rmd = gmask.UNI
		}
	case 33:
		//line gmask.y:115
		{
			yyVAL.rmd = gmask.LIN
		}
	case 34:
		//line gmask.y:116
		{
			yyVAL.rmd = gmask.RLIN
		}
	case 35:
		//line gmask.y:117
		{
			yyVAL.rmd = gmask.TRI
		}
	case 36:
		//line gmask.y:118
		{
			yyVAL.rmd = gmask.EXP
		}
	case 37:
		//line gmask.y:119
		{
			yyVAL.rmd = gmask.REXP
		}
	case 38:
		//line gmask.y:120
		{
			yyVAL.rmd = gmask.BEXP
		}
	case 39:
		//line gmask.y:121
		{
			yyVAL.rmd = gmask.GAUSS
		}
	case 40:
		//line gmask.y:122
		{
			yyVAL.rmd = gmask.CAUCHY
		}
	case 41:
		//line gmask.y:123
		{
			yyVAL.rmd = gmask.BETA
		}
	case 42:
		//line gmask.y:124
		{
			yyVAL.rmd = gmask.WEI
		}
	case 43:
		//line gmask.y:127
		{
			yyVAL.gen = gmask.RndGen(yyS[yypt-0].rmd)
		}
	case 44:
		//line gmask.y:128
		{
			yyVAL.gen = gmask.RndGen(yyS[yypt-1].rmd, yyS[yypt-0].val)
		}
	case 45:
		//line gmask.y:129
		{
			yyVAL.gen = gmask.RndGen(yyS[yypt-1].rmd, yyS[yypt-0].gen)
		}
	case 46:
		//line gmask.y:130
		{
			yyVAL.gen = gmask.RndGen(yyS[yypt-2].rmd, yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 47:
		//line gmask.y:131
		{
			yyVAL.gen = gmask.RndGen(yyS[yypt-2].rmd, yyS[yypt-1].gen, yyS[yypt-0].val)
		}
	case 48:
		//line gmask.y:132
		{
			yyVAL.gen = gmask.RndGen(yyS[yypt-2].rmd, yyS[yypt-1].val, yyS[yypt-0].gen)
		}
	case 49:
		//line gmask.y:133
		{
			yyVAL.gen = gmask.RndGen(yyS[yypt-2].rmd, yyS[yypt-1].gen, yyS[yypt-0].gen)
		}
	case 50:
		//line gmask.y:136
		{
			yyVAL.omd = gmask.SIN
		}
	case 51:
		//line gmask.y:137
		{
			yyVAL.omd = gmask.COS
		}
	case 52:
		//line gmask.y:138
		{
			yyVAL.omd = gmask.SQUARE
		}
	case 53:
		//line gmask.y:139
		{
			yyVAL.omd = gmask.TRIANGLE
		}
	case 54:
		//line gmask.y:140
		{
			yyVAL.omd = gmask.SAWUP
		}
	case 55:
		//line gmask.y:141
		{
			yyVAL.omd = gmask.SAWDOWN
		}
	case 56:
		//line gmask.y:142
		{
			yyVAL.omd = gmask.POWUP
		}
	case 57:
		//line gmask.y:143
		{
			yyVAL.omd = gmask.POWDOWN
		}
	case 58:
		//line gmask.y:146
		{
			yyVAL.gen = gmask.OscGen(yyS[yypt-1].omd, yyS[yypt-0].val)
		}
	case 59:
		//line gmask.y:147
		{
			yyVAL.gen = gmask.OscGen(yyS[yypt-1].omd, yyS[yypt-0].gen)
		}
	case 60:
		//line gmask.y:148
		{
			yyVAL.gen = gmask.OscGen(yyS[yypt-2].omd, yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 61:
		//line gmask.y:149
		{
			yyVAL.gen = gmask.OscGen(yyS[yypt-2].omd, yyS[yypt-1].gen, yyS[yypt-0].val)
		}
	case 62:
		//line gmask.y:150
		{
			yyVAL.gen = gmask.OscGen(yyS[yypt-3].omd, yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 63:
		//line gmask.y:151
		{
			yyVAL.gen = gmask.OscGen(yyS[yypt-3].omd, yyS[yypt-2].gen, yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 64:
		//line gmask.y:154
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 65:
		//line gmask.y:155
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-1].gen, yyS[yypt-0].val)
		}
	case 66:
		//line gmask.y:156
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-1].val, yyS[yypt-0].gen)
		}
	case 67:
		//line gmask.y:157
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-1].gen, yyS[yypt-0].gen)
		}
	case 68:
		//line gmask.y:158
		{
			yyVAL.sli = append(yyS[yypt-2].sli, yyS[yypt-0].val)
		}
	case 69:
		//line gmask.y:161
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-0].val)
		}
	case 70:
		//line gmask.y:162
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-0].gen)
		}
	case 71:
		//line gmask.y:163
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 72:
		//line gmask.y:164
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-1].val, yyS[yypt-0].gen)
		}
	case 73:
		//line gmask.y:165
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-1].gen, yyS[yypt-0].val)
		}
	case 74:
		//line gmask.y:166
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-1].gen, yyS[yypt-0].gen)
		}
	case 75:
		//line gmask.y:167
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 76:
		//line gmask.y:168
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].val, yyS[yypt-1].val, yyS[yypt-0].gen)
		}
	case 77:
		//line gmask.y:169
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].val, yyS[yypt-1].gen, yyS[yypt-0].val)
		}
	case 78:
		//line gmask.y:170
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].val, yyS[yypt-1].gen, yyS[yypt-0].gen)
		}
	case 79:
		//line gmask.y:171
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].gen, yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 80:
		//line gmask.y:172
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].gen, yyS[yypt-1].val, yyS[yypt-0].gen)
		}
	case 81:
		//line gmask.y:173
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].gen, yyS[yypt-1].gen, yyS[yypt-0].val)
		}
	case 82:
		//line gmask.y:174
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].gen, yyS[yypt-1].gen, yyS[yypt-0].gen)
		}
	case 83:
		//line gmask.y:177
		{
			yyVAL.amd = gmask.LIMIT
		}
	case 84:
		//line gmask.y:178
		{
			yyVAL.amd = gmask.MIRROR
		}
	case 85:
		//line gmask.y:179
		{
			yyVAL.amd = gmask.WRAP
		}
	case 86:
		//line gmask.y:182
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].amd, yyS[yypt-1].val, yyS[yypt-0].val)
		}
	case 87:
		//line gmask.y:183
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].amd, yyS[yypt-1].gen, yyS[yypt-0].val)
		}
	case 88:
		//line gmask.y:184
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].amd, yyS[yypt-1].val, yyS[yypt-0].gen)
		}
	case 89:
		//line gmask.y:185
		{
			yyVAL.sli = NewInterfaceSlice(yyS[yypt-2].amd, yyS[yypt-1].gen, yyS[yypt-0].gen)
		}
	case 90:
		//line gmask.y:186
		{
			yyVAL.sli = append(yyS[yypt-2].sli, yyS[yypt-0].val)
		}
	case 91:
		//line gmask.y:189
		{
			yyVAL.n = -1
		}
	case 92:
		//line gmask.y:190
		{
			yyVAL.n = int(yyS[yypt-0].val)
		}
	}
	goto yystack /* stack new state and value */
}
