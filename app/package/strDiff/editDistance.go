package strDiff

import "unicode/utf8"

func EditDistance(s string, t string) int {
	// initialize
	table := make([][]int, utf8.RuneCountInString(s)+1)
	for index := range table {
		table[index] = make([]int, utf8.RuneCountInString(t)+1)
	}
	for i := 1; i <= utf8.RuneCountInString(s); i += 1 {
		table[i][0] = i
	}
	for j := 1; j <= utf8.RuneCountInString(t); j += 1 {
		table[0][j] = j
	}

	// calculate
	for i := 1; i <= utf8.RuneCountInString(s); i += 1 {
		s_i := []rune(s)[i-1]
		for j := 1; j <= utf8.RuneCountInString(t); j += 1 {
			var replaceCost int
			substring_j := []rune(t)[j-1]
			if s_i == substring_j {
				replaceCost = 0
			} else {
				replaceCost = 1
			}
			replace := table[i-1][j-1] + replaceCost
			delete := table[i-1][j] + 1
			insert := table[i][j-1] + 1
			costs := [...]int{delete, insert, replace}
			minimum := -1
			for _, c := range costs {
				if (minimum >= 0 && minimum > c) || minimum < 0 {
					minimum = c
				}
			}
			table[i][j] = minimum
		}
	}
	return table[utf8.RuneCountInString(s)][utf8.RuneCountInString(t)]
}
