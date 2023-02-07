package reader

import "fmt"

type Expr interface {
	isExpr()
	String() string
}

type Atom interface {
	Expr
	isAtom()
}

type Int struct {
	Atom
	value int
}

func (i *Int) isExpr() {}
func (i *Int) isAtom() {}
func (i *Int) String() string {
	return fmt.Sprintf("%d", i.value)
}

type Symbol struct {
	Atom
	name string
}

func (s *Symbol) isExpr() {}
func (s *Symbol) isAtom() {}
func (s *Symbol) String() string {
	return s.name
}

type Cons struct {
	Expr
	car Expr
	cdr Expr
}

func (c *Cons) isExpr() {}
func (c *Cons) String() string {
	return fmt.Sprintf("(%s . %s)", c.car, c.cdr)
}

var (
	T = Symbol{name: "t"}
	NIL = Symbol{name: "nil"}
)
