func twoSum(nums []int, target int) []int {
	result:=[]int{}
    for i:=0;i<len(nums);i++ {
		for j:=0;j<len(nums);j++{
			if i!=j{
				sum:=nums[i]+nums[j]
				if sum==target{
					result=append(result,i)
					result=append(result,j)
					return result
				}
			}
		}
	}
	return []int{}
}
