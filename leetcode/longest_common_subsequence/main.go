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
	cache := make(map[string]int)
	return longestCommonSubsequenceRec(text1, text2, len(text1)-1, len(text2)-1, cache)
}

func longestCommonSubsequenceRec(text1 string, text2 string, i int, j int, cache map[string]int) int {
	key := fmt.Sprintf("%d_%d", i, j)
	if i == -1 || j == -1 {
		return 0
	} else if cached, found := cache[key]; found {
		return cached
	} else if text1[i] == text2[j] {
		res := longestCommonSubsequenceRec(text1, text2, i-1, j-1, cache) + 1
		cache[key] = res
		return res
	} else {
		res := max(
			longestCommonSubsequenceRec(text1, text2, i-1, j, cache),
			longestCommonSubsequenceRec(text1, text2, i, j-1, cache),
		)
		cache[key] = res
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
