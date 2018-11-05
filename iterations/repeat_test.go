package iterations

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 3)
	expected := "aaa"

	if repeated != expected {
		t.Errorf("got '%s' and '%s' expected", repeated, expected)
	}
}
