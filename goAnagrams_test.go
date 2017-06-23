package main

import "testing"
import "errors"

func TestCreatePrimeMap(t *testing.T) {
	tests := []struct {
		word   []rune
		primes map[int32]int64
		err    error
	}{
		{[]rune{'A', 'n', 'n', 'a'}, map[int32]int64{65: 2, 110: 3, 97: 5}, nil},
		{[]rune{'世', '界'}, map[int32]int64{19990: 2, 30028: 3}, nil},
		{[]rune{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}, map[int32]int64{72: 2, 101: 3, 111: 7, 87: 13, 114: 17, 100: 19, 108: 5, 32: 11, 33: 23}, nil},
	}

	for _, test := range tests {
		primes, err := createPrimeMap(test.word)
		if err != test.err {
			t.Errorf("Expected: %v \t Got: %v", test.err, err)
		}
		for i, prime := range primes {
			if test.primes[i] != prime {
				t.Errorf("Expected: %d \t Got: %d", prime, primes[i])
			}
		}
	}

}

func TestCreateNumber(t *testing.T) {
	tests := []struct {
		word    []rune
		mapping map[int32]int64
		num     int64
		err     error
	}{
		{[]rune{'A', 'n', 'n', 'a'}, map[int32]int64{65: 2, 110: 3, 97: 5}, 90, nil},
		{[]rune{'a', 'n', 'n', 'a'}, map[int32]int64{110: 2, 97: 3}, 36, nil},
		{[]rune{'世', '界'}, map[int32]int64{19990: 2, 30028: 3}, 6, nil},
		{[]rune{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!'}, map[int32]int64{72: 2, 101: 3, 111: 7, 87: 13, 114: 17, 100: 19, 108: 5, 32: 11, 33: 23}, 39041252250, nil},
		{[]rune{'a', 'b', 'c'}, map[int32]int64{97: 2, 98: 3}, -1, errors.New("\"c\" is missing in mapping")},
	}

	for _, test := range tests {
		num, err := createNumber(test.word, test.mapping)
		if err != nil && test.err != nil {
			t.Log("Expected: ", test.err, "\t Got: ", err)
		} else if err != nil && test.err == nil {

			t.Errorf("Expected: %v \t Got: %v", test.err, err)
		}
		if num != test.num {
			t.Errorf("Expected: %v \t Got: %v", test.num, num)
		}
	}
}
