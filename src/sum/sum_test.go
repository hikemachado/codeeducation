package main

import "testing"

func TestSum(t *testing.T) {

	//var in1, in2 int

	in1 := 5
	in2 := 5

	result := sum(in1, in2)

	if result != 10 {
		t.Errorf("failed result ")
	}
}
