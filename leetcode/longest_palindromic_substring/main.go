package main

import "fmt"

func main() {
	fmt.Printf("%v\n", longestPalindrome("babad"))
	fmt.Printf("%v\n", longestPalindrome("ababad"))
	fmt.Printf("%v\n", longestPalindrome("cbbd"))
	fmt.Printf("%v\n", longestPalindrome("a"))
	fmt.Printf("%v\n", longestPalindrome("ac"))
	//fmt.Printf("%v\n", longestPalindrome("jsfpmgkuxqnmtruslzgyvexhqjoamvyuyedhybqqcjhhhgmwqudgldvspgugibdsqfhucpfcqzriqqusvspgbqhgkswlzdkytyqphexemwxpduxplkquvgvhefsxubjluopighxbpscekijrqjhcgmqcuoczwbvueuviyfokdoqqsckjdorsettkkpiyyxxdsfczyhkyxlvrmhvflqbvlrukqcplbxnyokdxvhubsisxrodolmpmkdczavqlsnrggffagoddaldlcexwvozjxxdjtfjrfciwpacpbajcpmgfpefngqfbzehaaqyfvthtrbhkzrzqmzdcgrkezpqgbqjembeqaziuubbvdfpfyqanilcjggkudsyigiqgrcmauyugyhepvduudvpehyguyuamcrgqigiysdukggjclinaqyfpfdvbbuuizaqebmejqbgqpzekrgcdzmqzrzkhbrthtvfyqaahezbfqgnfepfgmpcjabpcapwicfrjftjdxxjzovwxecldladdogaffggrnslqvazcdkmpmlodorxsisbuhvxdkoynxblpcqkurlvbqlfvhmrvlxykhyzcfsdxxyyipkkttesrodjkcsqqodkofyivueuvbwzcoucqmgchjqrjikecspbxhgipouljbuxsfehvgvuqklpxudpxwmexehpqytykdzlwskghqbgpsvsuqqirzqcfpcuhfqsdbigugpsvdlgduqwmghhhjcqqbyhdeyuyvmaojqhxevygzlsurtmnqxukgmpfsj"))
	//fmt.Printf("%v\n", longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}

func longestPalindrome(s string) string {
	longest := ""

	validated := make(map[string]bool)
	for i := 0; i < len(s); i++ {
		for j := len(s) - 1; j >= i; j-- {
			start := i
			end := j
			valid := true
			for valid && end > start {
				current := s[start : end+1]
				currentValid, found := validated[current]
				if found && currentValid {
					break
				}
				if s[start] != s[end] {
					valid = false
				}
				start++
				end--
			}
			subString := s[i : j+1]
			validated[subString] = valid
			if valid && len(subString) > len(longest) {
				longest = subString
			}
		}
	}

	return longest
}
