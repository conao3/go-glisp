package evaluator

import (
	"github.com/conao3/go-glisp/types"
)

func extendedFunctionEnv(arg_syms_ *types.Cons, args_ *types.Cons, env *types.Environment) *types.Environment {
	extended_env := types.NewEnclosedEnvironment(env)

	for {
		arg_sym_ := arg_syms_.Car
		arg_ := args_.Car

		arg_sym := arg_sym_.(*types.Symbol)
		arg := Eval(arg_, env)

		extended_env.SetValue(arg_sym.Name, arg)

		if arg_syms_.Cdr == &types.NIL && args_.Cdr == &types.NIL {
			break
		}
		if arg_syms_.Cdr == &types.NIL || args_.Cdr == &types.NIL {
			panic("argument length mismatch")
		}
		arg_syms_ = arg_syms_.Cdr.(*types.Cons)
		args_ = args_.Cdr.(*types.Cons)
	}

	return extended_env
}

var builtin = map[string]func(args *types.Cons, env *types.Environment) types.Expr{
	"atom":   doAtom,
	"eq":     doEq,
	"car":    doCar,
	"cdr":    doCdr,
	"cons":   doCons,
	"cond":   doCond,
	"quote":  doQuote,
	"lambda": doLambda,
	"set":    doSet,
	"fset":   doFset,
}

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
			fn, ok := env.GetFunction(car.Name)
			if ok {
				if fn.(*types.Cons).Car.(*types.Symbol).Name != "lambda" {
					panic("this is not a function")
				}
				arg_syms := fn.(*types.Cons).Cdr.(*types.Cons).Car.(*types.Cons)
				body := fn.(*types.Cons).Cdr.(*types.Cons).Cdr.(*types.Cons).Car
				args := expr.Cdr.(*types.Cons)

				extended_env := extendedFunctionEnv(arg_syms, args, env)

				return Eval(body, extended_env)
			}

			fn_builtin, ok_builtin := builtin[car.Name]
			if ok_builtin {
				return fn_builtin(expr.Cdr.(*types.Cons), env)
			}

			panic("undefined function")
		case *types.Cons:
			switch caar := car.Car.(type) {
			case *types.Symbol:
				switch caar.Name {
				case "lambda":
					arg_syms := car.Cdr.(*types.Cons).Car.(*types.Cons)
					body := car.Cdr.(*types.Cons).Cdr.(*types.Cons).Car
					args := expr.Cdr.(*types.Cons)

					extended_env := extendedFunctionEnv(arg_syms, args, env)

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
