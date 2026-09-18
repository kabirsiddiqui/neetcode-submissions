func topKFrequent(nums []int, k int) []int {
    freq:=map[int]int{}
    for _,value := range nums{
        freq[value]++
    }
    temp:=[]int{}
    for num:=range freq{
        temp=append(temp,num)
    }
    sort.Slice(temp, func(i,j int) bool{
        return freq[temp[i]]>freq[temp[j]]
    })
    return temp[:k]
}
