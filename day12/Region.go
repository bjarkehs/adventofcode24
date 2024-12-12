package day12

import "sort"

type Region struct {
	id          string
	coordinates map[Coordinate]bool
}

func (region Region) perimeterSize() int {
	perimeter := 0

	for coordinate := range region.coordinates {
		neighbors := coordinate.adjacentCoordinates()
		for _, neighbor := range neighbors {
			if _, ok := region.coordinates[neighbor]; !ok {
				perimeter++
			}
		}
	}

	return perimeter
}

func (region Region) amountOfSides() int {
	sidesMap := make(map[SidePosition][]int)

	for coordinate := range region.coordinates {
		neighbors := coordinate.adjacentCoordinates()
		for direction, neighbor := range neighbors {
			if _, ok := region.coordinates[neighbor]; ok {
				continue
			}

			sidePosition := sidePositionForCoordinate(coordinate, direction)

			if _, ok := sidesMap[sidePosition]; !ok {
				sidesMap[sidePosition] = make([]int, 0)
			}

			sidesMap[sidePosition] = append(sidesMap[sidePosition], sidePosition.pointForCoordinate(coordinate))
		}
	}

	sides := 0
	for _, sideCoordinates := range sidesMap {
		sides++
		if len(sideCoordinates) == 1 {
			continue
		}

		sort.Ints(sideCoordinates)

		for i := 1; i < len(sideCoordinates); i++ {
			if sideCoordinates[i]-sideCoordinates[i-1] > 1 {
				sides++
			}
		}
	}

	return sides
}
