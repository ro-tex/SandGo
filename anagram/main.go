package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const wordListUrl = "http://codekata.com/data/wordlist.txt"
const wordListFile = "./anagram/wordList.dat"

var wordList map[string]interface{}

func init() {
	wordList = loadWords()
}

func fetchWordList() ([]byte, error) {
	resp, err := http.Get(wordListUrl)
	if err != nil {
		return []byte{}, err
	}

	defer func() { _ = resp.Body.Close() }()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	err = ioutil.WriteFile(wordListFile, body, 0644)
	if err != nil {
		println("ERROR writing file:", err.Error())
	}

	return body, nil
}

func loadWords() map[string]interface{} {
	var wordList = make(map[string]interface{})

	data, err := ioutil.ReadFile(wordListFile)
	if err != nil || len(data) == 0 {
		data, err = fetchWordList()
		if err != nil {
			return map[string]interface{}{}
		}
	}

	for _, v := range strings.Split(string(data), "\n") {
		wordList[v] = struct{}{}
	}
	return wordList
}

//func findAnagrams(word string) []string {
//
//}

func generatePermutations(runes []rune) [][]rune {
	var permutations [][]rune

	for i := range runes {
		subrunes := append(runes[:i], runes[i+1:]...)
		fmt.Println("entering i loop", string(runes[i]), "subrunes", subrunes)
		if len(subrunes) == 0 {
			permutations = append(permutations, []rune{runes[i]})
		} else {
			perms := generatePermutations(subrunes)
			for j, p := range perms {
				fmt.Println("entering j loop", string(perms[j]), string(p))
				fmt.Println("j", j, perms[j])
				pr := append([]rune{runes[i]}, perms[j]...)
				fmt.Println("p", permutations, pr)
				permutations = append(permutations, pr)
				fmt.Println("permutations", permutations)
			}
		}

	}
	//fmt.Println("will return ", permutations, " for ", string(runes))
	return permutations
}

func p(runewords [][]rune) {
	var words []string
	for i := range runewords {
		words = append(words, string(runewords[i]))
	}
	fmt.Println(words)
}

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	var anagrams []string
	//for scanner.Scan() {
	//word := scanner.Text()
	perms := generatePermutations([]rune("ab"))

	p(perms)

	//anagrams = []string{}
	//var anagram string
	//for i := range perms {
	//	anagram = perms[i]
	//	fmt.Println("checking for", anagram)
	//	if wordList[anagram] != nil {
	//		anagrams = append(anagrams, anagram)
	//	}
	//}
	//}

	fmt.Println(anagrams)
}
