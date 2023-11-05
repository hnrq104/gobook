package main

import (
	"fmt"
	"gobook/ch7/surfserver/surface"
	"log"
	"math"
	"net/http"

	"gopl.io/ch7/eval"
)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}

	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}

	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}

	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}

	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Surface(w, func(a, b float64) float64 {
		r := math.Hypot(a, b) // distance from 0.0
		return expr.Eval(eval.Env{"x": a, "y": b, "r": r})
	}, "")
}

func main() {
	http.HandleFunc("/", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
