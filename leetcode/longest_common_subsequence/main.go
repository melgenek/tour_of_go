package main

import "fmt"

func main() {
	fmt.Printf("%v\n", 3 == longestCommonSubsequence("ace", "abcde"))
	fmt.Printf("%v\n", 3 == longestCommonSubsequence("abcde", "ace"))
	fmt.Printf("%v\n", 3 == longestCommonSubsequence("abc", "abc"))
	fmt.Printf("%v\n", 2 == longestCommonSubsequence("aatr", "asdfa"))
	fmt.Printf("%v\n", 0 == longestCommonSubsequence("abc", "def"))
	fmt.Printf("%v\n", 6 == longestCommonSubsequence("mhunuzqrkzsnidwbun", "szulspmhwpazoxijwbq"))
	fmt.Printf("%v\n", 323 == longestCommonSubsequence("fcvafurqjylclorwfoladwfqzkbebslwnmpmlkbezkxoncvwhstwzwpqxqtyxozkpgtgtsjobujezgrkvevklmludgtyrmjaxyputqbyxqvupojutsjwlwluzsbmvyxifqtglwvcnkfsfglwjwrmtyxmdgjifyjwrsnenuvsdedsbqdovwzsdghclcdexmtsbexwrszihcpibwpidixmpmxshwzmjgtadmtkxqfkrsdqjcrmxkbkfoncrcvoxuvcdytajgfwrcxivixanuzerebuzklyhezevonqdsrkzetsrgfgxibqpmfuxcrinetyzkvudghgrytsvwzkjulmhanankxqfihenuhmfsfkfepibkjmzybmlkzozmluvybyzsleludsxkpinizoraxonmhwtkfkhudizepyzijafqlepcbihofepmjqtgrsxorunshgpazovuhktatmlcfklafivivefyfubunszyvarcrkpsnglkduzaxqrerkvcnmrurkhkpargvcxefovwtapedaluhclmzynebczodwropwdenqxmrutuhehadyfspcpuxyzodifqdqzgbwhodcjonypyjwbwxepcpujerkrelunstebopkncdazexsbezmhynizsvarafwfmnclerafejgnizcbsrcvcnwrolofyzulcxaxqjqzunedidulspslebifinqrchyvapkzmzwbwjgbyrqhqpolwjijmzyduzerqnadapudmrazmzadstozytonuzarizszubkzkhenaxivytmjqjgvgzwpgxefatetoncjgjsdilmvgtgpgbibexwnexstipkjylalqnupexytkradwxmlmhsnmzuxcdkfkxyfgrmfqtajatgjctenqhkvyrgvapctqtyrufcdobibizihuhsrsterozotytubefutaxcjarknynetipehoduxyjstufwvkvwvwnuletybmrczgtmxctuny", "nohgdazargvalupetizezqpklktojqtqdivcpsfgjopaxwbkvujilqbclehulatshehmjqhyfkpcfwxovajkvankjkvevgdovazmbgtqfwvejczsnmbchkdibstklkxarwjqbqxwvixavkhylqvghqpifijohudenozotejoxavkfkzcdqnoxydynavwdylwhatslyrwlejwdwrmpevmtwpahatwlaxmjmdgrebmfyngdcbmbgjcvqpcbadujkxaxujudmbejcrevuvcdobolcbstifedcvmngnqhudixgzktcdqngxmruhcxqxypwhahobudelivgvynefkjqdyvalmvudcdivmhghqrelurodwdsvuzmjixgdexonwjczghalsjopixsrwjixuzmjgxydqnipelgrivkzkxgjchibgnqbknstspujwdydszohqjsfuzstyjgnwhsrebmlwzkzijgnmnczmrehspihspyfedabotwvwxwpspypctizyhcxypqzctwlspszonsrmnyvmhsvqtkbyhmhwjmvazaviruzqxmbczaxmtqjexmdudypovkjklynktahupanujylylgrajozobsbwpwtohkfsxeverqxylwdwtojoxydepybavwhgdehafurqtcxqhuhkdwxkdojipolctcvcrsvczcxedglgrejerqdgrsvsxgjodajatsnixutihwpivihadqdotsvyrkxehodybapwlsjexixgponcxifijchejoxgxebmbclczqvkfuzgxsbshqvgfcraxytaxeviryhexmvqjybizivyjanwxmpojgxgbyhcruvqpafwjslkbohqlknkdqjixsfsdurgbsvclmrcrcnulinqvcdqhcvwdaxgvafwravunurqvizqtozuxinytafopmhchmxsxgfanetmdcjalmrolejidylkjktunqhkxchyjmpkvsfgnybsjedmzkrkhwryzan"))
}

func longestCommonSubsequence(text1 string, text2 string) int {
	n1 := len(text1)
	n2 := len(text2)

	cache := make([][]int, n1+1)
	for i := range cache {
		cache[i] = make([]int, n2+1)
	}
	for i := 0; i < n1; i++ {
		for j := 0; j < n2; j++ {
			cache[i][j] = -1
		}
	}

	for i := n1 - 1; i >= 0; i-- {
		for j := n2 - 1; j >= 0; j-- {
			if text1[i] == text2[j] {
				cache[i][j] = cache[i+1][j+1] + 1
			} else {
				cache[i][j] = max(
					cache[i+1][j],
					cache[i][j+1],
				)
			}
		}
	}

	return cache[0][0]
}

func longestCommonSubsequence2(text1 string, text2 string) int {
	cache := make([][]int, len(text1))
	for i := range cache {
		cache[i] = make([]int, len(text2))
		for j := range cache[i] {
			cache[i][j] = -1
		}
	}
	return longestCommonSubsequenceRec(text1, text2, 0, 0, cache)
}

func longestCommonSubsequenceRec(text1 string, text2 string, i1, i2 int, cache [][]int) int {
	if i1 == len(text1) || i2 == len(text2) {
		return 0
	} else if cached := cache[i1][i2]; cached >= 0 {
		return cached
	} else {
		var res int
		if text1[i1] == text2[i2] {
			res = longestCommonSubsequenceRec(text1, text2, i1+1, i2+1, cache) + 1
		} else {
			res = max(
				longestCommonSubsequenceRec(text1, text2, i1+1, i2, cache),
				longestCommonSubsequenceRec(text1, text2, i1, i2+1, cache),
			)
		}
		cache[i1][i2] = res
		return res
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
