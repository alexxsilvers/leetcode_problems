package main

func main() {

}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */

//
//func minOperations(nums []int) int {
//	op := 0
//
//	sort.Ints(nums)
//
//	for _, n := range nums {
//		if n % 2 ==
//	}
//
//	return op
//}

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	memo := make(map[int]int, 6)
	for i := 0; i < n; i++ {
		memo[i] = 0
	}

	for _, edge := range edges {
		memo[edge[1]]++
	}

	res := make([]int, 0)
	for k, v := range memo {
		if v == 0 {
			res = append(res, k)
		}
	}

	return res
}

func thousandSeparator(n int) string {
	if n == 0 {
		return "0"
	}

	res := make([]byte, 0)
	tCounter := 0
	for n > 0 {
		res = append(res, byte(n%10)+'0')
		tCounter++

		if tCounter%3 == 0 {
			res = append(res, '.')
		}

		n = n / 10
	}

	for i := 0; i < len(res)/2; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	if res[0] == '.' {
		res = res[1:]
	}

	return string(res)
}
