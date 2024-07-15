package sigremover

import (
	"fmt"
	"regexp"
)

func Remove() {
	fmt.Println("Sig remover")
}

func RemoveComment(input string) string {
	pattern := `(?m)^\s*<!--.*Made with Postcards.*-->\s*$\n?`
	re := regexp.MustCompile(pattern)
	cleanedHTML := re.ReplaceAllString(input, "")
	return cleanedHTML
}
