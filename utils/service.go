package utils

import (
	"crypto/rand"
	"encoding/hex"
	"regexp"
	"strings"
)

// GenSlug generates a slug from the title and adds a short random string at the end
func GenSlug(title string) string {
	// convert to lowercase
	slug := strings.ToLower(title)

	// replace spaces with hyphens
	slug = strings.ReplaceAll(slug, " ", "-")

	// remove non-alphanumeric characters except hyphen
	re := regexp.MustCompile(`[^\w-]`)
	slug = re.ReplaceAllString(slug, "")

	// generate 6-byte random string (12 hex chars)
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		panic(err) // you can handle error differently
	}
	randomStr := hex.EncodeToString(b)

	// combine slug and random string
	return slug + "-" + randomStr
}
