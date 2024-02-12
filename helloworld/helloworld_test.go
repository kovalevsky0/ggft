package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	got := HelloWorld()
	want := "Hello, World!"

	if got != want {
		t.Errorf("Failed! Expected %q but received %q", want, got)
	}
}
