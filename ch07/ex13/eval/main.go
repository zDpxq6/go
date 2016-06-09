package main

import (
	"fmt"
)

func main() {

	a := Var("x")
	b := literal(3.14)
	c := unary{'-', a}
	d := binary{x: a, op: '*', y: b}
	e := call{"sin", []Expr{a}}
	exprs := []Expr{a, b, c, d, e}
	x := make([]string, 10)
	for _, e := range exprs {
		x = append(x, e.String())
		fmt.Println(e.String())
	}
	for _, e := range x {
		expl,err := Parse(e)
		if err == nil {
			fmt.Println(expl.String())
		}
	}
}
