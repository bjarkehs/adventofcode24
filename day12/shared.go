package day12

import (
	"container/list"
)

func findRegionsInMap(areaMap [][]string) []Region {
	regions := make([]Region, 0)
	maxX := len(areaMap[0])
	maxY := len(areaMap)
	processedCoordinates := make(map[Coordinate]bool)
	for y, row := range areaMap {
		for x, cell := range row {
			coordinate := Coordinate{x, y}

			if processedCoordinates[coordinate] {
				continue
			}

			processedCoordinates[coordinate] = true

			coordinates := make(map[Coordinate]bool)
			coordinates[coordinate] = true

			neighbors := coordinate.adjacentCoordinates()
			queue := list.New()

			for _, neighbor := range neighbors {
				if !neighbor.isValid(maxX, maxY) || processedCoordinates[neighbor] {
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

				if processedCoordinates[queuedCoordinate] {
					continue
				}

				if areaMap[queuedCoordinate.y][queuedCoordinate.x] != cell {
					continue
				}

				nestedNeighbors := queuedCoordinate.adjacentCoordinates()

				coordinates[queuedCoordinate] = true
				processedCoordinates[queuedCoordinate] = true

				for _, neighbor := range nestedNeighbors {
					if !neighbor.isValid(maxX, maxY) {
						continue
					}

					if processedCoordinates[neighbor] || processedNeighbors[neighbor] {
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
