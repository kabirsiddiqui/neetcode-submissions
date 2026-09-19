func productExceptSelf(nums []int) []int {
	n:=len(nums)
	result:=make([]int,n)
	zeroCount:=0
	for _,value:=range nums{
		if value==0{
			zeroCount++
		}
	}
	if zeroCount>1{
		return result
	}
	if zeroCount==1{
		prod:=1
		for _,value:=range nums{
			if value!=0{
				prod*=value
			}
		}
		for index,value:=range nums{
			if value==0{
				result[index]=prod
				return result
			}
		}
	}
	total_product:=1
	for _,value:=range nums{
		total_product*=value
	}
	for i:=0;i<len(nums);i++{
		result[i]=total_product/nums[i]
	}
	return result
}
