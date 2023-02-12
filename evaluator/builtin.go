package evaluator

import (
	"github.com/conao3/go-glisp/types"
)

func doAtom(args *types.Cons, env *types.Environment) types.Expr {
	cadr := args.Car
	_, ok := Eval(cadr, env).(types.Atom)
	if ok {
		return &types.T
	}
	return &types.NIL
}

func doEq(args *types.Cons, env *types.Environment) types.Expr {
	lhs := args.Car
	rhs := args.Cdr.(*types.Cons).Car
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
}

func doCar(args *types.Cons, env *types.Environment) types.Expr {
	arg := args.Car
	return Eval(arg, env).(*types.Cons).Car
}

func doCdr(args *types.Cons, env *types.Environment) types.Expr {
	arg := args.Car
	return Eval(arg, env).(*types.Cons).Cdr
}

func doCons(args *types.Cons, env *types.Environment) types.Expr {
	lhs := args.Car
	rhs := args.Cdr.(*types.Cons).Car
	return &types.Cons{
		Car: Eval(lhs, env),
		Cdr: Eval(rhs, env),
	}
}

func doCond(args *types.Cons, env *types.Environment) types.Expr {
	for {
		cond := args.Car.(*types.Cons)
		if Eval(cond.Car, env) != &types.NIL {
			return Eval(cond.Cdr.(*types.Cons).Car, env)
		}
		cur := args.Cdr
		if cur == &types.NIL {
			return &types.NIL
		}
		args = cur.(*types.Cons).Cdr.(*types.Cons)
	}
}

func doQuote(args *types.Cons, env *types.Environment) types.Expr {
	return args.Car
}

func doLambda(args *types.Cons, env *types.Environment) types.Expr {
	return &types.Cons{
		Car: &types.Symbol{Name: "lambda"},
		Cdr: args,
	}
}

func doSet(args *types.Cons, env *types.Environment) types.Expr {
	sym_ := args.Car
	val_ := args.Cdr.(*types.Cons).Car
	sym := Eval(sym_, env).(*types.Symbol)
	val := Eval(val_, env)
	env.SetValue(sym.Name, val)
	return val
}

func doFset(args *types.Cons, env *types.Environment) types.Expr {
	sym_ := args.Car
	val_ := args.Cdr.(*types.Cons).Car
	sym := Eval(sym_, env).(*types.Symbol)
	val := Eval(val_, env)
	env.SetFunction(sym.Name, val)
	return val
}
