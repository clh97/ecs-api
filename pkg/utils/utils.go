package utils

import (
	"fmt"
	"regexp"
	"strings"
)

// Slugify transforms text with spacebars to lowercase text with dashes
func Slugify(text string) string {
	re, err := regexp.Compile("[^a-z0-9]+")

	if err != nil {
		fmt.Println(err)
	}

	return strings.Trim(re.ReplaceAllString(strings.ToLower(text), "-"), "-")
}
