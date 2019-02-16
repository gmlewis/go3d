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

func TestIntersects(t *testing.T) {
	a := Rect{T{-0.07499999999999983, 1.3818309886183793}, T{15.031830988618381, 1.5318309886183792}}
	b := Rect{T{2.306134154342999, -11.530670771714988}, T{25.367475697772974, 11.530670771714988}}
	if got, want := a.Intersects(&b), true; got != want {
		t.Errorf("failed Intersects: got %v, want %v", got, want)
	}
	if got, want := b.Intersects(&a), true; got != want {
		t.Errorf("failed Intersects: got %v, want %v", got, want)
	}

	a = Rect{T{0, 0}, T{1, 1}}
	b = Rect{T{0, 2}, T{0, 3}}
	if got, want := a.Intersects(&b), false; got != want {
		t.Errorf("failed Intersects: got %v, want %v", got, want)
	}
	if got, want := b.Intersects(&a), false; got != want {
		t.Errorf("failed Intersects: got %v, want %v", got, want)
	}
}
