package services_test

import (
	"gobasic/services"
	"testing"
)

func TestCheckGrade(t *testing.T) {

	// struct temp
	type testCase struct {
		name     string
		score    int
		expected string
	}

	cases := []testCase{
		{name: "A", score: 100, expected: "A"},
		{name: "A", score: 80, expected: "A"},
		{name: "B", score: 70, expected: "B"},
		{name: "C", score: 60, expected: "C"},
		{name: "D", score: 50, expected: "D"},
		{name: "F", score: 20, expected: "F"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			grade := services.CheckGrade(c.score)
			expected := c.expected
			if grade != c.expected {
				t.Errorf("Get %v expected %v", grade, expected)
			}
		})
	}
}
