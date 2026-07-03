package exos

import (
	"sort"
	"unicode"
)

type CharCount struct {
	Char  rune
	Count int
}

// TopChars prend un texte et retourne les n caractères les plus fréquents,
// du plus fréquent au moins fréquent.
// Pattern map[count]→slice[sort]
func TopChars(text string, n int) []CharCount {

	counts := make(map[rune]int)

	// r,c = range : r et c dependent de ce sur quoi on boulce ,
	// par defaulr r= range , r forcement la clé
	for _, r := range text {
		// ignorer les espaces ici (unicode.IsSpace)
		if !unicode.IsSpace(r) {
			counts[r]++

		}
	}

	var results []CharCount
	for r, c := range counts {

		results = append(results, CharCount{r, c})
	}

	sort.Slice(results, func(i, j int) bool {

		if results[i].Count != results[j].Count {

			//Ordre decroissant
			return results[i].Count > results[j].Count

		}

		//Ordre croissant
		return results[i].Char < results[j].Char
	})

	// j'étais parti sur
	//if len(results) > n {
	//finalresult := results[:n]
	//return finalresult
	//}
	//finalresult := results[:len(results)]
	//return finalresult

	// trop long

	return results[:min(n, len(results))]

}
