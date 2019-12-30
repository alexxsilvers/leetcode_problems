package main

import "sort"

// Given two binary search trees root1 and root2.
// Return a list containing all the integers from both trees sorted in ascending order.
func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	ret := make([]int, 0)

	getAllValues(root1, &ret)
	getAllValues(root2, &ret)

	sort.Ints(ret)
	return ret
}

func getAllValues(bTree *TreeNode, numbers *[]int) {
	if bTree == nil {
		return
	}

	*numbers = append(*numbers, bTree.Val)

	getAllValues(bTree.Left, numbers)
	getAllValues(bTree.Right, numbers)
}

func mergeChans(chs []<-chan int) chan int {
	resultCh := make(chan int, len(chs))
	for _, ch := range chs {
		select {
		case val := <-ch:
			resultCh <- val
		}
	}

	strconv.Ito

	return resultCh
}
