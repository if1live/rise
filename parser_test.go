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

func TestMakeSafeImageURI(t *testing.T) {
	cases := []struct {
		imguri   string
		uri      string
		expected string
	}{
		{"http://foo.com/bar.png", "http://foo.com/", "http://foo.com/bar.png"},
		{"/bar.png", "http://foo.com/", "http://foo.com/bar.png"},
		{"./bar.png", "http://foo.com/sample/", "http://foo.com/sample/bar.png"},
		{"../bar.png", "http://foo.com/sample/", "http://foo.com/bar.png"},
	}
	for _, c := range cases {
		v := makeSafeImageURI(c.imguri, c.uri)
		if v != c.expected {
			t.Error(
				"For makeSafeImageURI",
				"expected", c.expected,
				"got", v,
			)
		}
	}
}
