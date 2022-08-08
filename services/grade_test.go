package services_test

import (
	"gobasic/services"
	"testing"
)

func TestCheckGrade(t *testing.T) {

	// subtest
	t.Run("A", func(t *testing.T) {
		grade := services.CheckGrade(85)
		expected := "A"
		if grade != "A" {
			t.Errorf("Get %v expected %v", grade, expected)
		}
	})

	t.Run("B", func(t *testing.T) {
		grade := services.CheckGrade(70)
		expected := "B"
		if grade != "B" {
			t.Errorf("Get %v expected %v", grade, expected)
		}
	})
}
