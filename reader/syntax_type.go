package reader

type SyntaxType int

const (
	Whitespace SyntaxType = iota
	Constituent
	TerminatingMacro
	NonTerminatingMacro
	SingleEscape
	MultipleEscape
	Invalid
)

func GetSyntaxType(chr byte) SyntaxType {
	switch {
	case syntaxTypeWhitespace(chr):
		return Whitespace
	case syntaxTypeConstituent(chr):
		return Constituent
	case syntaxTypeTerminatingMacro(chr):
		return TerminatingMacro
	case chr == '#':
		return NonTerminatingMacro
	case chr == '\\':
		return SingleEscape
	case chr == '|':
		return MultipleEscape
	}
	return Invalid
}

func syntaxTypeWhitespace(chr byte) bool {
	return chr == '\t' || chr == '\n' || chr == '\r' || chr == ' '
}

func syntaxTypeConstituent(chr byte) bool {
	arr := []byte{
		// constituent
		'_', '-', ':', '.', '@', '*', '/', '&', '%', '^', '+', '<', '=', '>', '~', '$',
		// constituent*
		'!', '?', '[', ']', '{', '}',
	}
	for _, v := range arr {
		if chr == v {
			return true
		}
	}
	if 'a' <= chr && chr <= 'z' || 'A' <= chr && chr <= 'Z' || '0' <= chr && chr <= '9' {
		return true
	}

	return false
}

func syntaxTypeTerminatingMacro(chr byte) bool {
	arr := []byte{
		',', ';', '\'', '"', '(', ')', '`',
	}
	for _, v := range arr {
		if chr == v {
			return true
		}
	}
	return false
}
