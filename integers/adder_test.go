package integers

import "testing"

func TestAdder(t *testing.T) {
	got := Adder(2, 2)
	expected := 4

	if got != expected {

		t.Errorf("Got '%d' expected '%d' ", got, expected)

	}

}
