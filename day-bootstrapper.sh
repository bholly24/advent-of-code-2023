#!/bin/bash

if [ $# -eq 0 ]; then
    echo "Usage: $0 [directory_name]"
    exit 1
fi

mkdir "day0$1"
echo "package day0$1" > "day0$1/day0$1.go"

testing_file="day0$1/day0$1_test.go"
#
#echo "package day0$1" > "$testing_file"
#echo "" >> "$testing_file"
#echo "import \"testing\"" >> "$testing_file"

cat <<EOF > "$testing_file"
package day0$1

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
EOF

echo "Project set up for day '$1' created successfully!"
