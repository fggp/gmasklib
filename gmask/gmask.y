// Gmask parser.

%{

package main

import (
	"bufio"
	"fmt"
	"github.com/fggp/gmask"
	"io"
	"io/ioutil"
	"log"
	"os"
)

%}

%union {
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

%token SCO
%token NUMBER
%token F PARAM
%token IPL COS OFF
%token PREC
%token ITEM CYCLE SWING HEAP RANDOM
%token CONST SEG RANGE
%token RND UNI LIN RLIN TRI EXP REXP BEXP GAUSS CAUCHY BETA WEI
%token OSC SIN SQUARE TRIANGLE SAWUP SAWDOWN POWUP POWDOWN
%token MASK MAP QUANT
%token ACCUM ON LIMIT MIRROR WRAP INIT

%type <sco> SCO, csco
%type <val> NUMBER
%type <n>   PARAM prec
%type <fie> field
%type <par> param
%type <gen> generator bpf rndgen oscgen
%type <itm> item
%type <lst> list bplist
%type <ipl> ipl
%type <rmd> rnd
%type <omd> osc
%type <sli> mask quant accum
%type <amd> accummode

%% // Grammar rules and actions follow.

input: csco     { fmt.Fprintf(w, "%s\n", $1) }
| input field	{ fieldNum++; $2.EvalToScore(w, fieldNum) }
;

csco: SCO       { $$ = $1 }
;

field: F NUMBER NUMBER    { $$ = gmask.NewField($2, $3) }
| field param             { $1.AddParam($2) }
;

param: PARAM generator prec   { $$ = gmask.NewParam($1, $2, $3) }
;

generator: CONST NUMBER         { $$ = gmask.ConstGen($2) }
| item '(' list ')'             { $$ = gmask.ItemGen($1, $3.GetVal()) }
| SEG bpf                       { $$ = $2 }
| RANGE NUMBER NUMBER           { $$ = gmask.RangeGen($2, $3) }
| rndgen
| oscgen
| generator mask                  { $$ = gmask.MaskGen($1, $2...) }
| generator quant                 { $$ = gmask.QuantGen($1, $2...) }
| generator ACCUM ON              { $$ = gmask.AccumGen($1, gmask.ON) }
| generator ACCUM ON INIT NUMBER  { $$ = gmask.AccumGen($1, gmask.ON, $5) }
| generator accum                 { $$ = gmask.AccumGen($1, $2...) }
;

item: ITEM CYCLE          { $$ = gmask.CYCLE }
| ITEM SWING              { $$ = gmask.SWING }
| ITEM HEAP               { $$ = gmask.HEAP }
| ITEM RANDOM             { $$ = gmask.RANDOM }
;

list: NUMBER              { $$ = NewList($1) } 
| list NUMBER             { $1.AddVal($2) }
;

bpf: '(' bplist ipl ')'      { $$ = gmask.BpfGen($2.GetVal(), $3) }
| '[' NUMBER NUMBER ipl ']'  { $$ = gmask.BpfGen([]float64{$2, $3}, $4) }
;

bplist: NUMBER NUMBER     { $$ = NewBpList($1, $2) } 
| bplist NUMBER NUMBER    { $1.AddBp($2, $3) }
;

ipl: /* empty */          { $$ = gmask.NewInterpolation(0, false, false) }
| IPL NUMBER              { $$ = gmask.NewInterpolation($2, false, false) }
| IPL COS                 { $$ = gmask.NewInterpolation(0, true, false) }
| IPL OFF                 { $$ = gmask.NewInterpolation(0, false, true) }
;

rnd: RND UNI    { $$ = gmask.UNI }
| RND LIN       { $$ = gmask.LIN }
| RND RLIN      { $$ = gmask.RLIN }
| RND TRI       { $$ = gmask.TRI }
| RND EXP       { $$ = gmask.EXP }
| RND REXP      { $$ = gmask.REXP }
| RND BEXP      { $$ = gmask.BEXP }
| RND GAUSS     { $$ = gmask.GAUSS }
| RND CAUCHY    { $$ = gmask.CAUCHY }
| RND BETA      { $$ = gmask.BETA }
| RND WEI       { $$ = gmask.WEI }
;

rndgen: rnd            { $$ = gmask.RndGen($1) }
| rnd NUMBER           { $$ = gmask.RndGen($1, $2) }
| rnd bpf              { $$ = gmask.RndGen($1, $2) }
| rnd NUMBER NUMBER    { $$ = gmask.RndGen($1, $2, $3) }
| rnd bpf NUMBER       { $$ = gmask.RndGen($1, $2, $3) }
| rnd NUMBER bpf       { $$ = gmask.RndGen($1, $2, $3) }
| rnd bpf bpf          { $$ = gmask.RndGen($1, $2, $3) }
;

osc: OSC SIN      { $$ = gmask.SIN }
| OSC COS         { $$ = gmask.COS }
| OSC SQUARE      { $$ = gmask.SQUARE }
| OSC TRIANGLE    { $$ = gmask.TRIANGLE }
| OSC SAWUP       { $$ = gmask.SAWUP }
| OSC SAWDOWN     { $$ = gmask.SAWDOWN }
| OSC POWUP       { $$ = gmask.POWUP }
| OSC POWDOWN     { $$ = gmask.POWDOWN }
;

oscgen: osc NUMBER          { $$ = gmask.OscGen($1, $2) }
| osc bpf                   { $$ = gmask.OscGen($1, $2) }
| osc NUMBER NUMBER         { $$ = gmask.OscGen($1, $2, $3) }
| osc bpf NUMBER            { $$ = gmask.OscGen($1, $2, $3) }
| osc NUMBER NUMBER NUMBER  { $$ = gmask.OscGen($1, $2, $3, $4) }
| osc bpf NUMBER NUMBER     { $$ = gmask.OscGen($1, $2, $3, $4) }
;

mask: MASK NUMBER NUMBER    { $$ = NewInterfaceSlice($2, $3) }
| MASK bpf NUMBER           { $$ = NewInterfaceSlice($2, $3) }
| MASK NUMBER bpf           { $$ = NewInterfaceSlice($2, $3) }
| MASK bpf bpf              { $$ = NewInterfaceSlice($2, $3) }
| mask MAP NUMBER           { $$ = append($1, $3) }
;

quant: QUANT NUMBER            { $$ = NewInterfaceSlice($2) }
| QUANT bpf                    { $$ = NewInterfaceSlice($2) }
| QUANT NUMBER NUMBER          { $$ = NewInterfaceSlice($2, $3) }
| QUANT NUMBER bpf             { $$ = NewInterfaceSlice($2, $3) }
| QUANT bpf NUMBER             { $$ = NewInterfaceSlice($2, $3) }
| QUANT bpf bpf                { $$ = NewInterfaceSlice($2, $3) }
| QUANT NUMBER NUMBER NUMBER   { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT NUMBER NUMBER bpf      { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT NUMBER bpf NUMBER      { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT NUMBER bpf bpf         { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT bpf NUMBER NUMBER      { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT bpf NUMBER bpf         { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT bpf bpf NUMBER         { $$ = NewInterfaceSlice($2, $3, $4) }
| QUANT bpf bpf bpf            { $$ = NewInterfaceSlice($2, $3, $4) }
;

accummode: ACCUM LIMIT    { $$ = gmask.LIMIT }
| ACCUM MIRROR            { $$ = gmask.MIRROR }
| ACCUM WRAP              { $$ = gmask.WRAP }
;

accum: accummode NUMBER NUMBER  { $$ = NewInterfaceSlice($1, $2, $3) }
| accummode bpf NUMBER          { $$ = NewInterfaceSlice($1, $2, $3) }
| accummode NUMBER bpf          { $$ = NewInterfaceSlice($1, $2, $3) }
| accummode bpf bpf             { $$ = NewInterfaceSlice($1, $2, $3) }
| accum INIT NUMBER             { $$ = append($1, $3) }
;

prec: /* empty */         { $$ = -1 }
| PREC NUMBER             { $$ = int($2) }
;

%%
	
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
