package ctci

import "testing"

func Test_IsPermutation(t *testing.T) {
	type testCase struct {
		first  string
		second string
	}
	cases := map[testCase]bool{
		testCase{
			first:  "asdfg",
			second: "gfdsa",
		}: true,
		testCase{
			first:  "asdfg",
			second: "gfds",
		}: false,
		testCase{
			first:  "",
			second: "4",
		}: false,
		testCase{
			first:  "",
			second: "",
		}: true,
	}

	for testCase, want := range cases {
		got := isPermutation(testCase.first, testCase.second)
		if got != want {
			t.Errorf("isPermutation failed. got %v - want %v", got, want)
		}
	}
}
