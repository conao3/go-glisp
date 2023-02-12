package evaluator

import (
	"github.com/conao3/go-glisp/types"
)

func Eval(expr types.Expr, env *types.Environment) types.Expr {
	switch expr := expr.(type) {
	case *types.Symbol:
		r, ok := env.GetValue(expr.Name)
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
			case "set":
				sym_ := expr.Cdr.(*types.Cons).Car
				val_ := expr.Cdr.(*types.Cons).Cdr.(*types.Cons).Car
				sym := Eval(sym_, env).(*types.Symbol)
				val := Eval(val_, env)
				env.SetValue(sym.Name, val)
				return val
			case "fset":
				sym_ := expr.Cdr.(*types.Cons).Car
				val_ := expr.Cdr.(*types.Cons).Cdr.(*types.Cons).Car
				sym := Eval(sym_, env).(*types.Symbol)
				val := Eval(val_, env)
				env.SetFunction(sym.Name, val)
				return val
			default:
				panic("not implemented")
			}
		case *types.Cons:
			switch caar := car.Car.(type) {
			case *types.Symbol:
				switch caar.Name {
				case "lambda":
					arg_syms := car.Cdr.(*types.Cons).Car.(*types.Cons)
					body := car.Cdr.(*types.Cons).Cdr.(*types.Cons).Car
					args_ := expr.Cdr.(*types.Cons)

					extended_env := types.NewEnclosedEnvironment(env)

					for {
						arg_sym_ := arg_syms.Car
						arg_ := args_.Car

						arg_sym := arg_sym_.(*types.Symbol)
						arg := Eval(arg_, env)

						extended_env.SetValue(arg_sym.Name, arg)

						if arg_syms.Cdr == &types.NIL && args_.Cdr == &types.NIL {
							break
						}
						if arg_syms.Cdr == &types.NIL || args_.Cdr == &types.NIL {
							panic("argument length mismatch")
						}
						arg_syms = arg_syms.Cdr.(*types.Cons)
						args_ = args_.Cdr.(*types.Cons)
					}

					return Eval(body, extended_env)
				default:
					panic("not implemented")
				}
			default:
				panic("not implemented")
			}
		default:
			panic("unreachable")
		}
	default:
		panic("unreachable")
	}
}
