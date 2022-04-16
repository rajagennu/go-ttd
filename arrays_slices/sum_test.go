package arrslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sending for sum", func(t *testing.T) {
		arr := []int{1, 2, 3, 4, 5}
		expected := 15
		got := Sum(arr)

		if got != expected {
			t.Errorf("Got %d , expected %d for sum %v", got, expected, arr)
		}
	})

}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 3, 4}, []int{4, 5, 6})
	expected := []int{8, 15}

	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Got %v expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("sending non empty list of slices", func(t *testing.T) {
		got := []int{5, 15, 25}
		expected := SumAllTails([]int{5, 5}, []int{5, 5, 5, 5}, []int{5, 5, 5, 5, 5, 5})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("Got %v expected %v ", got, expected)
		}
	})

	t.Run("Sending one or more slices empty, expecting zero at those", func(t *testing.T) {
		got := []int{0, 10, 0}
		expected := SumAllTails([]int{1, 0}, []int{5, 5, 5}, []int{})

		if !reflect.DeepEqual(got, expected) {
			t.Errorf("got %v expected %v ", got, expected)
		}
	})
}
