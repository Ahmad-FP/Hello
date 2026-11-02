package mixed



func Truncate(s string, max int) string{
	if len([]rune(s)) <= max {
		return s
	} else{
		return s[:max] + "\u2026"
	}
}