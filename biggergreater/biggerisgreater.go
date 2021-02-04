package biggergreater

import (
	"fmt"
	"sort"
)

func Test() {

	// f, err := os.Open("testcase1.txt")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()

	// b := bufio.NewReader(f)
	// for {
	// 	line, err := b.ReadString('\n')
	// 	if err != nil {
	// 		if err == io.EOF {
	// 			break
	// 		}
	// 		log.Fatal(err)
	// 	}
	// 	line = strings.TrimSuffix(line, "\n")
	// 	r := biggerIsGreater(line)
	// 	fmt.Println(r)
	// }

	sa := []string{
		"biehzcmjckznhwrfgglverxsezxuqpj",
		// "biehzcmjckznhwrfgglverxs jepquxz"

		"rebjvsszebhehuojrkkhszxltyqfdvayusylgmgkdivzlpmmtvbsavxvydldmsym",
		// // "rebjvsszebhehuojrkkhszxltyqfdvayusylgmgkdivzlpmmtvbsavxvydldm yms"

		"jprfovzkdlmdcesdcpdchcwoedjchcovklhrhlzfeeptoewcqpxg",
		// // "jprfovzkdlmdcesdcpdchcwoedjchcovklhrhlzfeeptoewcq xgp"

		// "gwakhcpkolybihkmxyecrdhsvycjrljajlmlqgpcnmvvkjlkvdowzdfikh",
		// // "gwakhcpkolybihkmxyecrdhsvycjrljajlmlqgpcnmvvkjlkvdowzdf khi"

		// "nebsajjbbuifimjpdcqfygeitief",
		// // "nebsajjbbuifimjpdcqfygeiti fe"

		// "xuqfobndhsnsztifmqpnencxkllnpmbfqihtgdgxjhsvitlgtodhcydfb",
		// // "xuqfobndhsnsztifmqpnencxkllnpmbfqihtgdgxjhsvitlgtodhcy fbd"

		"pqommldkafmnwzidydgjghxcbnwyjdxpvmkztdfmcxlkargafjzeye",
		// "pqommldkafmnwzidydgjghxcbnwyjdxpvmkztdfmcxlkargafjzyee"

		"ynwjdeyufytvgltwzdjnzkafnzwlndzrixrjckc",
		// "ynwjdeyufytvgltwzdjnzkafnzwlndzrixrjkcc"

		"gxfbcmzkpcywzudenlrpenxudmjliaqkmbqbixmtlgmbvevxqjzvjpuptiprdsixcvrc",
		// "gxfbcmzkpcywzudenlrpenxudmjliaqkmbqbixmtlgmbvevxqjzvjpuptiprdsixrccv"

		"ab",
		"bb",
		"hefg",
		"dhck",
		"dkhc",
	}

	for _, s := range sa {
		r := biggerIsGreater(s)
		fmt.Println(r)
	}

}

// Complete the biggerIsGreater function below.
func biggerIsGreater(w string) string {

	var i int
	for i = len(w) - 1; i > 0; i-- {
		if w[i-1] < w[i] {
			break
		}
	}

	if i-1 < 0 {
		return "no answer"
	}

	b := []byte(w[i-1:])
	sort.SliceStable(b, func(i, j int) bool {
		return b[i] < b[j]
	})

	var j int
	for j = 0; j < len(b); j++ {
		if w[i-1] < b[j] {
			break
		}
	}
	return fmt.Sprintf("%s%c%s%s", w[:i-1], b[j], b[0:j], b[j+1:])
}
