package main

func main() {

}

func sumZero(n int) []int {
	s := 0
	tmp := n
	ret := make([]int, 0, n)

	for i := 0; i < n; i++ {
		if i == n-1 { // last
			ret = append(ret, s*-1)
			break
		}

		tmp -= 1
		s += tmp
		ret = append(ret, tmp)
	}

	return ret
}
