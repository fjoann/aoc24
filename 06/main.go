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

type direction struct {
	dx, dy int
}

type guard struct {
	pos position
	dir direction
}

func (g *guard) move(p position) {
	g.pos = p
}

func (g *guard) turnRight() {
	g.dir.dx, g.dir.dy = -g.dir.dy, g.dir.dx
}

type step struct {
	pos position
	dir direction
}

const (
	start    = '^'
	space    = '.'
	obstacle = '#'
)

func main() {
	input, err := os.Open("06/input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	var labMap [][]rune
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		row := []rune(scanner.Text())
		labMap = append(labMap, []rune(row))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	labBounds := struct {
		x, y int
	}{
		x: len(labMap[0]) - 1,
		y: len(labMap) - 1,
	}

	directions := struct {
		up, right, down, left direction
	}{
		up:    direction{0, -1},
		right: direction{1, 0},
		down:  direction{0, 1},
		left:  direction{-1, 0},
	}

	// Part 1: visited positions
	var initialPosition position
	for y, row := range labMap {
		for x, tile := range row {
			if tile == start {
				initialPosition = position{x, y}
			}
		}
	}
	initialDirection := directions.up
	guard := &guard{
		pos: initialPosition,
		dir: initialDirection,
	}

	var path []step
	visitedPositions := make(map[position]struct{})

	initialStep := step{
		pos: guard.pos,
		dir: guard.dir,
	}
	path = append(path, initialStep)
	visitedPositions[guard.pos] = struct{}{}

	for {
		nextPosition := position{
			x: guard.pos.x + guard.dir.dx,
			y: guard.pos.y + guard.dir.dy,
		}

		if nextPosition.x < 0 || nextPosition.x > labBounds.x ||
			nextPosition.y < 0 || nextPosition.y > labBounds.y {
			break
		}

		if labMap[nextPosition.y][nextPosition.x] == obstacle {
			guard.turnRight()
			continue
		}

		guard.move(nextPosition)

		step := step{
			pos: guard.pos,
			dir: guard.dir,
		}
		path = append(path, step)

		visitedPositions[guard.pos] = struct{}{}
	}

	fmt.Printf("Visited positions: %d\n", len(visitedPositions))

	// Part 2: obstructions & loops
	obstructionPositionsCausingLoops := make(map[position]struct{})
	obstructionPositionsChecked := make(map[position]struct{})

	for i := 0; i < len(path)-1; i++ {
		obstructionPosition := path[i+1].pos
		if _, alreadyChecked := obstructionPositionsChecked[obstructionPosition]; alreadyChecked {
			continue
		}
		obstructionPositionsChecked[obstructionPosition] = struct{}{}
		labMap[obstructionPosition.y][obstructionPosition.x] = obstacle

		guard.pos = path[i].pos
		guard.dir = path[i].dir
		stepsTaken := make(map[step]struct{})

		initialStep := step{
			pos: guard.pos,
			dir: guard.dir,
		}
		stepsTaken[initialStep] = struct{}{}

		for {
			nextPosition := position{
				x: guard.pos.x + guard.dir.dx,
				y: guard.pos.y + guard.dir.dy,
			}

			if nextPosition.x < 0 || nextPosition.x > labBounds.x ||
				nextPosition.y < 0 || nextPosition.y > labBounds.y {
				break
			}

			if labMap[nextPosition.y][nextPosition.x] == obstacle {
				guard.turnRight()
				continue
			}

			guard.move(nextPosition)

			step := step{
				pos: guard.pos,
				dir: guard.dir,
			}
			if _, alreadyTaken := stepsTaken[step]; alreadyTaken {
				obstructionPositionsCausingLoops[obstructionPosition] = struct{}{}
				break
			}

			stepsTaken[step] = struct{}{}
		}

		labMap[path[i+1].pos.y][path[i+1].pos.x] = space
	}

	fmt.Printf("Obstruction positions causing loops: %d\n", len(obstructionPositionsCausingLoops))
}
