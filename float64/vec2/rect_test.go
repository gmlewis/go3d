package vec2

import "testing"

func TestContains(t *testing.T) {
	a := Rect{T{0, 0}, T{2, 2}}
	b := Rect{T{1, 1}, T{3, 3}}
	if got, want := a.Contains(&b), false; got != want {
		t.Errorf("failed Contains: got %v, want %v", got, want)
	}
	if got, want := b.Contains(&a), false; got != want {
		t.Errorf("failed Contains: got %v, want %v", got, want)
	}
	c := Rect{T{0, 0}, T{1, 1}}
	if got, want := c.Contains(&a), false; got != want {
		t.Errorf("failed Contains: got %v, want %v", got, want)
	}
	if got, want := a.Contains(&c), true; got != want {
		t.Errorf("failed Contains: got %v, want %v", got, want)
	}
}

func TestJoined(t *testing.T) {
	a := Rect{T{0, 0}, T{2, 2}}
	b := Rect{T{1, 1}, T{3, 3}}
	if got, want := Joined(&a, &b), (Rect{T{0, 0}, T{3, 3}}); got != want {
		t.Errorf("failed Joined: got %v, want %v", got, want)
	}
}
