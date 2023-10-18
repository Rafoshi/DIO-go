package main

import "testing"

func TestTestShouldSumCorrect(t *testing.T) {
	test := sum(1, 2)
	result := 3.0
	if test != result {
		t.Error("Expected", result, " got ", test)
	}
}
func TestShouldSumIncorrect(t *testing.T) {
	test := sum(1, 2)
	result := 4.0
	if test != result {
		t.Error("Expected", result, " got ", test)
	}
}

func TestShouldSubCorrect(t *testing.T) {
	test := sub(1, 2)
	result := -1.0
	if test != result {
		t.Error("Expected", result, " got ", test)
	}
}
func TestShouldSubIncorrect(t *testing.T) {
	test := sub(1, 2)
	result := 4.0
	if test != result {
		t.Error("Expected", result, " got ", test)
	}
}

func TestShouldMulCorrect(t *testing.T) {
	test := mul(1, 2)
	result := 2.0
	if test != result {
		t.Error("Expected", result, " got ", test)
	}
}
func TestShouldMulIncorrect(t *testing.T) {
	test := mul(1, 2)
	result := 3.0
	if test != result {
		t.Error("Expected", result, " got ", test)
	}
}

func TestShouldDivCorrect(t *testing.T) {
	test := div(1, 2)
	result := 0.5
	if test != result {
		t.Error("Expected", result, " got ", test)
	}
}
func TestShouldDivIncorrect(t *testing.T) {
	test := div(1, 2)
	result := 4.0
	if test != result {
		t.Error("Expected", result, " got ", test)
	}
}
