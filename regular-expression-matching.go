func isMatch(s string, p string) bool {

	lengthS := len(s)
	lengthP := len(p)

	if lengthP == 0 {
		return lengthS == 0
	}

	//'same char' or '.'
	matchedCharDot := lengthS != 0 && (p[0] == s[0] || p[0] == '.')

	//star
	matchedStar := lengthP >= 2 && p[1] == '*'
	if matchedStar {
		return (isMatch(s, string(p[2:])) || (matchedCharDot && isMatch(string(s[1:]), p)))
	}

	//not star
	return matchedCharDot && isMatch(string(s[1:]), string(p[1:]))

}