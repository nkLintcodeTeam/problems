package match

import (
	"fmt"
)

func isMatch(s string, p string) bool {
	sBytes := []byte(s)
	pBytes := []byte(p)
	dp := make([][]bool, len(s)+1)
	for k, _ := range dp {
		dp[k] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	for i := 0; i < len(p); i = i + 1 {
		if pBytes[i] == '*' {
			// * will not be the frist character in p, otherwise the pattern is illegal
			dp[0][i+1] = dp[0][i-1]
		}
	}
	fmt.Printf("dp[0]: %+v\n", dp[0])

	for i := 0; i < len(s); i = i + 1 {
		for j := 0; j < len(p); j = j + 1 {
			if sBytes[i] == pBytes[j] {
				dp[i+1][j+1] = dp[i][j]
			} else if pBytes[j] == '.' {
				dp[i+1][j+1] = dp[i][j]
			} else if pBytes[j] == '*' {
				// *: 0 or >0
				dp[i+1][j+1] = (dp[i+1][j-1]) || ((sBytes[i] == pBytes[j-1] || pBytes[j-1] == '.') && dp[i][j+1])
			}
		}
	}
	return dp[len(s)][len(p)]
}
