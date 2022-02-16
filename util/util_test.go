package util

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	x := "AbbcasdsdSDAFAS"
	XX := "2sdjfklsadh"
	fmt.Printf("CheckWords(x): %v\n", CheckWords(x))
	fmt.Printf("CheckWords(XX): %v\n", CheckWords(XX))
}
