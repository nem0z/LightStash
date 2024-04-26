package utils

import (
	"bufio"
	"os"
)

func MoveCursorStart(file *os.File) error {
	_, err := file.Seek(0, 0)
	return err
}

func CountLines(file *os.File) (int, error) {
	if err := MoveCursorStart(file); err != nil {
		return 0, err
	}

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {
		count++
	}

	return count, scanner.Err()
}

func ReadLines(file *os.File, indexs ...int) ([]string, error) {
	if err := MoveCursorStart(file); err != nil {
		return []string{}, err
	}

	scanner := bufio.NewScanner(file)
	lines := make([]string, len(indexs))
	fileIdx := 0

	for i := range lines {
		for scanner.Scan() {
			if fileIdx == indexs[i] {
				lines[i] = scanner.Text()
				fileIdx++
				break
			}

			fileIdx++
		}
	}

	return lines, scanner.Err()
}
