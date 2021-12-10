package utilities

import (
	"bufio"
	"os"
	"strconv"
)

func ReadFileToIntSlice(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	result := []int{}
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, number)
	}

	return result, scanner.Err()
}
