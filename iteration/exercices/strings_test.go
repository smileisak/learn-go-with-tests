package exercices

import "testing"

func TestCount(t *testing.T) {
	count := Count("Ismail", "i")
	expected := 1

	if count != expected {
		t.Errorf("got %d but expected %d", count, expected)
	}
}

func TestToUpper(t *testing.T) {
	upper := ToUpper("ismail")
	expected := "ISMAIL"

	if upper != expected {
		t.Errorf("got %q and expected %q.", upper, expected)
	}
}
