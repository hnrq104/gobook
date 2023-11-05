package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gopl.io/ch7/eval"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	fmt.Printf("present expression:")
	input.Scan()
	expr, err := eval.Parse(input.Text())
	if err != nil {
		log.Fatalf("found a problem during parsing: %v", err)
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		log.Fatalf("found a problem during checking: %v", err)
	}

	fmt.Println("present values in right format, e.g. 'x=3.14 y=34 z=25'")
	input.Scan()
	env := make(eval.Env)
	for _, str := range strings.Split(input.Text(), " ") {

		parts := strings.Split(str, "=")
		var v eval.Var = eval.Var(parts[0])
		var floatString string = parts[1]

		if _, ok := vars[v]; !ok {
			log.Fatalf("variable %s not in expression\n", v)
		}

		if value, ok := env[v]; ok {
			log.Fatalf("variable %s already has assigned value %v", v, value)
		}

		f, err := strconv.ParseFloat(floatString, 64)
		if err != nil {
			log.Fatalf("problem parsing float from %s: %v", floatString, err)
		}

		env[v] = f
	}
	fmt.Printf("%v => %g\n", env, expr.Eval(env))
}
