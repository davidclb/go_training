package exos

import (
	"reflect"
	"testing"
)

func TestSafeAppend(t *testing.T) {

	s := make([]int, 2, 10)
	s[0] = 1
	s[1] = 2

	SafeAppend(s, 99)

	got := s[:cap(s)]                           // les 10 cases
	want := []int{1, 2, 0, 0, 0, 0, 0, 0, 0, 0} // backing intact = que des zéros après [1 2]
	if !reflect.DeepEqual(got, want) {
		t.Errorf("backing de s modifié ! got %v; want %v", got, want)
	}

}
