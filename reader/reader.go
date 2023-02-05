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

func (r *Reader) readList() []string {
	var list []string
	for r.chr != ')' {
		list = append(list, r.readExpr()...)
	}
	r.readChar()  // skip ')'
	return list
}

func (r *Reader) readSymbol() string {
	pos := r.position
	for GetSyntaxType(r.chr) == Constituent {
		r.readChar()
	}
	return r.input[pos:r.position]
}

func (r *Reader) readExpr() []string {
	for r.chr != 0 {
		switch GetSyntaxType(r.chr) {
		case Invalid:
			panic("got <invalid>; invalid character")

		case Whitespace:
			r.readChar()
			continue

		case TerminatingMacro, NonTerminatingMacro:
			if r.chr == '(' {
				r.readChar()  // skip '('
				return r.readList()
			}
			panic("got <macro>; not implemented")

		case SingleEscape:
			panic("got <single escape>; not implemented")

		case MultipleEscape:
			panic("got <multiple escape>; not implemented")

		case Constituent:
			return []string{r.readSymbol()}
		}
	}
	panic("Unexpected EOF")
}

func (r *Reader) Read() []string {
	return r.readExpr()
}
