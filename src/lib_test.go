package main

import "testing"

func TestSqrtSum(t *testing.T) {
	res := sqrtSum(0.1,5)
	if(res != 6.246264369941972){
		t.Errorf("Expected 6.246264 but got %v", res)
	}
}