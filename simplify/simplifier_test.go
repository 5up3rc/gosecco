package simplify

import (
	"testing"

	"github.com/twtiger/go-seccomp/tree"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type SimplifierSuite struct{}

var _ = Suite(&SimplifierSuite{})

func (s *SimplifierSuite) Test_simplifyAddition(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.PLUS, Left: tree.NumericLiteral{1}, Right: tree.NumericLiteral{2}})
	c.Assert(tree.ExpressionString(sx), Equals, "3")
}

func (s *SimplifierSuite) Test_simplifySubtraction(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.MINUS, Left: tree.NumericLiteral{32}, Right: tree.NumericLiteral{3}})
	c.Assert(tree.ExpressionString(sx), Equals, "29")
}

func (s *SimplifierSuite) Test_simplifyMult(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.MULT, Left: tree.NumericLiteral{12}, Right: tree.NumericLiteral{3}})
	c.Assert(tree.ExpressionString(sx), Equals, "36")
}

func (s *SimplifierSuite) Test_simplifyDiv(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.DIV, Left: tree.NumericLiteral{37}, Right: tree.NumericLiteral{3}})
	c.Assert(tree.ExpressionString(sx), Equals, "12")
}

func (s *SimplifierSuite) Test_simplifyMod(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.MOD, Left: tree.NumericLiteral{37}, Right: tree.NumericLiteral{3}})
	c.Assert(tree.ExpressionString(sx), Equals, "1")
}

func (s *SimplifierSuite) Test_simplifyBinAnd(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.BINAND, Left: tree.NumericLiteral{7}, Right: tree.NumericLiteral{4}})
	c.Assert(tree.ExpressionString(sx), Equals, "4")
}

func (s *SimplifierSuite) Test_simplifyBinOr(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.BINOR, Left: tree.NumericLiteral{3}, Right: tree.NumericLiteral{8}})
	c.Assert(tree.ExpressionString(sx), Equals, "11")
}

func (s *SimplifierSuite) Test_simplifyBinXoe(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.BINXOR, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{12}})
	c.Assert(tree.ExpressionString(sx), Equals, "38")
}

func (s *SimplifierSuite) Test_simplifyLsh(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.LSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}})
	c.Assert(tree.ExpressionString(sx), Equals, "168")
}

func (s *SimplifierSuite) Test_simplifyRsh(c *C) {
	sx := Simplify(tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{84}, Right: tree.NumericLiteral{2}})
	c.Assert(tree.ExpressionString(sx), Equals, "21")
}

func (s *SimplifierSuite) Test_simplifyCall(c *C) {
	sx := Simplify(tree.Call{"foo", []tree.Any{tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{84}, Right: tree.NumericLiteral{2}}}})
	c.Assert(tree.ExpressionString(sx), Equals, "(foo 21)")
}

