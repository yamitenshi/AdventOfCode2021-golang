package main

import (
	"AdventOfCode2021-golang/day1"
	"AdventOfCode2021-golang/utilities"
	"fmt"
	"path"
	"runtime"
)

func main() {
	_, currentFilename, _, _ := runtime.Caller(0)
	currentDirectory := path.Dir(currentFilename)

	inputs, _ := utilities.ReadFileToIntSlice(currentDirectory + "/day1/input")

	fmt.Printf("Regular sonar sweep: %d\n", day1.SonarSweep(inputs))
	fmt.Printf("Sliding sonar sweep: %d\n", day1.SlidingSonarSweep(inputs))
}
