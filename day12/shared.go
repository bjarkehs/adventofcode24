package day12

import (
	"adventofcode2024-go/helpers"
	"container/list"
	"sort"
)

type Coordinate struct {
	x int
	y int
}

type Region struct {
	id          string
	coordinates map[Coordinate]bool
}

func findRegionsInMap(areaMap [][]string) []Region {
	regions := make([]Region, 0)
	maxX := len(areaMap[0])
	maxY := len(areaMap)
	processedCoordinates := make(map[Coordinate]bool)
	for y, row := range areaMap {
		for x, cell := range row {
			coordinate := Coordinate{x, y}
			if _, ok := processedCoordinates[coordinate]; ok {
				continue
			}
			processedCoordinates[coordinate] = true
			coordinates := make(map[Coordinate]bool)
			coordinates[coordinate] = true
			neighbors := adjacentCoordinates(coordinate)
			queue := list.New()
			for _, neighbor := range neighbors {
				if !isValidCoordinate(neighbor, maxX, maxY) {
					continue
				}

				if _, ok := processedCoordinates[neighbor]; ok {
					continue
				}

				queue.PushBack(neighbor)
			}

			processedNeighbors := make(map[Coordinate]bool)

			for queue.Len() > 0 {
				front := queue.Front()
				queue.Remove(front)
				queuedCoordinate := front.Value.(Coordinate)

				processedNeighbors[queuedCoordinate] = true

				if areaMap[queuedCoordinate.y][queuedCoordinate.x] != cell {
					continue
				}

				if _, ok := processedCoordinates[queuedCoordinate]; ok {
					continue
				}

				nestedNeighbors := adjacentCoordinates(queuedCoordinate)

				coordinates[queuedCoordinate] = true
				processedCoordinates[queuedCoordinate] = true

				for _, neighbor := range nestedNeighbors {
					if !isValidCoordinate(neighbor, maxX, maxY) {
						continue
					}

					if _, ok := processedCoordinates[neighbor]; ok {
						continue
					}

					if _, ok := processedNeighbors[neighbor]; ok {
						continue
					}

					queue.PushBack(neighbor)
				}
			}

			regions = append(regions, Region{id: cell, coordinates: coordinates})
		}
	}

	return regions
}

func findPerimeterSize(region Region) int {
	perimeter := 0

	for coordinate := range region.coordinates {
		neighbors := adjacentCoordinates(coordinate)
		for _, neighbor := range neighbors {
			if _, ok := region.coordinates[neighbor]; !ok {
				perimeter++
			}
		}
	}

	return perimeter
}

type SidePosition struct {
	position  int
	direction int
}

type Neighbor struct {
	coordinate Coordinate
	direction  int
}

func findSidesCount(region Region) int {
	sidesMap := make(map[SidePosition][]Coordinate)

	for coordinate := range region.coordinates {
		neighbors := findNeighbors(coordinate)
		for _, neighbor := range neighbors {
			if _, ok := region.coordinates[neighbor.coordinate]; !ok {
				sidePosition := sidePositionForNeighbor(neighbor)
				if _, ok := sidesMap[sidePosition]; !ok {
					sidesMap[sidePosition] = make([]Coordinate, 0)
				}
				sidesMap[sidePosition] = append(sidesMap[sidePosition], coordinate)
			}
		}
	}

	sides := 0
	for side, sideCoordinates := range sidesMap {
		sides++
		if len(sideCoordinates) == 1 {
			continue
		}

		points := make([]int, len(sideCoordinates))
		for index, coordinate := range sideCoordinates {
			if side.direction%2 == 0 {
				points[index] = coordinate.x
			} else {
				points[index] = coordinate.y
			}
		}

		sort.Ints(points)

		for i := 1; i < len(points); i++ {
			if points[i]-points[i-1] > 1 {
				sides++
			}
		}
	}

	return sides
}

func isCoordinateAdjacentToOtherCoordinate(coordinate Coordinate, otherCoordinate Coordinate, direction int) bool {
	if direction%2 == 0 {
		return helpers.AbsInt(coordinate.x-otherCoordinate.x) <= 1
	}
	return helpers.AbsInt(coordinate.y-otherCoordinate.y) <= 1
}

func sidePositionForNeighbor(neighbor Neighbor) SidePosition {
	position := 0
	switch neighbor.direction {
	case 0:
		position = neighbor.coordinate.y + 1
	case 1:
		position = neighbor.coordinate.x - 1
	case 2:
		position = neighbor.coordinate.y - 1
	case 3:
		position = neighbor.coordinate.x + 1
	}

	return SidePosition{position: position, direction: neighbor.direction}
}

func findNeighbors(coordinate Coordinate) []Neighbor {
	return []Neighbor{
		{Coordinate{coordinate.x, coordinate.y - 1}, 0},
		{Coordinate{coordinate.x + 1, coordinate.y}, 1},
		{Coordinate{coordinate.x, coordinate.y + 1}, 2},
		{Coordinate{coordinate.x - 1, coordinate.y}, 3},
	}
}

func adjacentCoordinates(coordinate Coordinate) []Coordinate {
	return []Coordinate{
		{coordinate.x + 1, coordinate.y},
		{coordinate.x - 1, coordinate.y},
		{coordinate.x, coordinate.y + 1},
		{coordinate.x, coordinate.y - 1},
	}
}

func isValidCoordinate(coordinate Coordinate, maxX int, maxY int) bool {
	return coordinate.x >= 0 && coordinate.x < maxX && coordinate.y >= 0 && coordinate.y < maxY
}
