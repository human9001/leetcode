package main

import (
	"reflect"
	"testing"
)

type addTest struct {
	arg1 *ListNode
	arg2 *ListNode
	want *ListNode
}

var arg1 = ListNode{
	2, &ListNode{4, &ListNode{3, nil}},
}
var addTests = []addTest{
	{&ListNode{2, &ListNode{4, &ListNode{3, nil}}}, &ListNode{5, &ListNode{6, &ListNode{4, nil}}}, &ListNode{7, &ListNode{0, &ListNode{8, nil}}}},
}

func TestAddTwoNumbers(t *testing.T) {
	for _, test := range addTests {
		if output := addTwoNumbers(test.arg1, test.arg2); !reflect.DeepEqual(output, test.want) {
			t.Errorf("Output %v not equal to expected %v", *output, *test.want)

		}

	}
}

func BenchmarkAddTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		addTwoNumbers(addTests[0].arg1, addTests[0].arg2)
	}
}
