package hello

import "testing"

func TestHello(t *testing.T)  {
	assertCorrectMessage := func(t *testing.T, expected, actual string) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in
		// our function call rather than inside our test helper
		t.Helper()
		if expected != actual {
			// For tests %q is very useful as it wraps your values in double quotes.
			t.Errorf("Expected %q but actual is %q\n", expected, actual)
		}
	}

	// Sometimes it is useful to group tests around a "thing" and then have sub-tests describing different scenarios
	t.Run("saying hello to people", func(t *testing.T) {
		assertCorrectMessage(t, "Hello, Victor", Hello("Victor", ""))
	})

	t.Run("say 'Hello World' when an empty string is supplied", func(t *testing.T) {
		assertCorrectMessage(t, "Hello, World", Hello("", ""))
	})

	t.Run("say hello in Spanish", func(t *testing.T) {
		assertCorrectMessage(t, "Hola, Victor", Hello("Victor", "Spanish"))
	})
}