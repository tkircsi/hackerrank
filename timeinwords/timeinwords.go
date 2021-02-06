package timeinwords

import "fmt"

func Test() {
	clocks := [][]int32{
		{3, 00}, {7, 15}, {6, 30}, {8, 45}, {12, 00}, {12, 30}, {1, 1}, {1, 59},
	}
	for _, c := range clocks {
		r := timeInWords(c[0], c[1])
		fmt.Printf("%d:%d - %s\n", c[0], c[1], r)
	}
}

// Complete the timeInWords function below.
func timeInWords(h int32, m int32) string {
	nums := map[int32]string{
		0: "zero", 1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven", 8: "eight", 9: "nine",
		10: "ten", 11: "eleven", 12: "twelve", 13: "thirteen", 14: "fourteen", 15: "fifteen", 16: "sixteen", 17: "seventeen", 18: "eighteen", 19: "nineteen", 20: "twenty", 21: "twenty one", 22: "twenty two", 23: "twenty three", 24: "twenty four", 25: "twenty five", 26: "twenty six", 27: "twenty seven", 28: "twenty eight", 29: "twenty nine", 30: "thirty",
	}
	var sm string
	switch {
	case m == 0:
		sm = fmt.Sprintf("%s o' clock", nums[h])
	case m == 15:
		sm = fmt.Sprintf("quarter past %s", nums[h])
	case m == 30:
		sm = fmt.Sprintf("half past %s", nums[h])
	case m == 45:
		sm = fmt.Sprintf("quarter to %s", nums[h+1])
	case m == 1:
		sm = fmt.Sprintf("%s minute past %s", nums[m], nums[h])
	case 0 < m && m < 30:
		sm = fmt.Sprintf("%s minutes past %s", nums[m], nums[h])
	case m == 59:
		sm = fmt.Sprintf("%s minute to %s", nums[60-m], nums[h+1])
	case 30 < m && m <= 59:
		sm = fmt.Sprintf("%s minutes to %s", nums[60-m], nums[h+1])
	default:
		sm = "undefined"
	}
	return sm
}
