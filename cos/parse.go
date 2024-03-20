package cos

import "fmt"

// GenerateMarkdownImageTag generates markdown image tag for the provided URL.
func GenerateMarkdownImageTag(imageURL string) string {
	return fmt.Sprintf("![](%s)", imageURL)
}
