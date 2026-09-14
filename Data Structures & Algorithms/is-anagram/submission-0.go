func isAnagram(s string, t string) bool {
	freq1:=map[byte]int{}
	for _,value:=range s{
		freq1[byte(value)]++
	}
	freq2:=map[byte]int{}
	for _,value:=range t{
		freq2[byte(value)]++
	}
	if sameFreq(freq1,freq2){
		return true
	}else{
		return false
	}
}
func sameFreq(a, b map[byte]int) bool {
    if len(a) != len(b) {
        return false
    }

    for key, count := range a {
        if b[key] != count {
            return false
        }
    }

    return true
}
