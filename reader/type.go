package reader

import "fmt"

type Expr interface {
	isExpr()
	String() string
}

func (e *Expr) String() string {
	switch e.(type) {
	case *Int:
		return fmt.Sprintf("%d", e.(Int).value)
	case *Symbol:
		return e.(Symbol).name
	case *Cons:
		return fmt.Sprintf("(%s . %s)", e.(Cons).car, e.(Cons).cdr)
	default:
		panic("Not implemented")
	}
}

type Atom interface {
	Expr
	isAtom()
}

type Int struct {
	Atom
	value int
}

type Symbol struct {
	Atom
	name string
}

type Cons struct {
	Expr
	car Expr
	cdr Expr
}

var (
	T = Symbol{name: "t"}
	NIL = Symbol{name: "nil"}
)
