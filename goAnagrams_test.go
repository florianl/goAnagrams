package main

import "testing"
import "fmt"

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
