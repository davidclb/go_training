package hello

import (
	"testing"
)

func TestHello(t *testing.T) {

	want := ("Hello, david, coulibaly!")
	got := Say([]string{"david", "coulibaly"})
	if got != want {
		t.Errorf("Say(\"david\") = %s; want (\"Hello, david, coulibaly!\")", got)
	}
}
