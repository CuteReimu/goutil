package slices

import "testing"

func TestRangeFunctionWithPositiveNumbers(t *testing.T) {
	for i := range Range(1, 5) {
		if i < 1 || i > 4 {
			t.Fail()
		}
	}
}

func TestRangeFunctionWithNegativeNumbers(t *testing.T) {
	for i := range Range(-5, 0) {
		if i < -5 || i >= 0 {
			t.Fail()
		}
	}
}

func TestRangeFunctionWithZero(t *testing.T) {
	for range Range(0, 0) {
		t.Fail() // Should not be called
	}
}

func TestRangeFunctionWithSameStartAndEnd(t *testing.T) {
	for range Range(10, 10) {
		t.Fail() // Should not be called
	}
}

func TestRangeFunctionWithStartGreaterThanEnd(t *testing.T) {
	for range Range(5, 1) {
		t.Fail() // Should not be called
	}
}

func TestProgressionFunctionWithPositiveNumbers(t *testing.T) {
	for i := range Progression[int](1, 5, 1) {
		if i < 1 || i > 4 {
			t.Fail()
		}
	}
}

func TestProgressionFunctionWithNegativeNumbers(t *testing.T) {
	for i := range Progression[int](-5, 0, 1) {
		if i < -5 || i >= 0 {
			t.Fail()
		}
	}
}

func TestProgressionFunctionWithZero(t *testing.T) {
	for range Progression[int](0, 0, 1) {
		t.Fail() // Should not be called
	}
}

func TestProgressionFunctionWithSameStartAndEnd(t *testing.T) {
	for range Progression[int](10, 10, 1) {
		t.Fail() // Should not be called
	}
}

func TestProgressionFunctionWithStartGreaterThanEnd(t *testing.T) {
	for range Progression[int](5, 1, 1) {
		t.Fail() // Should not be called
	}
}

func TestProgressionFunctionWithStepGreaterThanRange(t *testing.T) {
	for i := range Progression[int](1, 5, 10) {
		if i != 1 {
			t.Fail()
		}
	}
}
