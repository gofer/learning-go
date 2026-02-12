package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	add = func(i int, j int) (int, error) { return i + j, nil }
	sub = func(i int, j int) (int, error) { return i - j, nil }
	mul = func(i int, j int) (int, error) { return i * j, nil }
	div = func(i int, j int) (int, error) {
		if j == 0 {
			return 0, errors.New("0で割ることはできません")
		}
		return i / j, nil
	}
)

func exercise01() {
	type opFuncType func(int, int) (int, error)
	var opMap = map[string]opFuncType{
		"+": add,
		"-": sub,
		"*": mul,
		"/": div,
	}

	expressions := [][]string{
		[]string{"2", "+", "3"},
		[]string{"2", "-", "3"},
		[]string{"2", "*", "3"},
		[]string{"2", "/", "3"},
		[]string{"2", "%", "3"},
		[]string{"two", "+", "three"},
		[]string{"2", "+", "three"},
		[]string{"5"},
		[]string{"2", "/", "0"},
	}
	for _, expression := range expressions {
		if len(expression) != 3 {
			fmt.Print(expression, " -- 不正な式です\n")
			continue
		}
		p1, err := strconv.Atoi(expression[0])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		op := expression[1]
		opFunc, ok := opMap[op]
		if !ok {
			fmt.Print(expression, " -- ", "定義されていない演算子です: ", op, "\n")
			continue
		}
		p2, err := strconv.Atoi(expression[2])
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		result, err := opFunc(p1, p2)
		if err != nil {
			fmt.Print(expression, " -- ", err, "\n")
			continue
		}
		fmt.Print(expression, " → ", result, "\n")
	}
}
