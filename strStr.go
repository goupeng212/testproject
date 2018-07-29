package main

//import "fmt"

func strStr(haystack string, needle string) int {
	lenhay := len(haystack)
	lenneed := len(needle)
	if lenneed == 0 {
		return 0
	}
	if lenhay < lenneed {
		return -1
	}

	for i := 0; i < lenhay; i++ {
		j := 0
		for ; j < lenneed; j++ {

			if i+j > lenhay-1 || haystack[i+j] != needle[j] {
				break
			}
			if j == lenneed-1 {
				return i
			}
		}
	}
	return -1
}

//func main() {
//	hay := "mississippi"
//	needle := "sip"
//	fmt.Println(strStr(hay, needle))
//}
