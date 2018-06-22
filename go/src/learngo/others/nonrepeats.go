package others

import "fmt"

func nonOfLongestReoeateStr(s string) int {
	lastOccurred := make(map[rune]int)
	start := 0
	maxlength := 0

	for i, ch := range []rune(s) {
		if lastI,ok := lastOccurred[ch];ok && lastI >= start{
			start = lastI + 1
		}

		if i - start + 1 > maxlength{
			maxlength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxlength
}

func main() {
	fmt.Println(
		nonOfLongestReoeateStr("abcabcbb"))
	fmt.Println(
		nonOfLongestReoeateStr("bbbbb"))
	fmt.Println(
		nonOfLongestReoeateStr(""))
	fmt.Println(
		nonOfLongestReoeateStr("b"))
	fmt.Println(
		nonOfLongestReoeateStr("abcdef"))
	fmt.Println(
		nonOfLongestReoeateStr("阿森纳长发少女"))
	fmt.Println(
		nonOfLongestReoeateStr("一二三二一"))
}