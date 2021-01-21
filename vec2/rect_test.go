package vec2

import "testing"

func TestIntersects(t *testing.T) {
	tests := []struct {
		name string
		a    *Rect
		b    *Rect
		want bool
	}{
		{
			name: "a touches top left of b",
			a:    &Rect{Min: T{0, 0}, Max: T{1, 1}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		{
			name: "a touches top middle of b",
			a:    &Rect{Min: T{1, 0}, Max: T{2, 1}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		{
			name: "a touches top right of b",
			a:    &Rect{Min: T{2, 0}, Max: T{3, 1}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		{
			name: "a touches center left of b",
			a:    &Rect{Min: T{0, 1}, Max: T{1, 2}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		{
			name: "a on top of b",
			a:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		{
			name: "a touches center right of b",
			a:    &Rect{Min: T{2, 1}, Max: T{3, 2}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		{
			name: "a touches bottom left of b",
			a:    &Rect{Min: T{0, 2}, Max: T{1, 3}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		{
			name: "a touches bottom middle of b",
			a:    &Rect{Min: T{1, 2}, Max: T{2, 3}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		{
			name: "a touches bottom right of b",
			a:    &Rect{Min: T{2, 2}, Max: T{3, 3}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: true,
		},
		// does not touch
		{
			name: "a does not touch top left of b",
			a:    &Rect{Min: T{0, -1}, Max: T{1, 0}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch top left of b - version 2",
			a:    &Rect{Min: T{-1, 0}, Max: T{0, 1}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch top middle of b",
			a:    &Rect{Min: T{1, -1}, Max: T{2, 0}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch top right of b",
			a:    &Rect{Min: T{2, -1}, Max: T{3, 0}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch top right of b - version 2",
			a:    &Rect{Min: T{3, 0}, Max: T{4, 1}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch center left of b",
			a:    &Rect{Min: T{-1, 1}, Max: T{0, 2}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch center right of b",
			a:    &Rect{Min: T{3, 1}, Max: T{4, 2}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch bottom left of b",
			a:    &Rect{Min: T{0, 3}, Max: T{1, 4}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch bottom left of b - version 2",
			a:    &Rect{Min: T{-1, 2}, Max: T{0, 3}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch bottom middle of b",
			a:    &Rect{Min: T{1, 3}, Max: T{2, 4}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch bottom right of b",
			a:    &Rect{Min: T{2, 3}, Max: T{3, 4}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
		{
			name: "a does not touch bottom right of b - version 2",
			a:    &Rect{Min: T{3, 2}, Max: T{4, 3}},
			b:    &Rect{Min: T{1, 1}, Max: T{2, 2}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.a.Intersects(tt.b)

			if got != tt.want {
				t.Errorf("got = %v, want %v", got, tt.want)
			}
		})
	}
}
