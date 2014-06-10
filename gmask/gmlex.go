package main

import (
	"bytes"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"text/scanner"
)

var keywords = map[string]int{
	"f":        F,
	"ipl":      IPL,
	"cos":      COS,
	"off":      OFF,
	"prec":     PREC,
	"const":    CONST,
	"item":     ITEM,
	"cycle":    CYCLE,
	"swing":    SWING,
	"heap":     HEAP,
	"random":   RANDOM,
	"seg":      SEG,
	"range":    RANGE,
	"rnd":      RND,
	"uni":      UNI,
	"lin":      LIN,
	"rlin":     RLIN,
	"tri":      TRI,
	"exp":      EXP,
	"rexp":     REXP,
	"bexp":     BEXP,
	"gauss":    GAUSS,
	"cauchy":   CAUCHY,
	"beta":     BETA,
	"wei":      WEI,
	"osc":      OSC,
	"sin":      SIN,
	"square":   SQUARE,
	"triangle": TRIANGLE,
	"sawup":    SAWUP,
	"sawdown":  SAWDOWN,
	"powup":    POWUP,
	"powdown":  POWDOWN,
	"mask":     MASK,
	"map":      MAP,
	"quant":    QUANT,
	"accum":    ACCUM,
	"on":       ON,
	"limit":    LIMIT,
	"mirror":   MIRROR,
	"wrap":     WRAP,
	"init":     INIT,
}

type Lexer struct {
	scanner.Scanner
}

var validPField = regexp.MustCompile("p[1-9][0-9]*")

func (l *Lexer) Lex(lval *yySymType) int {
	tok := l.Scan()
	if tok == scanner.Float || tok == scanner.Int {
		lval.val, _ = strconv.ParseFloat(l.TokenText(), 64)
		return NUMBER
	}
	if tok == scanner.EOF {
		return 0
	}
	if tok == scanner.Ident {
		if _, ok := keywords[l.TokenText()]; ok {
			return keywords[l.TokenText()]
		} else if validPField.MatchString(l.TokenText()) {
			lval.n, _ = strconv.Atoi(l.TokenText()[1:len(l.TokenText())])
			return PARAM
		}
		return -1
	}
	c := int(tok)
	switch c {
	case '-', '+':
		tok = l.Scan()
		if tok == scanner.Float || tok == scanner.Int {
			lval.val, _ = strconv.ParseFloat(l.TokenText(), 64)
			if c == '-' {
				lval.val = -lval.val
			}
			return NUMBER
		}
		if tok == scanner.EOF {
			return 0
		}
		return -1
	case '(', '[', ']', ')', '\n':
		return c
	case '{':
		var b bytes.Buffer
		for {
			r := l.Next()
			if r == '}' {
				sco := b.String()
				lval.sco = strings.Replace(sco, "//", ";", -1)
				return SCO
			}
			if r == scanner.EOF {
				return -1
			}
			b.WriteRune(r)
		}
	}
	return -1
}

func (l *Lexer) Error(s string) {
	fmt.Printf("syntax error: %s\n", s)
}

func NewLexer(text []byte) *Lexer {
	text = bytes.Replace(text, []byte(";"), []byte("//"), -1)
	var l Lexer
	l.Init(bytes.NewReader(text))
	l.Mode = scanner.ScanIdents | scanner.ScanFloats | scanner.ScanComments | scanner.SkipComments
	l.Whitespace = 1<<'\t' | 1<<'\r' | 1<<'\n' | 1<<' '
	return &l
}
