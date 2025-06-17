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
	Length int
}

func (r *Rectangle) Area() int {
	return 2*r.Height*r.Length + 2*r.Length*r.Width + 2*r.Width*r.Height
}

func (r *Rectangle) FoundSmallAreaSide() int {
	sideArea1 := r.Height * r.Width
	sideArea2 := r.Height * r.Length
	sideArea3 := r.Length * r.Width

	return min(sideArea1, sideArea2, sideArea3)
}

func (r *Rectangle) TotalAreaNeeded() int {
	return r.Area() + r.FoundSmallAreaSide()
}

func getRectanglesByFile(filePath string) []Rectangle {
	allRectangles := []Rectangle{}

	fileContent, _ := os.ReadFile(filePath)
	for line := range strings.SplitSeq(string(fileContent), "\n") {
		if len(line) > 0 {
			dimensions := strings.Split(line, "x")

			length, _ := strconv.Atoi(dimensions[0])
			width, _ := strconv.Atoi(dimensions[1])
			height, _ := strconv.Atoi(dimensions[2])
			allRectangles = append(allRectangles, Rectangle{Length: length, Width: width, Height: height})
		}
	}
	return allRectangles
}

func (r *Rectangle) SmallestPerimeterOfFace() int{
	perimeterHeightWidth := 2*r.Height + 2*r.Width
	perimeterHeightLength := 2*r.Height + 2*r.Length
	perimeterLengthWidth := 2*r.Length + 2*r.Width

	return min(perimeterHeightWidth,perimeterHeightLength,perimeterLengthWidth)
}

func (r *Rectangle) GiftVolume() int {
	return r.Height * r.Width * r.Length
}

func (r *Rectangle) TotalRibbonLength() int {
	return r.GiftVolume() + r.SmallestPerimeterOfFace()
}


func main() {
	rectangles := getRectanglesByFile("./input.txt")

	var totalWrappingPaperArea int
	var totalRibbonRequired int
	for _, rect := range rectangles {
		totalWrappingPaperArea += rect.TotalAreaNeeded()
		totalRibbonRequired += rect.TotalRibbonLength()
	}
	fmt.Printf("Total area of all rectangles: %d\n", totalWrappingPaperArea)
	fmt.Printf("Total of the ribbon that will be use: %d\n", totalRibbonRequired)
}
