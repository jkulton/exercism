package bob

import (
	"strings"
)

// Hey determines what Bob should say.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if remark == "" {
		return "Fine. Be that way!"
	}

	hasLetters := strings.ToUpper(remark) != strings.ToLower(remark)
	isUpper := strings.ToUpper(remark) == remark
	isYelling := hasLetters && isUpper
	lastChar := remark[len(remark)-1:]
	isQuestion := lastChar == "?"

	if isYelling && isQuestion {
		return "Calm down, I know what I'm doing!"
	}

	if isYelling {
		return "Whoa, chill out!"
	}

	if isQuestion {
		return "Sure."
	}

	return "Whatever."
}
