package lexer

type Lexer struct {
	input        string
	position     int
	readPosition int
	chr          byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.chr = 0
	} else {
		l.chr = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
