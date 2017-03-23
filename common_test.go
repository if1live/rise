package main

import "testing"

func TestApplyFilter(t *testing.T) {
	cases := []struct {
		link     string
		mode     int
		expected string
	}{
		{
			"http://google.com/helloworld.jpg/w650",
			1,
			"http://google.com/helloworld.jpg",
		},
		{
			"http://google.com/helloworld.jpg?w650",
			2,
			"http://google.com/helloworld.jpg",
		},
	}
	for _, c := range cases {
		actual := ApplyFilter(c.link, c.mode)
		if actual != c.expected {
			t.Error(
				"For ApplyFilter",
				"expected", c.expected,
				"got", actual,
			)
		}
	}
}

func TestMakeFilename(t *testing.T) {
	cases := []struct {
		idx      int
		basename string
		expected string
	}{
		{1, "foo.png", "001-foo.png"},
		{12, "foo.png", "012-foo.png"},
		{123, "foo.png", "123-foo.png"},
		{1234, "foo.png", "1234-foo.png"},
	}
	for _, c := range cases {
		v := MakeFilename(c.basename, c.idx)
		if v != c.expected {
			t.Error(
				"For makeFilename",
				"expected", c.expected,
				"got", v,
			)
		}
	}
}
