package reader

import (
	"regexp"
	"strconv"

	"github.com/conao3/go-glisp/types"
)

type Reader struct {
	input        string
	position     int
	readPosition int
	chr          byte
}

func New(input string) *Reader {
	r := &Reader{input: input}
	r.readChar()
	return r
}

func (r *Reader) readChar() {
	if r.readPosition >= len(r.input) {
		r.chr = 0
	} else {
		r.chr = r.input[r.readPosition]
	}
	r.position = r.readPosition
	r.readPosition += 1
}

func (r *Reader) peakChar() byte {
	if r.readPosition >= len(r.input) {
		return 0
	}
	return r.input[r.readPosition]
}

func (r *Reader) skipWhitespace() {
	for GetSyntaxType(r.chr) == Whitespace {
		r.readChar()
	}
}

func (r *Reader) readList() types.Expr {
	if r.chr == ')' {
		r.readChar() // skip ')'
		return &types.NIL
	}

	lst := &types.Cons{Car: r.readExpr(), Cdr: &types.NIL}
	cur := lst

	r.skipWhitespace()

L:
	for r.chr != ')' {
		cur.Cdr = &types.Cons{Car: r.readExpr(), Cdr: &types.NIL}
		cur = cur.Cdr.(*types.Cons)

		r.skipWhitespace()
		switch r.chr {
		case 0:
			panic("Unexpected EOF")
		case '.':
			if GetSyntaxType(r.peakChar()) == Whitespace {
				r.readChar() // skip '.'

				cur.Cdr = r.readExpr()

				r.skipWhitespace()
				if r.chr != ')' {
					panic("Expected ')'")
				}

				break L
			}
		}
	}
	r.readChar() // skip ')'
	return lst
}

func (r *Reader) readSymbol() types.Expr {
	pos := r.position
	for GetSyntaxType(r.chr) == Constituent {
		r.readChar()
	}
	name := r.input[pos:r.position]

	if regexp.MustCompile(`[0-9]+`).Match([]byte(name)) {
		num, err := strconv.Atoi(name)
		if err != nil {
			panic(err)
		}
		return &types.Int{Value: num}
	}
	if name == "t" {
		return &types.T
	}
	if name == "nil" {
		return &types.NIL
	}
	return &types.Symbol{Name: name}
}

func (r *Reader) readExpr() types.Expr {
	for r.chr != 0 {
		switch GetSyntaxType(r.chr) {
		case Invalid:
			panic("got <invalid>; invalid character")

		case Whitespace:
			r.readChar()
			continue

		case TerminatingMacro, NonTerminatingMacro:
			if r.chr == '(' {
				r.readChar() // skip '('
				return r.readList()
			}
			panic("got <macro>; not implemented")

		case SingleEscape:
			panic("got <single escape>; not implemented")

		case MultipleEscape:
			panic("got <multiple escape>; not implemented")

		case Constituent:
			return r.readSymbol()
		}
	}
	panic("Unexpected EOF")
}

func (r *Reader) Read() types.Expr {
	return r.readExpr()
}
