package main

import "fmt"

// IsInterleave reports whether s3 is an interleaving of s1 and s2: all
// characters of both, relative orders preserved. One rolling row over s2's
// prefix length; dp[j] = "s1's first i and s2's first j can weave s3's
// first i+j".
func IsInterleave(s1, s2, s3 string) bool {
	n, m := len(s1), len(s2)
	if n+m != len(s3) {
		return false
	}
	dp := make([]bool, m+1)
	dp[0] = true
	for j := 1; j <= m; j++ {
		dp[j] = dp[j-1] && s2[j-1] == s3[j-1]
	}
	for i := 1; i <= n; i++ {
		dp[0] = dp[0] && s1[i-1] == s3[i-1]
		for j := 1; j <= m; j++ {
			fromS1 := dp[j] && s1[i-1] == s3[i+j-1]   // extend with s1's char
			fromS2 := dp[j-1] && s2[j-1] == s3[i+j-1] // extend with s2's char
			dp[j] = fromS1 || fromS2
		}
	}
	return dp[m]
}

func main() {
	fmt.Println(IsInterleave("aabcc", "dbbca", "aadbbcbcac"))
	fmt.Println(IsInterleave("aabcc", "dbbca", "aadbbbaccc"))
}
