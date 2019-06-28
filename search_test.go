package pork

import "testing"

func TestSearchByKeyword(t *testing.T) {
	if err := SearchByKeyword([]string{"one, two", "three"}); err == nil {
		t.Fail()
	}
}
