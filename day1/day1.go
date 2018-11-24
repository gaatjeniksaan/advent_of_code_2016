package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	in, err := ioutil.ReadFile("./data.txt")
	if err != nil {
		fmt.Println("ERROR ERROR!")
	}

	inStr := string(in)
	data := strings.Split(inStr, ", ")

	facing := 0 // 0 = North, 1 = East, 2 = South, 3 = West
	x := 0      // x coordinate is East (pos) <-> West (neg)
	y := 0      // y coordinate is North (pos) <-> South (neg)

	for _, value := range data {
		dir := string(value[0])
		facing = updateFacing(dir, facing)
		stepSize, err := strconv.Atoi(strings.Trim(value[1:], "\n"))
		if err != nil {
			fmt.Println("String to int conversion failed: ", stepSize, err)
		}
		x, y = updateDirections(facing, stepSize, x, y)
	}

	result := shortestDistance(x, y)
	fmt.Println("The shortest distance is: ", result)
}

func updateFacing(dir string, facing int) int {
	if dir == "R" {
		facing = facing + 1
	} else {
		facing = facing - 1
	}
	facing = (facing + 4) % 4 // We don't want a negative facing value

	return facing
}

func updateDirections(facing int, stepSize, x, y int) (int, int) {
	if facing == 0 {
		y = y + stepSize
	} else if facing == 1 {
		x = x + stepSize
	} else if facing == 2 {
		y = y - stepSize
	} else if facing == 3 {
		x = x - stepSize
	} else {
		fmt.Println("Error, facing not within range 0-3")
	}

	return x, y
}

func shortestDistance(x, y int) float64 {
	a := float64(x)
	b := float64(y)
	a = math.Abs(a)
	b = math.Abs(b)

	return a + b
}
