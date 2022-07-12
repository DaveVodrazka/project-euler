package main

import (
	"testing"
)

func TestIsAbundant(t *testing.T) {
	if !isAbundant(12) {
		t.Errorf("Incorrect for 12, should be abundant")
	}
	if isAbundant(16) {
		t.Errorf("Incorrect for 16, should not be abundant")
	}
}

func TestDivSum(t *testing.T) {
	if divSum(1) != 0 {
		t.Errorf("Incorrect for 12, should be %d, but got %d.", 0, divSum(1))
	}
	if divSum(12) != 16 {
		t.Errorf("Incorrect for 12, should be %d, but got %d.", 16, divSum(12))
	}
}
