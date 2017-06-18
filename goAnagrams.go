package main

import (
	"fmt"
	"math/big"
	"os"
)

func createNumber(word []rune, mapping map[int32]int64) (int64, error) {
	var number int64 = 1

	for _, i := range word {
		if _, ok := mapping[i]; ok {
			number *= mapping[i]
		} else {
			return -1, fmt.Errorf("\"%v\" is missing in mapping", string(i))
		}
	}
	return number, nil
}

func createPrimeMap(word []rune) (map[int32]int64, error) {
	mapping := make(map[int32]int64)
	var number int64 = 2

	for _, i := range word {
		if _, ok := mapping[i]; ok {
			continue
		}
		for {
			if big.NewInt(number).ProbablyPrime(31) == true {
				mapping[i] = number
				number += 1
				break
			}
			number += 1
		}
	}

	return mapping, nil
}

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Incorrect number of arguments")
		os.Exit(-1)
	}

	word1 := []rune(os.Args[1])
	word2 := []rune(os.Args[2])

	if len(word1) != len(word2) {
		fmt.Println(string(word1), "and", string(word2), "are not anagrams")
		os.Exit(1)
	}

	primeMap, err := createPrimeMap(word1)
	if err != nil {
		fmt.Errorf("Could not create mapping")
		os.Exit(-1)
	}

	num1, err := createNumber(word1, primeMap)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	num2, err := createNumber(word2, primeMap)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	if num1 == num2 {
		fmt.Println(string(word1), "and", string(word2), "are probably anagrams")
	} else {
		fmt.Println(string(word1), "and", string(word2), "are not anagrams")
	}
}
