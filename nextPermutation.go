//https://blog.csdn.net/menghan1224/article/details/52269064
package main

func nextPermutation(nums []int) {
	lastincr := -1
	for i := len(nums) - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			lastincr = i - 1
			break
		}
	}
	if lastincr == -1 {
		nextPermutationReverse(nums, 0, len(nums)-1)
	}
	// lastincr之后把最大的交换到lastincr 剩下的升序
	for i := len(nums) - 1; i > lastincr; i-- {
		if nums[i] > nums[lastincr] {
			nums[i], nums[lastincr] = nums[lastincr], nums[i]
		}
	}
	nextPermutationReverse(nums, lastincr+1, len(nums)-1)

}
func nextPermutationReverse(nums []int, start int, end int) {
	l, e := start, end
	for l < e {
		nums[l], nums[e] = nums[e], nums[l]
		l++
		e--
	}
}
