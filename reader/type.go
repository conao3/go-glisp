package reader

type Expr interface {
	isExpr()
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