func (s *SimplifierSuite) Test_simplifyAnd(c *C) {
	sx := Simplify(tree.And{
		tree.Comparison{
			Left:  tree.Arithmetic{Op: tree.LSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
			Op:    tree.EQL,
			Right: tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		},
		tree.Comparison{
			Left:  tree.Argument{2},
			Op:    tree.EQL,
			Right: tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		},
	})
	c.Assert(tree.ExpressionString(sx), Equals, "false")

	sx = Simplify(tree.And{
		tree.Comparison{
			Left:  tree.Arithmetic{Op: tree.LSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
			Op:    tree.EQL,
			Right: tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{42}, Right: tree.Variable{"foo"}},
		},
		tree.Comparison{
			Left:  tree.Argument{2},
			Op:    tree.EQL,
			Right: tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		},
	})
	c.Assert(tree.ExpressionString(sx), Equals, "(and (eq 168 (rsh 42 foo)) (eq arg2 10))")
}

func (s *SimplifierSuite) Test_simplifyComparison(c *C) {
	sx := Simplify(tree.Comparison{Left: tree.NumericLiteral{42}, Op: tree.EQL, Right: tree.NumericLiteral{41}})
	c.Assert(tree.ExpressionString(sx), Equals, "false")

	sx = Simplify(tree.Comparison{Left: tree.NumericLiteral{42}, Op: tree.NEQL, Right: tree.NumericLiteral{41}})
	c.Assert(tree.ExpressionString(sx), Equals, "true")

	sx = Simplify(tree.Comparison{Left: tree.NumericLiteral{42}, Op: tree.GT, Right: tree.NumericLiteral{41}})
	c.Assert(tree.ExpressionString(sx), Equals, "true")

	sx = Simplify(tree.Comparison{Left: tree.NumericLiteral{42}, Op: tree.GTE, Right: tree.NumericLiteral{41}})
	c.Assert(tree.ExpressionString(sx), Equals, "true")

	sx = Simplify(tree.Comparison{Left: tree.NumericLiteral{42}, Op: tree.LT, Right: tree.NumericLiteral{41}})
	c.Assert(tree.ExpressionString(sx), Equals, "false")

	sx = Simplify(tree.Comparison{Left: tree.NumericLiteral{42}, Op: tree.LTE, Right: tree.NumericLiteral{41}})
	c.Assert(tree.ExpressionString(sx), Equals, "false")

	sx = Simplify(tree.Comparison{Left: tree.NumericLiteral{3}, Op: tree.BIT, Right: tree.NumericLiteral{2}})
	c.Assert(tree.ExpressionString(sx), Equals, "true")
}

func (s *SimplifierSuite) Test_simplifyOr(c *C) {
	sx := Simplify(tree.Or{
		tree.Comparison{
			Left:  tree.Arithmetic{Op: tree.LSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
			Op:    tree.EQL,
			Right: tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		},
		tree.Comparison{
			Left:  tree.Argument{2},
			Op:    tree.EQL,
			Right: tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		},
	})
	c.Assert(tree.ExpressionString(sx), Equals, "(eq arg2 10)")

	sx = Simplify(tree.Or{
		tree.Comparison{
			Left:  tree.Arithmetic{Op: tree.LSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
			Op:    tree.EQL,
			Right: tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{42}, Right: tree.Variable{"foo"}},
		},
		tree.Comparison{
			Left:  tree.Argument{2},
			Op:    tree.EQL,
			Right: tree.Arithmetic{Op: tree.RSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		},
	})
	c.Assert(tree.ExpressionString(sx), Equals, "(or (eq 168 (rsh 42 foo)) (eq arg2 10))")
}

func (s *SimplifierSuite) Test_Argument(c *C) {
	sx := Simplify(tree.Argument{3})
	c.Assert(tree.ExpressionString(sx), Equals, "arg3")
}

func (s *SimplifierSuite) Test_simplifyBinaryNegation(c *C) {
	sx := Simplify(tree.BinaryNegation{tree.NumericLiteral{42}})
	c.Assert(tree.ExpressionString(sx), Equals, "4294967253")
}

func (s *SimplifierSuite) Test_simplifyBooleanLiteral(c *C) {
	sx := Simplify(tree.BooleanLiteral{true})
	c.Assert(tree.ExpressionString(sx), Equals, "true")
}

func (s *SimplifierSuite) Test_simplifyInclusion(c *C) {
	sx := Simplify(tree.Inclusion{
		Positive: true,
		Left:     tree.BinaryNegation{tree.NumericLiteral{42}},
		Rights: []tree.Numeric{
			tree.NumericLiteral{42},
			tree.Arithmetic{Op: tree.LSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		}})
	c.Assert(tree.ExpressionString(sx), Equals, "false")

	sx = Simplify(tree.Inclusion{
		Positive: true,
		Left:     tree.BinaryNegation{tree.NumericLiteral{42}},
		Rights: []tree.Numeric{
			tree.Argument{0},
			tree.Arithmetic{Op: tree.LSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		}})
	c.Assert(tree.ExpressionString(sx), Equals, "(in 4294967253 arg0)")

	sx = Simplify(tree.Inclusion{
		Positive: true,
		Left:     tree.Argument{0},
		Rights: []tree.Numeric{
			tree.NumericLiteral{42},
			tree.Arithmetic{Op: tree.LSH, Left: tree.NumericLiteral{42}, Right: tree.NumericLiteral{2}},
		}})
	c.Assert(tree.ExpressionString(sx), Equals, "(in arg0 42 168)")
}

func (s *SimplifierSuite) Test_simplifyNegation(c *C) {
	sx := Simplify(tree.Negation{tree.BooleanLiteral{true}})
	c.Assert(tree.ExpressionString(sx), Equals, "false")
}

func (s *SimplifierSuite) Test_simplifyNumericLiteral(c *C) {
	sx := Simplify(tree.NumericLiteral{42})
	c.Assert(tree.ExpressionString(sx), Equals, "42")
}

func (s *SimplifierSuite) Test_Variable(c *C) {
	sx := Simplify(tree.Variable{"foo"})
	c.Assert(tree.ExpressionString(sx), Equals, "foo")
}