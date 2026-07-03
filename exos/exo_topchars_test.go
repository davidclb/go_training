package exos

import (
	"reflect"
	"testing"
)

func TestTopChars(t *testing.T) {
	tests := []struct {
		name string
		text string
		n    int
		want []CharCount
	}{
		{
			name: "ascii basique",
			text: "aabbbc",
			n:    2,
			want: []CharCount{{'b', 3}, {'a', 2}},
		},
		{
			name: "chinois",
			text: "好好学习",
			n:    2,
			want: []CharCount{{'好', 2}, {'习', 1}},
		},
		{
			name: "depart par rune sur egalite",
			text: "cba",
			n:    3,
			want: []CharCount{{'a', 1}, {'b', 1}, {'c', 1}},
		},
		{
			name: "n plus grand que le nombre de caracteres",
			text: "xy",
			n:    10,
			want: []CharCount{{'x', 1}, {'y', 1}},
		},
		{
			name: "espaces ignores",
			text: "a a b",
			n:    3,
			want: []CharCount{{'a', 2}, {'b', 1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopChars(tt.text, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TopChars(%q, %d) = %v; want %v", tt.text, tt.n, got, tt.want)
			}
		})
	}
}
