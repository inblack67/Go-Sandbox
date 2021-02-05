package main

import (
	"testing"
)

func TestAdd(t *testing.T){
	if (add(2,2) != 4){
		t.Error("cant even add?")
	}
}

func TestAdds(t *testing.T){
	tests := []struct{
		input1 int
		input2 int
		expected int
	}{
		{2,4,6},
		{1,4,5},
		{3,4,7},
		{2,4,6},
	}
	for _, test := range tests{
		if output := add(test.input1, test.input2); output != test.expected{
			t.Error("what is this?", test.input1, test.input2, test.expected)
		}
	}
}