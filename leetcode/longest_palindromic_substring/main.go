package main

import "fmt"

func main() {
	fmt.Printf("%v\n", longestPalindrome("babad") == "bab")
	fmt.Printf("%v\n", longestPalindrome("ababad") == "ababa")
	fmt.Printf("%v\n", longestPalindrome("cbbd") == "bb")
	fmt.Printf("%v\n", longestPalindrome("a") == "a")
	fmt.Printf("%v\n", longestPalindrome("ac") == "a")
	fmt.Printf("%v\n", longestPalindrome("babaddtattarrattatddetartrateedredividerb") == "ddtattarrattatdd")
	fmt.Printf("%v\n", "lzyzl" == longestPalindrome("lfzwymgfcaqlozazylwpafzwgjhxbibllvdgsaiadtpryangehchkwkprhzpbszkobjcfmhffqxdcvghqftqyxorllrrheptcrhhlhrytwkytqmqlnvgoogjdlejnslpehtndmounvrtxplzyzlvcyuviknxoyhomwjzigiufhwqmjnwqpuukcxxhatxrgqiayqkkuwbxxbyejvxjpiflbeqjqemvkzcayitimalelkqmvrydiknqeghabhfuogujutrnzkmqqphbnrbnxhlgotbyghsbgmxschmbuhkobwvwajkcghrmgvvfzmxmaihcenxerznbnkotjubqxhbfqrcwsyfeowixusgfdgreywudrxjbylrnydtpfawayptifhlbmvrklplxahkxqahqalwsivszwvblpnozfmabzmouaxxbvbsibbzirgiqurhoitzlmpsovcjnkbeeydtkpelxmaulsvozwomofyvcafcenaprlnfxhvvkwpuyycqokybyqrujpdgpnpqcfrmdunejkidxpkdipigmkqwasfdewnhumokvubzqxserhpsxoskmvhsflmtvootrhpnjguqmqhpuiosqpiwmmahvuimwcquktrfnniybyhuftrfzqpmvvklgoilbwvtvaprddkwiwiezxarnxnzgqzqxhseodyyleerusznmmyxxvlmokiyhpsghcububxzrglgskrkbagamwvxxrkplpjrcsxvvvcjmjzsemvjvfmesckkrfabzfxxzmwthxldyoyhbsdsqmrugnsyracggnsextkzjqyivpiiambvsulqjefbheakvwkffcvjnuvkgusnawxdtibaycabnzeobaljpfhlhbaismpplqckycavmhttyakpngcnuawxdwwfhswyllbbhbkmuvgdu"))
	//fmt.Printf("%v\n", longestPalindrome("jsfpmgkuxqnmtruslzgyvexhqjoamvyuyedhybqqcjhhhgmwqudgldvspgugibdsqfhucpfcqzriqqusvspgbqhgkswlzdkytyqphexemwxpduxplkquvgvhefsxubjluopighxbpscekijrqjhcgmqcuoczwbvueuviyfokdoqqsckjdorsettkkpiyyxxdsfczyhkyxlvrmhvflqbvlrukqcplbxnyokdxvhubsisxrodolmpmkdczavqlsnrggffagoddaldlcexwvozjxxdjtfjrfciwpacpbajcpmgfpefngqfbzehaaqyfvthtrbhkzrzqmzdcgrkezpqgbqjembeqaziuubbvdfpfyqanilcjggkudsyigiqgrcmauyugyhepvduudvpehyguyuamcrgqigiysdukggjclinaqyfpfdvbbuuizaqebmejqbgqpzekrgcdzmqzrzkhbrthtvfyqaahezbfqgnfepfgmpcjabpcapwicfrjftjdxxjzovwxecldladdogaffggrnslqvazcdkmpmlodorxsisbuhvxdkoynxblpcqkurlvbqlfvhmrvlxykhyzcfsdxxyyipkkttesrodjkcsqqodkofyivueuvbwzcoucqmgchjqrjikecspbxhgipouljbuxsfehvgvuqklpxudpxwmexehpqytykdzlwskghqbgpsvsuqqirzqcfpcuhfqsdbigugpsvdlgduqwmghhhjcqqbyhdeyuyvmaojqhxevygzlsurtmnqxukgmpfsj"))
	//fmt.Printf("%v\n", longestPalindrome("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"))
}

func longestPalindrome(s string) string {
	start, end := 0, 0
	for i := 0; i < len(s); i++ {
		oddLeft, oddRight := checkPalindromeAtPosition(s, i, i)
		evenLeft, evenRight := checkPalindromeAtPosition(s, i, i+1)

		if oddRight-oddLeft > end-start {
			start = oddLeft
			end = oddRight
		}

		if evenRight-evenLeft > end-start {
			start = evenLeft
			end = evenRight
		}
	}
	return s[start : end+1]
}

func checkPalindromeAtPosition(s string, left, right int) (int, int) {
	resLeft, resRight := left, left
	for left >= 0 && right < len(s) && s[left] == s[right] {
		resLeft, resRight = left, right
		right++
		left--
	}
	return resLeft, resRight
}

func longestPalindromeBrute(s string) string {
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

func longestPalindromeDynamic(s string) string {
	result := ""
	cache := make(map[string]bool)

	var isPalindrome func(s string) bool

	isPalindrome = func(s string) bool {
		res := false
		cached, found := cache[s]
		if found {
			return cached
		} else if len(s) <= 1 {
			res = true
		} else {
			if s[0] == s[len(s)-1] && isPalindrome(s[1:len(s)-1]) {
				res = true
			}
		}
		cache[s] = res
		return res
	}

	n := len(s)
	for i := 0; i < n; i++ {
		for j := n; j > i; j-- {
			subString := s[i:j]
			if len(subString) > len(result) && isPalindrome(subString) {
				result = subString
			}
		}
	}

	return result
}
