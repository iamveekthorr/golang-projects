package greetings

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Chris"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := Greet(name)
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Greet("Chris") = %q, %v, want match for %#v, nil`, msg, err, want)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestHelloEmpty(t *testing.T) {
	msg, err := Greet("")
	if msg != "" || err == nil {
		t.Errorf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}
