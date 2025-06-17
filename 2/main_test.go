package main

import (
	"testing"
)

func twoSmallest(a, b, c int) (int, int) {
    if a > b {
        a, b = b, a
    }

    if b > c {
        b, c = c, b
    }

    if a > b {
        a, b = b, a
    }

    return a, b
}

func getWrapRibbon(present Rectangle) int {
    a, b := twoSmallest(present.Length, present.Width, present.Height)

    return 2 * (a + b)
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{Width: 3, Height: 4, Length: 2}
	got := rectangle.Area()
	want := 52

	assetsCorrectMessage(t, got, want)
}

func TestFoundSmallAreaSide(t *testing.T) {
	rectangle := Rectangle{Width: 1, Height: 1, Length: 10}
	got := rectangle.FoundSmallAreaSide()
	want := 1

	assetsCorrectMessage(t, got, want)
}

func TestTotalAreaNeeded(t *testing.T) {
	rectangle := Rectangle{Width: 10, Height: 1, Length: 1}
	got := rectangle.TotalAreaNeeded()
	want := 43

	assetsCorrectMessage(t, got, want)
}

func TestRibbon(t *testing.T) {
	t.Run("Testing for found the 2 smaller sides and calc your perimeter", func(t *testing.T) {
		rectangle := Rectangle{Width: 2, Height: 3, Length: 4}
		got := rectangle.SmallestPerimeterOfFace()
		want := 10

		assetsCorrectMessage(t, got, want)
	})
	t.Run("Testing the Bow size", func(t *testing.T) {
		rectangle := Rectangle{Width: 2, Height: 3, Length: 4}
		got := rectangle.GiftVolume()
		want := 24

		assetsCorrectMessage(t, got, want)

	})
	t.Run("Testing the total of the ribbon that it will be use", func(t *testing.T) {
		rectangle := Rectangle{Width: 2, Height: 3, Length: 4}
		got := rectangle.TotalRibbonLength()
		want := 34

		assetsCorrectMessage(t, got, want)

	})
}

func BenchmarkSmallPerimeter(b *testing.B) {
	testCases := []Rectangle{
		{2, 3, 4},
		{1, 1, 10},
		{5, 5, 5},
		{20, 15, 10},
		{100, 50, 25},
	}

	for i := 0; i < b.N; i++ {
		for _, present := range testCases {
			present.SmallestPerimeterOfFace()
		}
	}
}
func BenchmarkSmallPerimeterPart2(b *testing.B) {
	testCases := []Rectangle{
		{2, 3, 4},
		{1, 1, 10},
		{5, 5, 5},
		{20, 15, 10},
		{100, 50, 25},
	}

	for i := 0; i < b.N; i++ {
		for _, present := range testCases {
			getWrapRibbon(present)
		}
	}
}

func assetsCorrectMessage(t testing.TB, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("GOT %d, WANT %d", got, want)
	}

}
