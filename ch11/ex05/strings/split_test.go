package strings

import (
	"reflect"
	"strings"
	"testing"
)

func TestSplit(t *testing.T) {

	var tests = []struct {
		in   string
		sep  string
		want []string
	}{
		{"a:b:c", ":", []string{"a", "b", "c"}},
		{"a:b:", ":", []string{"a", "b", ""}},
	}
	for _, test := range tests {
		got := strings.Split(test.in, test.sep)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s wants %v, but %v", test.in, test.want, got)
		}
	}
}
