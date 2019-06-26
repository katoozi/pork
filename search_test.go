package pork

import "testing"

func TestSearchByKeyword(t *testing.T) {
	repos := SearchByKeyword([]string{"one, two", "three"})
	if repos[0] != "myrepo" {
		t.Fail()
	}
}
