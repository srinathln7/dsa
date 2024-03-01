package main

import "fmt"

func combinationSum(candidates []int, target int) [][]int {
	// Key Idea: This is a variation of a generation of all the subsets problem
	var result [][]int
	dfs(candidates, []int{}, &result, target, 0)
	return result
}

func dfs(candidates, combination []int, result *[][]int, target, startIdx int) {

	fmt.Printf("Target at index %d => %d\n", startIdx, target)
	// When the desired target is reached append the result to the count
	if target == 0 {
		*result = append(*result, append([]int{}, combination...))
		return
	}

	// Since the array consists of only positive numbers
	if target < 0 {
		return
	}

	for i := startIdx; i < len(candidates); i++ {
		combination = append(combination, candidates[i])
		fmt.Printf("combination at index %d => %d \n", i, combination)
		dfs(candidates, combination, result, target-candidates[i], i)

		// Backtrack the combination slice => Similar to the generation of a subset to the set problem
		fmt.Printf("Back tracking at index %d \n", i)
		combination = combination[:len(combination)-1]
	}
}

func main() {
	nums := []int{1, 7}
	target := 8

	fmt.Printf("Generating all possible combinations in %d who sum is equal to %d\n", nums, target)
	combinationSum(nums, target)
}
