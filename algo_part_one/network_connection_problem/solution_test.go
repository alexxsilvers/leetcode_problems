package network_connection_problem

import "testing"

func Test_Solution(t *testing.T) {
	members := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i"}
	log := []string{
		"1 a h",
		"2 f i",
		"3 d e",
		"4 g e",
		"5 e h",
		"6 b f",
		"7 c h",
		"8 b e",
		"9 a c",
		"10 f e",
	}

	ts := earliestConnectTime(members, log)
	if ts != "8" {
		t.Errorf("want %s, got %s", "8", ts)
	}
}
