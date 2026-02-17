package asciiart

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	charStart  = 32
	charEnd    = 126
	charHeight = 8
)

func Generate(text string, banner string) (string, error) {

	file, err := os.Open("banners/" + banner + ".txt")
	if err != nil {
		return "", fmt.Errorf("banner not found")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if scanner.Text() != "" {
			break
		}
	}

	asciiMap := make(map[rune][]string)

	char := rune(charStart)
	lines := []string{scanner.Text()}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			if len(lines) == charHeight {
				asciiMap[char] = lines
				char++
			}
			lines = []string{}
			continue
		}

		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	var result strings.Builder
	words := strings.Split(text, "\n")

	for _, word := range words {
		for row := 0; row < charHeight; row++ {
			for _, ch := range word {
				if ch < charStart || ch > charEnd {
					continue
				}
				result.WriteString(asciiMap[ch][row])
			}
			result.WriteByte('\n')
		}
	}

	return result.String(), nil
}
