package starter

import "testing"

func TestGetRockPaperAnswers(t *testing.T) {
	sampleInput := []string{
		"A Y", "B X", "C Z",
	}
	answer := GetRockPaperAnswers(sampleInput)
	expectation := 15
	if answer != expectation {
		t.Fatalf(`Rock paper scissors test: %d, we wanted a match for %d, nil`, answer, expectation)
	}
}
