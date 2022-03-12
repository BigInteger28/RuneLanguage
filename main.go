package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"unicode/utf8"
)

//CONSTANTS after import
var Runes []rune
var LanguageFile []string

var translation []rune

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func importRunes() {
	content, err := ioutil.ReadFile("alfabet.txt")
	check(err)
	for i := range content {
		Runes = append(Runes, rune(content[i]))
	}
}

func importLanguage() {
	file, err := os.Open("woordenboek.txt")
	check(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		LanguageFile = append(LanguageFile, scanner.Text())
	}
}

func writeToFile() {
	//file, err := os.Create("vertaling.txt")
	//check(err)

}

func positionWord(word string) int {
	var position int = -1
	for currentWord := 0; currentWord < len(LanguageFile); currentWord++ {
		if word == LanguageFile[currentWord] {
			position = currentWord
			goto result
		}
	}
result:
	return position
}

func translateWordToRunes(position int) {
	amountRunes := len(Runes)
	if position < amountRunes {

	}
}

func translateToDutch() {
	const nihongo = "日本語"
	for i, w := 0, 0; i < len(nihongo); i += w {
		runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
		fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
		w = width
	}
}

func main() {

}
