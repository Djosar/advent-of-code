package parsel

import (
	// Adjust the import path as necessary
	"advent-of-code-2023/util/geo"
	"bufio"
	"os"
	"strings"
)

func ParseFileToMatrix(filePath string) (*geo.Matrix[string], error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lineCount, lineLength int
	for scanner.Scan() {
		lineCount++
		if lineLength == 0 {
			lineLength = len(scanner.Text())
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	matrix := geo.NewMatrix[string](lineLength, lineCount)

	file.Seek(0, 0)
	for index := 0; scanner.Scan(); index++ {
		line := scanner.Text()
		row := strings.Split(line, "")
		for idx, char := range row {
			matrix.Set(index, idx, char)

		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}
