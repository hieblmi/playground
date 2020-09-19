/*

	This is a solution to a problem that was posed to a friend of mine during an Amazon coding interview.

	Problem:	Given is a grid of cells each representing a network router.
				Initially an arbitrary number of routers are selected to be
				patched with a specific version of software. On a single day
				each router	can update the software version of its neighbors.
				What is the minimum number of days that have to pass until all
				routers in the grid received the update?
*/
package main

import (
	"fmt"
)

var m int
var n int
var passedDays int
var grid [][]bool
var initialPos []int

func init() {
	m = 5
	n = 5
	passedDays = 0
	grid = make([][]bool, m)
	for i := range grid {
		grid[i] = make([]bool, n)
	}

	// Set initial positions on grid to true
	// Every pair referres to a point in the 2d grid, e.g. below (1,1),(1,3),(3,1),(3,3)
	initialPos = []int{1, 1, 1, 3, 3, 1, 3, 3}
	for i := 0; i < len(initialPos); i += 2 {
		grid[initialPos[i]][initialPos[i+1]] = true
	}
}

func spreadTheUpdates() {

	// Query neighbors of initial positions until no new neighbors are found
	for neighbors := initialPos; len(neighbors) > 1; passedDays++ {
		fmt.Printf("After day %d: Neighbors: %v\n", passedDays, neighbors)
		initialPos = []int{}
		for i := 0; i < len(neighbors); i += 2 {
			initialPos = append(initialPos, getNeighbors(neighbors[i], neighbors[i+1], grid)...)
		}
		neighbors = initialPos
	}

	fmt.Printf("Passed days: %d\n", passedDays-1)
}

func getNeighbors(i int, j int, grid [][]bool) []int {
	directions, neighbors := []int{-1, 0, 1}, []int{}
	for _, y := range directions {
		for _, x := range directions {
			if i+y >= 0 && i+y < len(grid) && j+x >= 0 && j+x < len(grid[0]) {
				if grid[i+y][j+x] {
					continue
				}
				neighbors = append(neighbors, i+y)
				neighbors = append(neighbors, j+x)
				grid[i+y][j+x] = true
			}
		}
	}
	return neighbors
}
