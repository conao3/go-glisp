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
			case "atom":
				cadr := expr.Cdr.(*types.Cons).Car
				_, ok := Eval(cadr, env).(types.Atom)
				if ok {
					return &types.T
				} else {
					return &types.NIL
				}
			case "eq":
				lhs := expr.Cdr.(*types.Cons).Car
				rhs := expr.Cdr.(*types.Cons).Cdr.(*types.Cons).Car
				lhs_atom, lhs_ok := Eval(lhs, env).(*types.Symbol)
				rhs_atom, rhs_ok := Eval(rhs, env).(*types.Symbol)
				if lhs_ok && rhs_ok {
					if lhs_atom.Name == rhs_atom.Name {
						return &types.T
					} else {
						return &types.NIL
					}
				} else {
					return &types.NIL
				}
			case "car":
				cadr := expr.Cdr.(*types.Cons).Car
				return Eval(cadr, env).(*types.Cons).Car
			case "cdr":
				cadr := expr.Cdr.(*types.Cons).Car
				return Eval(cadr, env).(*types.Cons).Cdr
			case "cons":
				lhs := expr.Cdr.(*types.Cons).Car
				rhs := expr.Cdr.(*types.Cons).Cdr.(*types.Cons).Car
				return &types.Cons{Car: Eval(lhs, env), Cdr: Eval(rhs, env)}
			case "cond":
				cur := expr.Cdr
				for {
					if c, ok := cur.(*types.Symbol); ok && c == &types.NIL {
						return &types.NIL
					}
					cur_ := cur.(*types.Cons)
					pair := cur_.Car.(*types.Cons)
					if r, ok := Eval(pair.Car, env).(*types.Symbol); ok && r != &types.NIL {
						return Eval(pair.Cdr.(*types.Cons).Car, env)
					}
					cur = cur_.Cdr
				}
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
