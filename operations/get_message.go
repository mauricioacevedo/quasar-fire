package operations

import "fmt"

// GetMessage return the composed message using the messages from the different satellites
func GetMessage(messages ...[]string) (msg string) {

	if len(messages) < 1 {
		return ""
	}

	l := len(messages[0])

	for _, m := range messages {
		if len(m) != l {
			return ""
		}
	}

	//TODO: improve for more sats
	m1 := messages[0]
	m2 := messages[1]
	m3 := messages[2]

	finalMessage := ""

	for i := 0; i < l; i++ {

		wordlist := []string{}

		wordlist = append(wordlist, m1[i])
		wordlist = append(wordlist, m2[i])
		wordlist = append(wordlist, m3[i])

		word := selectWord(wordlist)

		if word == "" {
			return ""
		}

		if finalMessage == "" {
			finalMessage = word
		} else {
			finalMessage = fmt.Sprintf("%s %s", finalMessage, word)
		}
	}

	return finalMessage
}

func selectWord(words []string) string {

	l := []WordType{}

	for _, w := range words {

		if w == "" {
			continue
		}
		val, index := findWord(w, l)

		if val {
			wt := l[index]
			wt.Count = wt.Count + 1
			l[index] = wt
		} else {
			l = append(l, WordType{Word: w, Count: 1})
		}
	}

	return GetCommonWord(l)
}

// WordType struct used to keep the frecuency of the word
type WordType struct {
	Word  string
	Count int16
}

func findWord(w string, words []WordType) (bool, int) {

	for i, ww := range words {
		if ww.Word == w {
			return true, i
		}
	}

	return false, -1
}

// GetCommonWord function to retrieve the most common word
func GetCommonWord(words []WordType) string {

	if len(words) < 1 {
		return ""
	}

	var (
		index  = -1
		bigOne = int16(-1)
	)

	for i, w := range words {
		if w.Count > bigOne {
			index = i
			bigOne = w.Count
		}
	}

	return words[index].Word
}
