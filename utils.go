package genum

import "strings"

func makeComment(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	parts := make([]string, len(lines))
	for i, l := range lines {
		parts[i] = "// " + l
	}
	return strings.Join(parts, "\n")
}
