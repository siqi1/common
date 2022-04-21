package test

import (
	"reflect"
	"testing"
)

type TestCase[I, E any] struct {
	t      *testing.T
	input  I
	expect E
	fn     func(I) E
}

func (tc *TestCase[I, E]) SetInput(input I) {
	tc.input = input
}

func (tc *TestCase[I, E]) SetExpect(expect E) {
	tc.expect = expect
}

func NewTest[I, E any, T TestCase[I, E]](t *testing.T, fn func(I) E) T {
	return T{fn: fn, t: t}
}

func (tc *TestCase[I, E]) RunTest() {
	if actual := tc.fn(tc.input); !reflect.DeepEqual(actual, tc.expect) {
		tc.t.Error("失败")
	} else {
		tc.t.Log("成功")
	}
}

func (tc *TestCase[I, E]) SetAndRun(input I, expect E) {
	tc.SetInput(input)
	tc.SetExpect(expect)
	tc.RunTest()
}
