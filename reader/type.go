package reader

import (
	"bytes"
	"fmt"
)

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
	var buf bytes.Buffer
	cur := c
	fmt.Fprint(&buf, "(")
	for {
		fmt.Fprint(&buf, cur.car)
		if cur.cdr == &NIL {
			break
		}
		if atom, ok := cur.cdr.(Atom); ok {
			fmt.Fprint(&buf, " . ", atom)
			break
		}
		fmt.Fprint(&buf, " ")
		cur = cur.cdr.(*Cons)
	}
	fmt.Fprint(&buf, ")")
	return buf.String()
}

var (
	T   = Symbol{name: "t"}
	NIL = Symbol{name: "nil"}
)
