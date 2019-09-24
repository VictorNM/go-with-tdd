package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T)  {
	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		expected := 15
		actual := Sum(numbers)

		if expected != actual {
			t.Errorf("Expected '%d' but actual is '%d', given '%v'", expected, actual, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	expected := []int{3, 9}
	actual := SumAll([]int{1, 2}, []int{2, 7})

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected '%v' but actual is '%v'", expected, actual)
	}
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, expected, actual []int) {
		t.Helper()
		if !reflect.DeepEqual(expected, actual) {
			t.Errorf("Expected '%v' but actual is '%v'", expected, actual)
		}
	}

	t.Run("make the sums of some slice", func(t *testing.T) {
		expected := []int{2, 9}
		actual := SumAllTails([]int{1, 2}, []int{0, 9})

		checkSums(t, expected, actual)
	})

	t.Run("safely sum empty slice", func(t *testing.T) {
		expected := []int{0, 9}
		actual := SumAllTails([]int{}, []int{3, 4, 5})

		checkSums(t, expected, actual)
	})
}