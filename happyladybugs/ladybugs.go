package happyladybugs

import "fmt"

func Test() {
	tcs := []string{
		"RBY_YBR",    // YES
		"X_Y__X",     // NO
		"__",         // YES
		"B_RRBR",     // YES
		"AABBC",      // NO
		"AABBC_C",    // YES
		"_",          // YES
		"DD__FQ_QQF", // YES
		"AABCBC",     // NO
	}

	for _, tc := range tcs {
		r := happyLadybugs(tc)
		fmt.Println(r)
	}
}

// Complete the happyLadybugs function below.
// Use cases:
//	1. Only underscore														YES		"___"
//  2. Only letters but in row										YES		"AAACCBBCC"
//  3. At least one underscore, no letter alone		YES		"AABBC_C"
//	4. Only one from a letter											NO		"X_Y__X"
//	5. No underscore and letters not in pair			NO		"AABCBC"
func happyLadybugs(b string) string {
	h := map[rune]int{}
	for _, c := range b {
		h[c]++
	}

	// Check if any letters is alone
	for k, v := range h {
		if k != '_' && v == 1 {
			return "NO"
		}
	}

	// Check if no underscore and letters not in pairs
	if _, ok := h['_']; !ok {
		for i := 1; i < len(b)-1; i++ {
			if b[i-1] != b[i] && b[i+1] != b[i] {
				return "NO"
			}
		}
	}

	return "YES"
}
