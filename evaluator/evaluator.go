package evaluator

import (
	"github.com/conao3/go-glisp/types"
)

func Eval(expr types.Expr, env *types.Environment) types.Expr {
	switch expr := expr.(type) {
	case *types.Symbol:
		r, ok := env.Get(expr.Name)
		if !ok {
			panic("undefined symbol")
		}
		return r
	case types.Atom:
		return expr
	case *types.Cons:
		switch car := expr.Car.(type) {
		case *types.Symbol:
			switch car.Name {
			case "quote":
				return expr.Cdr.(*types.Cons).Car
			default:
				panic("not implemented")
			}
		case *types.Cons:
			panic("not implemented")
		default:
			panic("unreachable")
		}
	default:
		panic("unreachable")
	}
}
