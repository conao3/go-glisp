package reader

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

func (r *Reader) readList() Expr {
	if r.chr == ')' {
		r.readChar() // skip ')'
		return &NIL
	}

	lst := &Cons{car: r.readExpr(), cdr: &NIL}
	cur := lst
	for r.chr != ')' {
		cur.cdr = &Cons{car: r.readExpr(), cdr: &NIL}
		cur = cur.cdr.(*Cons)
	}
	r.readChar() // skip ')'
	return lst
}

func (r *Reader) readSymbol() Expr {
	pos := r.position
	for GetSyntaxType(r.chr) == Constituent {
		r.readChar()
	}
	name := r.input[pos:r.position]
	return &Symbol{name: name}
}

func (r *Reader) readExpr() Expr {
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

func (r *Reader) Read() Expr {
	return r.readExpr()
}
