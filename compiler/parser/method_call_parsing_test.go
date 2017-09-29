package parser

import (
	"github.com/goby-lang/goby/compiler/ast"
	"github.com/goby-lang/goby/compiler/lexer"
	"testing"
)

func TestCallExpressionWithKeywordArgument(t *testing.T) {
	tests := []struct {
		input    string
		expected map[string]interface{}
	}{
		{`
		add(x: 111)
		`, map[string]interface{}{
			"x": 111,
		}},
		{`
		add(x:)
		`, map[string]interface{}{
			"x": nil,
		}},
		{`
		add(x: 111, y:)
		`, map[string]interface{}{
			"x": 111,
			"y": nil,
		}},
		{`
		add(x: 111, y: 222, z: 333)
		`, map[string]interface{}{
			"x": 111,
			"y": 222,
			"z": 333,
		}},
	}

	for i, tt := range tests {
		l := lexer.New(tt.input)
		p := New(l)
		program, err := p.ParseProgram()

		if err != nil {
			t.Fatalf("At case %d "+err.Message, i)
		}

		stmt := program.Statements[0].(*ast.ExpressionStatement)
		c := stmt.Expression.(*ast.CallExpression)

		h, ok := c.Arguments[0].(*ast.HashExpression)
		if !ok {
			t.Fatalf("At case %d c.Statments[0] is not ast.HashExpression. got=%T", i, c.Arguments[0])
		}

		for k, expected := range tt.expected {
			if expected == nil {
				if h.Data[k] != nil {
					t.Fatalf("Expect argument content to be empty. got=%v", h.Data[k])
				}
			} else {
				testIntegerLiteral(t, h.Data[k], expected.(int))
			}
		}
	}
}
