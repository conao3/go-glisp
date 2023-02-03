package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	LPAREN  = "("
	RPAREN  = ")"

	FUNCTION = "FUNCTION"
)

type Tokenizer struct {
	input        string
	position     int
	readPosition int
	chr 		 byte
}

func New(input string) *Tokenizer {
	l := &Tokenizer{input: input}
	l.readChar()
	return l
}

func (l *Tokenizer) readChar() {
	if l.readPosition >= len(l.input) {
		l.chr = 0
	} else {
		l.chr = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}
