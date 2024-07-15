package sigremover

import (
	"os"
	"testing"
)

const (
	InputPath      = "test_data/input.hbs"
	OutputPath     = "test_data/output.hbs"
	WithoutComment = "test_data/without-comment.hbs"
)

func TestRemoveComment(t *testing.T) {
	// Load input and output
	// Input
	inputBytes, err := os.ReadFile(InputPath)
	if err != nil {
		t.Error(err)
	}
	input := string(inputBytes)
	// Expected
	expectedBytes, err := os.ReadFile(WithoutComment)
	if err != nil {
		t.Error(err)
	}
	expected := string(expectedBytes)
	// input := `
	// <title>Vernus</title>
	// <!-- Made with Postcards by Designmodo https://designmodo.com/postcards -->
	// <style>
	// `
	// expected := `
	// <title>Vernus</title>
	// <style>
	// `

	result := RemoveComment(input)

	if result != expected {
		t.Error("Result is different then expected")
	}
}
