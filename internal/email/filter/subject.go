package filter

import "strings"

func IsTransaction(subject string) bool {

	subject = strings.ToLower(subject)

	keywords := []string{
		"transaksi",
		"debit",
		"kartu kredit",
		"mutasi",
		"pendebetan",
		"bca",
		"credit card",
	}

	for _, keyword := range keywords {

		if strings.Contains(subject, keyword) {
			return true
		}
	}

	return false
}
