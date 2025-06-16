package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rectangle struct {
	Width  int
	Height int
	Lenght int
}

func (r *Rectangle) Area() int {
	return 2*r.Height*r.Lenght + 2*r.Lenght*r.Width + 2*r.Width*r.Height
}

func (r *Rectangle) FoundSmallAreaSide() int {
	s1 := r.Height * r.Width
	s2 := r.Height * r.Lenght
	s3 := r.Lenght * r.Width

	return min(s1, s2, s3)
}

func (r *Rectangle) TotalAreaNeeded() int {
	return r.Area() + r.FoundSmallAreaSide()
}

func getRectanglesByFile(uri string) []Rectangle {
	rectangles := []Rectangle{}

	urlsBytes, _ := os.ReadFile(uri)
	for line := range strings.SplitSeq(string(urlsBytes), "\n") {
		if len(line) > 0 {
			rectangle := strings.Split(line, "x")

			lenghtR, _ := strconv.Atoi(rectangle[0])
			widthR, _ := strconv.Atoi(rectangle[1])
			heightR, _ := strconv.Atoi(rectangle[2])
			rectangles = append(rectangles, Rectangle{Lenght: lenghtR, Width: widthR, Height: heightR})

		}
	}
	return rectangles
}

func main() {
	rectangles := getRectanglesByFile("./input.txt")

	var areaAllRectangles int
	for _, rectangle := range rectangles {
		areaAllRectangles += rectangle.TotalAreaNeeded()
	}
	fmt.Println(rectangles)
	fmt.Printf("total area of the sum of all rectangles: %d\n", areaAllRectangles)
}
