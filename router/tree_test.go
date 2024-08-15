package router

import (
	"testing"
)

func TestTreePathFinding(t *testing.T) {
	tree := NewTree[func()]()
	values := []string{
		"/this/is/a/test",
		"/this/is/a/possible/test",
		"/this/path/is/not/here",
		"/this/is/a/potential/path",
		"/this/:is/a/test",
		"/this/is/a/:variable/test",
		"/this/is/a/test/with/more/segments",
		"/this/is/:a/test/with/:variable",
		"/this/might/:have/:more/:than/one/:variable",
		"/path/that/does-:not/have/leading",
		"/path/that/test:not/have/leading",
		"/:variable",
	}

	for _, v := range values {
		tree.Insert(v, nil)
	}

	tests := []struct {
		path     string
		expected bool
	}{
		{"/this/is/a/test", true},                         // exact match
		{"/this/is/a/possible/test", true},                // exact match
		{"/this/path/is/not/here", true},                  // exact match
		{"/this/is/a/potential/path", true},               // exact match
		{"/variableValue", true},                          // variable match
		{"/this/123/a/test", true},                        // path with variable segment
		{"/this/is/a/123/test", true},                     // path with variable at the end
		{"/this/is/a/test/with/more/segments", true},      // exact match with longer path
		{"/this/is/a/test/with/variableValue", true},      // variable match in the middle
		{"/notfound", false},                              // non-existent path
		{"/this/is/not/found", false},                     // non-existent nested path
		{"/this/is/a/test/extra", false},                  // partially matching path
		{"/this/is/a/possible/test/more", false},          // partially matching longer path
		{"/this/might/1/2/3/one/4", true},                 // nested variables
		{"/path/that/does-1/have/leading", true},          // partial variable paths
		{"/path/that/test1/have/leading", true},           // partial variable path mixed
		{"/path/that/does_not_exist/have/leading", false}, // partial variable check
	}

	for _, test := range tests {
		found := tree.Search(test.path) != nil
		if found != test.expected {
			t.Errorf("Path %s was found: %v, expected: %v", test.path, found, test.expected)
		}
	}
}
