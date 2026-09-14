func hasDuplicate(nums []int) bool {
    freq:=map[int]int{}
    for _,value:=range nums{
        freq[value]++
    }
    for _,value:=range nums{
        if freq[value]>1{
            return true
        }
    }
    return false
}
