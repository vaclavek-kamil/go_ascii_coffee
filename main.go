package main

import (
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

func main() {
	// Init
	var shader [10][14]int
	const MAX_HEALTH int = 11
	const MIN_HEALTH int = 4
	const DEGRADATION = 1.2
	const FPS = 3
	const lpad = "    "
	ASCII := [14]string{" ", ".", ",", "-", "*", ";", "!", "|", "l", "I", "§", "%"}
	for {
		// Display frame
		for i := 0; i < len(shader); i++ {
			fmt.Printf("%s  ", lpad)
			for j := 0; j < len(shader[i]); j++ {
				fmt.Printf("%s ", ASCII[shader[i][j]])
			}
			fmt.Printf("\n")
		}
		fmt.Println(lpad, "_____________________________")
		fmt.Println(lpad, "\\                           /")
		fmt.Println(lpad, " \\                         /__")
		fmt.Println(lpad, "  \\                       /    \\")
		fmt.Println(lpad, "   \\                     /     /")
		fmt.Println(lpad, "    \\                   /     /")
		fmt.Println(lpad, "     \\                 /-----")
		fmt.Println(lpad, "      \\               /      ")
		fmt.Println(lpad, "       ---------------        ")

		// Move existing particles up and decrease health
		for x := 0; x < len(shader); x++ {
			for y := 0; y < len(shader[x]); y++ {
				if shader[x][y] > 0 && x != 0 {
					shader[x-1][y] = shader[x][y] - 1
				}
				shader[x][y] = 0
			}
		}

		// Create new particles
		bottomRow := shader[len(shader)-1]
		for i := 0; i < len(bottomRow); i++ {
			distanceFromCenter := int(math.Abs(float64((len(bottomRow) / 2) - i)))
			if rand.IntN(int(2+DEGRADATION*float64(distanceFromCenter))) == 0 {
				// RNG check passed. Creating rng health particle
				bottomRow[i] = MIN_HEALTH + rand.IntN(MAX_HEALTH-MIN_HEALTH) - int(float64(distanceFromCenter)*DEGRADATION)
				if bottomRow[i] < 0 {
					bottomRow[i] = 0
				}
			}
		}
		shader[len(shader)-1] = bottomRow
		// Wait for x time, clear the fame and proceed
		time.Sleep(time.Second / FPS)
		for i := 0; i < len(shader)+9; i++ {
			fmt.Printf("\033[1A\033[K")
		}
	}
}
