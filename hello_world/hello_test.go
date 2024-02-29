package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := hello("Chris")
		want := "Hello, Chris"

		assert_correct_message(t, got, want)
	})
	t.Run("say 'Hello World' when empty string passed", func(t *testing.T) {
		got := hello("")
		want := "Hello, World!"

		assert_correct_message(t, got, want)
	})
}

func assert_correct_message(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
