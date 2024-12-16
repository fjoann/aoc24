package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type position struct {
	x, y int
}

type move struct {
	dx, dy int
}

type robot struct {
	pos position
}

func (r *robot) moveTo(p position) {
	r.pos = p
}

type warehouse struct {
	grid [][]rune
}

func (w *warehouse) getTile(p position) rune {
	return w.grid[p.y][p.x]
}

func (w *warehouse) updateTile(p position, tile rune) {
	w.grid[p.y][p.x] = tile
}

const (
	robotTile = '@'
	spaceTile = '.'
	wallTile  = '#'
	boxTile   = 'O'
)

func main() {
	input, err := os.Open("15/input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	directions := map[rune]move{
		'^': {0, -1},
		'>': {1, 0},
		'v': {0, 1},
		'<': {-1, 0},
	}

	var warehouse warehouse
	var moves []rune
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		if len(row) == 0 {
			break
		}
		warehouse.grid = append(warehouse.grid, row)
	}
	for scanner.Scan() {
		row := []rune(scanner.Text())
		moves = append(moves, row...)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var robot robot
	for y, row := range warehouse.grid {
		for x, tile := range row {
			if tile == robotTile {
				robot.pos = position{x, y}
			}
		}
	}

	for _, move := range moves {
		initialPosition := position{
			x: robot.pos.x,
			y: robot.pos.y,
		}
		nextPosition := position{
			x: robot.pos.x + directions[move].dx,
			y: robot.pos.y + directions[move].dy,
		}
		nextTile := warehouse.getTile(nextPosition)

		switch nextTile {
		case wallTile:
			continue
		case boxTile:
			for numBoxes := 1; ; {
				forwardDistance := numBoxes + 1

				lastPosition := position{
					x: robot.pos.x + forwardDistance*directions[move].dx,
					y: robot.pos.y + forwardDistance*directions[move].dy,
				}
				lastTile := warehouse.getTile(lastPosition)

				if lastTile == boxTile {
					numBoxes++
					continue
				}
				if lastTile == spaceTile {
					robot.moveTo(nextPosition)
					warehouse.updateTile(lastPosition, boxTile)
					warehouse.updateTile(nextPosition, robotTile)
					warehouse.updateTile(initialPosition, spaceTile)
					break
				}
				break
			}
		case spaceTile:
			robot.moveTo(nextPosition)
			warehouse.updateTile(nextPosition, robotTile)
			warehouse.updateTile(initialPosition, spaceTile)
		}
	}

	var gpsCoordinateSum int
	for y, row := range warehouse.grid {
		for x, tile := range row {
			if tile == boxTile {
				gpsCoordinate := y*100 + x
				gpsCoordinateSum += gpsCoordinate
			}
		}
	}

	fmt.Printf("GPS coordinate sum: %d\n", gpsCoordinateSum)
}
