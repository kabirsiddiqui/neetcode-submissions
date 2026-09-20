import "slices"
func isPalindrome(s string) bool {
	s=strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")
	s = strings.Map(func(r rune) rune {
    if unicode.IsLetter(r) || unicode.IsDigit(r) {
        return r
    }
    return -1
}, s)
	chars:=[]byte(s)
	slices.Reverse(chars)
	if string(chars)==s{
		return true
	}else{
		return false
	}
}
