package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {

	t.Run("expecting 5 repeats", func(t *testing.T) {
		got := Repeat("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("expecting 10 repeats", func(t *testing.T) {
		got := Repeat("b", 10)
		expected := "bbbbbbbbbb"

		if got != expected {
			t.Errorf("got %q expected %q ", got, expected)
		}
	})

}

func TestRepeatV2(t *testing.T) {

	t.Run("expecting 5 repeats", func(t *testing.T) {
		got := RepeatV2("a", 5)
		expected := "aaaaa"

		if got != expected {
			t.Errorf("got %q expected %q", got, expected)
		}
	})

	t.Run("expecting 10 repeats", func(t *testing.T) {
		got := RepeatV2("b", 10)
		expected := "bbbbbbbbbb"

		if got != expected {
			t.Errorf("got %q expected %q ", got, expected)
		}
	})

}

func BenchamrkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", b.N)
	}
}

func ExampleRepeat() {
	repeat := Repeat("a", 5)
	fmt.Println(repeat)
}
