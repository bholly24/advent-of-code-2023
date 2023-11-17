package printer

import "fmt"

func PrintIntroText() {
	printDivider()
	fmt.Println("Welcome to my advent of code for 2023. All answers are my own.")
}

func PrintAnswerDivider(day int, part int) {
	printDivider()
	fmt.Printf("Ho ho ho - Solution for day %d part %d\n", day, part)
}

func printDivider() {
	fmt.Println("------------------------------------------------------------------")
}

func PrintOutro() {
	printDivider()
}
