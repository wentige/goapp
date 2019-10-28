package u

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

// there are no guarantees in life, but this should be pretty unique string
func RandomString() string {
	var a [20]byte
	_, err := rand.Read(a[:])
	if err == nil {
		return hex.EncodeToString(a[:])
	}
	return fmt.Sprintf("__--##%d##--__", rand.Int63())
}

// Wraps text at the specified column lineWidth on word breaks
func WordWrap(text string, lineWidth int) []string {
	words := strings.Fields(strings.TrimSpace(text))
	if len(words) == 0 {
		return []string{text}
	}

    lines := make([]string, 0)

	wrapped := words[0]
	spaceLeft := lineWidth - len(wrapped)
	for _, word := range words[1:] {
		if len(word)+1 > spaceLeft {
            lines = append(lines, wrapped)
			wrapped = word
			spaceLeft = lineWidth - len(word)
		} else {
			wrapped += " " + word
			spaceLeft -= 1 + len(word)
		}
	}
    lines = append(lines, wrapped)
	return lines
}
