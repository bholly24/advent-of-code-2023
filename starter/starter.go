package starter

import (
	"fmt"
	"os"
	"strings"
)

func ReadInLines(file string) []string {
	content, err := os.ReadFile(file)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func GetRockPaperAnswers(input []string) int {
	score := 0
	for i := 0; i < len(input); i++ {
		score += getScore(input[i])
	}
	return score
}

func PartOneAnswer() {
	test := ReadInLines("./starter/input.txt")
	fmt.Printf("Here is score %d\n", GetRockPaperAnswers(test))
}

func getScore(moves string) int {
	win := 6
	lose := 0
	tie := 3
	rock := 1
	paper := 2
	scissors := 3
	scores := map[string]int{
		"A X": tie + rock,
		"B X": lose + rock,
		"C X": win + rock,
		"A Y": win + paper,
		"B Y": tie + paper,
		"C Y": lose + paper,
		"A Z": lose + scissors,
		"B Z": win + scissors,
		"C Z": tie + scissors,
	}

	return scores[moves]
}
