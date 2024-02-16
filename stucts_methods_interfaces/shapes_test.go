package stuctsmethodsinterfaces

import "testing"

func TestPermiter(t *testing.T) {
	t.Run("rectangle permiter", func(t *testing.T) {
		rectangle := Rectange{Height: 10.0, Width: 20.0}
		want := 60.0
		checkPermiter(t, rectangle, want)
	})

	t.Run("Circle permiter", func(t *testing.T) {
		circle := Circle{Radius: 5.0}
		want := 31.41592653589793
		checkPermiter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	t.Run("test rectangle", func(t *testing.T) {
		rectangle := Rectange{Height: 2.0, Width: 25.0}
		want := 50.0
		checkArea(t, rectangle, want)
	})

	t.Run("test circle", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		want := 314.1592653589793
		checkArea(t, circle, want)
	})
}

func checkArea(t testing.TB, s Shape, want float64) {
	got := s.Area()
	if want != got {
		t.Errorf("expected %g but got %g", want, got)
	}
}

func checkPermiter(t testing.TB, s Shape, want float64) {
	got := s.Permiter()
	if want != got {
		t.Errorf("expected %g but got %g", want, got)
	}
}
