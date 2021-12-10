package day1

func SonarSweep(readings []int) int {
	readingsHigherThanPrevious := 0
	previousValue := -1 // I assume sonar will not return negative distances

	for _, value := range readings {
		if value > previousValue && previousValue != -1 {
			readingsHigherThanPrevious++
		}

		previousValue = value
	}

	return readingsHigherThanPrevious
}

type slidingWindow [3]int

func emptySlidingWindow() slidingWindow {
	return slidingWindow{-1, -1, -1}
}

func (window slidingWindow) isFull() bool {
	return window[0] != -1 &&
		window[1] != -1 &&
		window[2] != -1
}

func (window slidingWindow) push(measurement int) slidingWindow {
	return slidingWindow{window[1], window[2], measurement}
}

func (window slidingWindow) sum() int {
	sum := 0

	for _, measurement := range window {
		if measurement == -1 {
			continue
		}

		sum += measurement
	}

	return sum
}

func SlidingSonarSweep(readings []int) int {
	readingsHigherThanPrevious := 0
	previousWindow := emptySlidingWindow()

	for _, value := range readings {
		currentWindow := previousWindow.push(value)

		if currentWindow.isFull() && previousWindow.isFull() && (currentWindow.sum() > previousWindow.sum()) {
			readingsHigherThanPrevious++
		}

		previousWindow = currentWindow
	}

	return readingsHigherThanPrevious
}
