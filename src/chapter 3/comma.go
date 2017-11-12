package main
func comma(s string) string{
	n := len(s)
	if n<3{
		return
	}
	return comma(s[:n-3])+","+s[n-3:]
}
