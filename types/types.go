package types

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
	Value int
}

func (i *Int) isExpr() {}
func (i *Int) isAtom() {}
func (i *Int) String() string {
	return fmt.Sprintf("%d", i.Value)
}

type Symbol struct {
	Atom
	Name string
}

func (s *Symbol) isExpr() {}
func (s *Symbol) isAtom() {}
func (s *Symbol) String() string {
	return s.Name
}

type Cons struct {
	Expr
	Car Expr
	Cdr Expr
}

func (c *Cons) isExpr() {}
func (c *Cons) String() string {
	var buf bytes.Buffer
	cur := c
	fmt.Fprint(&buf, "(")
	for {
		fmt.Fprint(&buf, cur.Car)
		if cur.Cdr == &NIL {
			break
		}
		if atom, ok := cur.Cdr.(Atom); ok {
			fmt.Fprint(&buf, " . ", atom)
			break
		}
		fmt.Fprint(&buf, " ")
		cur = cur.Cdr.(*Cons)
	}
	fmt.Fprint(&buf, ")")
	return buf.String()
}

var (
	T   = Symbol{Name: "t"}
	NIL = Symbol{Name: "nil"}
)

type Environment struct {
	outer *Environment
	store map[string]Expr
}

func NewEnvironment() *Environment {
	return &Environment{store: make(map[string]Expr)}
}

func (e *Environment) Get(name string) (Expr, bool) {
	val, ok := e.store[name]
	if !ok && e.outer != nil {
		return e.outer.Get(name)
	}
	return val, ok
}

func (e *Environment) Set(name string, val Expr) Expr {
	e.store[name] = val
	return val
}
