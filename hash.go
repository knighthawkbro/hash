package main

import (
	"bufio"
	"fmt"
	"hash/hashset"
	"log"
	"os"
	"regexp"
	"strings"
)

// hash (Private) - Interface to the hashset datastructure
type hash interface {
	Add(item string) bool
	RemoveItem(item string) bool
	Contains(item string) bool
	Remove() string
	Get() string
	Items() int
	Size() int
	String() string
}

// main (private) - Main function for all the code to run from, un
func main() {
	hash := hashset.New()
	fmt.Println("\n******************************************")
	fmt.Print("*\tRunning driver function...")
	fmt.Println("\n******************************************")
	fmt.Println("")
	driver(hash)
	hash = new(hashset.HashSet).Init(51)
	fmt.Println("\n******************************************")
	fmt.Print("*\tRunning seuss function...")
	fmt.Println("\n******************************************")
	fmt.Println("")
	seuss(hash)
	hash = loadDictionary()
	test := loadTest()
	//fmt.Println(hash.Items())
	fmt.Println("\n******************************************")
	fmt.Print("*\tRunning spell function...")
	fmt.Println("\n******************************************")
	fmt.Println("")
	for _, word := range test {
		if spell(word, hash) {
			fmt.Println(word)
		}
	}
}

// loadTest (Private) - takes the test.txt file and returns a slice
func loadTest() []string {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	words := []string{}
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			word = reg.ReplaceAllString(word, "")
			if word == "" {
				continue
			}
			words = append(words, strings.ToLower(word))
		}
	}
	return words
}

// spell (Private) - checks to see if the word is in the dictionary provided
func spell(word string, dictionary *hashset.HashSet) bool {
	if !dictionary.Contains(word) {
		return true
	}
	return false
}

// loadDictionary (Private) - loads the dictionary of choice and returns
// a hashset of all the words for fast lookup near O(1)
func loadDictionary() *hashset.HashSet {
	file, err := os.Open("dictionary.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	result := new(hashset.HashSet).Init(10000)
	for scanner.Scan() {
		word := scanner.Text()
		word = reg.ReplaceAllString(word, "")
		if word == "" {
			continue
		}
		result.Add(word)
	}
	return result
}

// seuss (Private) - Loads all the distinct words in Green Eggs and Ham
// then prints out the has table with all the words
func seuss(words hash) {
	file, err := os.Open("greenEggs.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	reg, err := regexp.Compile("[^a-zA-Z]+")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for _, word := range strings.Split(scanner.Text(), " ") {
			word = reg.ReplaceAllString(word, "")
			if !words.Contains(strings.ToLower(word)) {
				words.Add(strings.ToLower(word))
			}
		}
	}
	fmt.Println(words)
}

// driver (Private) - Program to test out all hashset functions
func driver(letters hash) {
	toAdd := []string{
		"C", "b", "c", "Q", "z",
		"D", "E", "P", "j", "F",
		"E", "I", "y", "f", "i",
		"U", "V", "m", "e", "W",
	}

	for _, letter := range toAdd {
		letters.Add(letter)
	}

	fmt.Println("Printing letters")
	fmt.Println(letters)

	fmt.Print("Does the set contain 'e'? ")
	if letters.Contains("e") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Print("Does the set contain 'a'? ")
	if letters.Contains("a") {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	fmt.Println("After removing e, P, i, D, and i:")
	letters.RemoveItem("e")
	letters.RemoveItem("P")
	letters.RemoveItem("i")
	letters.RemoveItem("D")
	letters.RemoveItem("i")
	fmt.Println(letters)
}
