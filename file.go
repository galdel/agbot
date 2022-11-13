package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Word struct {
	eng string
	rus string
}

func openFile(path string) ([]Word, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var words []Word
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var word Word
		if strings.Contains(scanner.Text(), " - ") {
			word.eng = strings.Split(scanner.Text(), " - ")[0]
			word.rus = strings.Split(scanner.Text(), " - ")[1]
			words = append(words, word)
		}
	}
	if scanner.Err() != nil {
		return nil, err
	}
	if len(words) == 0 {
		err = errors.New("Bad data!")
	}
	return words, err
}

func saveFile(words []Word, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	w := bufio.NewWriter(file)
	for _, word := range words {
		fmt.Fprintln(w, word.eng+" - "+word.rus)
	}
	return w.Flush()
}

func getPair() func() (string, string) {
	var counter int
	var first string
	var second string
	words, err := openFile("words.txt")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	//fmt.Println(words)
	return func() (string, string) {
		if rand.Intn(2) == 1 {
			first, second = words[counter].eng, words[counter].rus
		} else {
			first, second = words[counter].rus, words[counter].eng
		}
		if counter < len(words)-1 {
			//fmt.Println("Work")
			counter++
		} else {
			rand.Shuffle(len(words), func(i, j int) {
				words[i], words[j] = words[j], words[i]
			})
			counter = 0
		}
		if err != nil {
			return "", ""
		}
		//fmt.Println("Counter ", counter)
		return first, second
	}
}

func getPairs() (pairs [][2]string, err error) {
	words, err := openFile("words.txt")
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})
	//fmt.Println(words)
	for _, word := range words {
		var pair [2]string
		if rand.Intn(2) == 1 {
			pair[0], pair[1] = word.eng, word.rus
		} else {
			pair[0], pair[1] = word.rus, word.eng
		}
		pairs = append(pairs, pair)
	}
	return
}

func getLen() int {
	words, err := openFile("words.txt")
	if err != nil {
		return 0
	}
	return len(words)
}

func maino() {
	// ghjdrr
	k, _ := getPairs()
	s := getLen()
	fmt.Println(k, s)
}
