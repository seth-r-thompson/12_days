package io

import "fmt"

func ScansWithPrompt(prompt string, count int) ([]string, error) {
	inputs := make([]string, count)
	for i := 0; i < count; i++ {
		input, err := ScanWithPrompt(prompt)
		if err != nil {
			return inputs, err
		}
		inputs[i] = input
	}
	return inputs, nil
}

func ScanWithPrompt(prompt string) (string, error) {
	fmt.Print(prompt, " ")
	var input string
	_, err := fmt.Scanln(input)
	if err != nil && err.Error() != "unexpected newline" { // "unexpected newline" is scanned < args error, which is fine
		return "", err
	}
	return input, nil
}
