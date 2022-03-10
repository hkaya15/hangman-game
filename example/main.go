package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var words = []string{"masa", "sandalye", "elma", "almanya", "muz"}
var stateValue=3
const HANGMAN=9


func main() {
	randomWord := selectRandom()
	guessCount, encryptedWord := calculateAndEncrypt(randomWord)
	getGuessFromUser(guessCount, encryptedWord, randomWord)
}

func selectRandom() string {
	t := time.Now()
	rand.Seed(t.UnixNano())
	index := rand.Intn(len(words))
	word := words[index]
	return word
}

func calculateAndEncrypt(word string) (int, []string) {
	guessCount := HANGMAN-stateValue
	encryptedWord := []string{}
	for i := 0; i < len(word); i++ {
		encryptedWord = append(encryptedWord, "_")
	}
	return guessCount, encryptedWord
}

func getGuessFromUser(guessCount int, encryptedWord []string, word string) {
	for guessCount > 0 {
		fmt.Println(showState(stateValue))
		fmt.Println(guessCount, "hakkınız kalmıştır")
		letter, err := getLetter(encryptedWord)
		if err != nil {
			fmt.Println("Hata okunuyor!")
			return
		}
		if !checkMatches(word, []string{letter}) {
			guessCount--
			stateValue++
		}
		if updateEncryption(encryptedWord, word, letter) {
			fmt.Println("Kazandın! Kelime: ", word)
			return
		}
	}
	fmt.Println(showState(stateValue))
	fmt.Println("Kaybettin! Kelime: ", word)

}

func getLetter(encryptedWord []string) (string, error) {
	alphabet := "abcçdefgğhıijklmnoöprstuüvyz"
	for 1 > 0 {
		letter, err := getInput("Harf giriniz: ", strings.Join(encryptedWord, " "))
		letter=strings.ToLower(letter)
		if err != nil {
			return "", err
		}
		if len(letter) == 1 && checkMatches(alphabet, []string{letter}) {
			return letter, nil
		}
	}
	return "", nil
}

func checkMatches(s string, chars []string) bool {
	for _, ch := range s {
		for _, ch2 := range chars {
			if string(ch) == ch2 {
				return true
			}
		}
	}
	return false
}

func getInput(vals ...interface{}) (string, error) {
	if len(vals) != 0 {
		fmt.Println(vals...)
	}
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return "", err
	}
	word := scanner.Text()
	return word, nil
}

func updateEncryption(encryptedWord []string, word string, letter string) bool {
	complete := true
	for i, v := range word {
		if letter == string(v) {
			encryptedWord[i] = letter
		}
		if encryptedWord[i] == "_" {
			complete = false
		}
	}
	return complete
}

func showState(s int)string{
	stateFile:= fmt.Sprintf("./states/hangman%d",s)
	readed, err := os.ReadFile(stateFile)
	if err !=nil{
		fmt.Println(err)
	}
	state:=string(readed)
	positionHangman:="Adam Asılıyor"
	if s == HANGMAN{
		positionHangman="Adam Asıldı"
	}
	return fmt.Sprintf("\n%s\n%s",positionHangman,state)
}
