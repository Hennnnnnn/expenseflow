package filter

import "strings"

func IsBCA(sender string) bool {

	sender = strings.ToLower(sender)

	allowedSenders := []string{
		"klikbca",
		"bca.co.id",
		"mybca",
		"bank central asia",
	}

	for _, s := range allowedSenders {
		if strings.Contains(sender, s) {
			return true
		}
	}

	return false
}
