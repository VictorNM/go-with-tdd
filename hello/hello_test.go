package hello

import "testing"

func TestHello(t *testing.T)  {
	expected := "Hello World"
	actual := Hello()

	if expected != actual {
		// For tests %q is very useful as it wraps your values in double quotes.
		t.Errorf("Expect 'Hello' return: %q but it return %q\n", expected, actual)
	}
}