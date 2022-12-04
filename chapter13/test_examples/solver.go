package test_examples

import (
	"context"
	"errors"
	"io"
	"strings"
	"testing"
)

type Processor struct {
	Solver MathSolver
}

type MathSolver interface {
	Resolve(ctx context.Context, expression string) (float64, error)
}

func (p Processor) ProcessExpresion(ctx context.Context, r io.Reader) (float64, error) {
	curExpression, err := readToNewLine(r)
	if err != nil {
		return 0, err
	}
	if len(curExpression) == 0 {
		return 0, errors.New("no expression to read")
	}
	answer, err := p.Solver.Resolve(ctx, curExpression)
	return answer, err
}

type MathSolverStub struct{}

func (ms MathSolverStub) Resolve(ctx context.Context, expr string) (float64, error) {
	switch expr {
	case "2 + 2 * 10":
		return 22, nil
	case "( 2 + 2 ) * 10":
		return 40, nil
	case "( 2 + 2 * 10 ":
		return 0, errors.New("invalid expression: ( 2 + 2 * 10")
	}
	return 0, nil
}

func TestProcessExpression(t *testing.T) {
	p := Processor{MathSolverStub{}}
	in := strings.NewReader(`2 + 2 * 10
    ( 2 + 2 ) * 10
	( 2 + 2 * 10`)
	data := []float64{22, 40, 0}
	hasErr := []bool{false, false, true}
	for i, d := range data {
		result, err := p.ProcessExpresion(context.Background(), in)
		if err != nil && !hasErr[i] {
			t.Error(err)
		}
		if result != d {
			t.Errorf("Expected result %f, got %f", d, result)
		}
	}

}
