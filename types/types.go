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
	values map[string]Expr
	functions map[string]Expr
}

func NewEnvironment() *Environment {
	return &Environment{
		outer: nil,
		values: make(map[string]Expr),
		functions: make(map[string]Expr),
	}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) GetValue(name string) (Expr, bool) {
	val, ok := e.values[name]
	if !ok && e.outer != nil {
		return e.outer.GetValue(name)
	}
	return val, ok
}

func (e *Environment) SetValue(name string, val Expr) Expr {
	e.values[name] = val
	return val
}

func (e *Environment) GetFunction(name string) (Expr, bool) {
	val, ok := e.functions[name]
	if !ok && e.outer != nil {
		return e.outer.GetFunction(name)
	}
	return val, ok
}

func (e *Environment) SetFunction(name string, val Expr) Expr {
	e.functions[name] = val
	return val
}
