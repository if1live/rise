package main

import "testing"

func TestMakeSafeURI(t *testing.T) {
	cases := []struct {
		uri      string
		expected string
	}{
		{"http://google.com", "http://google.com"},
		{"https://google.com", "https://google.com"},
		{"google.com", "http://google.com"},
	}
	for _, c := range cases {
		v := makeSafeURI(c.uri)
		if v != c.expected {
			t.Error(
				"For makeSafeURI",
				"expected", c.expected,
				"got", v,
			)
		}
	}
}
