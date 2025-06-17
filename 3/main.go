package main

import (
	"fmt"
	"os"
)

const (
	NORTH rune = 94  // ^
	SOUTH rune = 118 // v
	WEST  rune = 60  // <
	EAST  rune = 62  // >
)

type Position struct {
	X int
	Y int
}

type MapSanta struct {
	Position
	Map map[Position]int
}

func (m MapSanta) MapSize() int {
	return len(m.Map)
}

func (m *MapSanta) StartDelivery() {
	m.Map[Position{0, 0}]++
}

func (m *MapSanta) ChangePoint(inputValue string) {
	for _, position := range inputValue {
		switch position {
		case NORTH:
			m.Y++
		case SOUTH:
			m.Y--
		case WEST:
			m.X--
		case EAST:
			m.X++
		default:
		}

		m.Map[m.Position]++

	}
}

func TotalHouseVisitedByBoth(santa, robotSanta MapSanta) int {
	biggestMap := make(map[Position]int)
	smallMap := make(map[Position]int)

	if santa.MapSize() > robotSanta.MapSize() {
		biggestMap = santa.Map
		smallMap = robotSanta.Map
	} else {
		biggestMap = robotSanta.Map
		smallMap = santa.Map
	}
	for key, value := range biggestMap {
		if smallMap[key] == 0 {
			smallMap[key] = value
		}
	}

	return len(smallMap)
}

func main() {
	input, _ := os.ReadFile("./input.txt")
	var santaInput string
	var santaRobotInput string
	for i, position := range input {
		if i%2 == 0 {
			santaInput += string(position)
		} else {
			santaRobotInput += string(position)
		}
	}
	mapSanta := MapSanta{Position{0, 0}, map[Position]int{}}
	mapSantaRobot := MapSanta{Position{0, 0}, map[Position]int{}}

	mapSanta.StartDelivery()
	mapSantaRobot.StartDelivery()

	mapSanta.ChangePoint(string(santaInput))
	mapSantaRobot.ChangePoint(string(santaRobotInput))

	fmt.Printf("Total of house has visited: %d\n", TotalHouseVisitedByBoth(mapSanta, mapSantaRobot))
}
