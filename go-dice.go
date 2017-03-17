package main

import (
	"bufio"
	"crypto/rand"
	"flag"
	"fmt"
	"math"
	"math/big"
	"os"
	"regexp"
	"strings"
)

func read_words(word_file string, min_length, max_length int) []string {
	// words list has all sorts of weird stuff.  filter out the stuff
	// with special characters
	re_word, _ := regexp.Compile("^[a-z]+$")
	words := make([]string, 0, 20000)
	file, err := os.Open(word_file)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// convert all words to lowercase
		word := strings.ToLower(scanner.Text())
		// only add words with valid char and min_length/max_length to word_list
		if len(word) >= min_length && len(word) <= max_length && re_word.MatchString(word) {
			words = append(words, word)
		}
	}
	return words
}

func random_words(word_list []string, count int) []string {
	// maximum index for randomly selected word
	max_index := big.NewInt(int64(len(word_list)))
	words := make([]string, count)
	for i := 0; i < count; i++ {
		index, err := rand.Int(rand.Reader, max_index)
		if err != nil {
			panic(err)
		}
		words[i] = word_list[index.Int64()]
	}
	return words
}

func main() {

	var count int
	var min_length int
	var max_length int
	var word_file string
	var quiet bool

	flag.IntVar(&count, "c", 5, "number of words to use")
	flag.IntVar(&min_length, "n", 6, "minimum word length")
	flag.IntVar(&max_length, "m", 15, "maximum word length")
	flag.StringVar(&word_file, "w", "/usr/share/dict/words", "word file")
	flag.BoolVar(&quiet, "q", false, "quiet mode")

	flag.Parse()

	word_list := read_words(word_file, min_length, max_length)
	word_list_length := len(word_list)

	if word_list_length == 0 {
		panic("empty word list")
	}

	words := random_words(word_list, count)

	// print out some stats about word list and entropy
	if !quiet {
		entropy := float64(count) * math.Log2(float64(len(word_list)))
		fmt.Printf("Total word list size: %v (%0.2f bits of entropy)\n\n", len(word_list), entropy)
	}

	// print out selected word list
	fmt.Println(strings.Join(words, " "))
}
