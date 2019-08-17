package hello

import "testing"

// Writing a test is just like writing a function, with a few rules
// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// t.Fail() when you want to fail.
// Errorf method on our t which will print out a message and fail the test

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t *testing.T, got, want string) {
		// t.Helper() is needed to tell the test suite that this method is a helper.
		// By doing this when it fails the line number reported will be in our function call rather than inside our test helper
		t.Helper()
		if got != want {
				t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
			got := Hello("Chris", "")
			want := "Hello, Chris"
			assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
			got := Hello("", "")
			want := "Hello, World"
			assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})
}