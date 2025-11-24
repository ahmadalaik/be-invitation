package slug

import (
	"regexp"
	"strings"
)

func GenerateSlug(s string) string {
	slug := strings.ToLower(s)

	reNonAlphaNum := regexp.MustCompile(`[^a-z0-9\s]`)
	slug = reNonAlphaNum.ReplaceAllString(slug, "")

	reSpaces := regexp.MustCompile(`\s+`)
	slug = reSpaces.ReplaceAllString(slug, "-")

	reRepeatedHyphens := regexp.MustCompile(`-+`)
	slug = reRepeatedHyphens.ReplaceAllString(slug, "-")

	slug = strings.Trim(slug, "-")

	return slug
}
