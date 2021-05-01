package ctci

import "testing"

func Test_IsUniqueChars(t *testing.T) {
	testMap := map[string]bool{
		"asdfgh":    true,
		"aasdfgh":   false,
		"a":         true,
		"":          true,
		"阿里巴巴及蚂蚁集团": false,
	}

	for k, v := range testMap {
		got := isUniqueChars(k)
		if got != v {
			t.Errorf("IsUniqueChars failed. Got %v - want %v", got, v)
		}
	}
}
