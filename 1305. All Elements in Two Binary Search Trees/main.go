package main

// Given two binary search trees root1 and root2.
// Return a list containing all the integers from both trees sorted in ascending order.
func main() {}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion version
func getAllElements(root1 *TreeNode, root2 *TreeNode) []int {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	inorder(&list1, root1)
	inorder(&list2, root2)

	i, j, ret := 0, 0, make([]int, 0, len(list1)+len(list2))
	for i < len(list1) && j < len(list2) {
		if list1[i] < list2[j] {
			ret = append(ret, list1[i])
			i++
		} else {
			ret = append(ret, list2[j])
			j++
		}
	}

	for i < len(list1) {
		ret = append(ret, list1[i])
		i++
	}

	for j < len(list2) {
		ret = append(ret, list2[j])
		j++
	}

	return ret
}

func inorder(nums *[]int, root *TreeNode) {
	if root == nil {
		return
	}

	inorder(nums, root.Left)
	*nums = append(*nums, root.Val)
	inorder(nums, root.Right)
}
