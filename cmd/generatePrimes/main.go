package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/golang-collections/go-datastructures/bitarray"
)

func main() {
	var limit uint64 = 100
	numsPerLine := uint64(1000)
	reverse := false
	if len(os.Args) > 1 {
		numStr := os.Args[1]
		base := 10
		hexPrefix := "0x"
		octPrefix := "0"
		if strings.HasPrefix(numStr, hexPrefix) {
			base = 16
			numStr = numStr[len(hexPrefix):]
		} else if strings.HasPrefix(numStr, octPrefix) {
			base = 8
			numStr = numStr[len(octPrefix):]
		}
		n, err := strconv.ParseUint(numStr, base, 64)

		if err == nil && n > 0 {
			limit = n
		}
		if len(os.Args) > 2 {
			reverse = os.Args[2] == "reverse"
			if len(os.Args) > 3 {
				tmp, err := strconv.ParseUint(numStr, 10, 32)
				if err == nil {
					numsPerLine = tmp
				}
			}
		}

	}

	SieveOfEratosthenes(limit, reverse, numsPerLine)
}

// Based off https://siongui.github.io/2017/04/17/go-sieve-of-eratosthenes/
func SieveOfEratosthenes(n uint64, isReverse bool, numsPerLine uint64) {

	// Create a boolean bitarray "prime[0..n]" and initialize
	// all entries it as true. A value in prime[i] will
	// finally be false if i is Not a prime, else true.

	bools := bitarray.NewBitArray(n, true)

	for p := uint64(2); p*p <= n; p++ {
		// if bools.getBit(p) is true then p is a prime
		b, e := bools.GetBit(p)
		if b == true && e == nil {
			// Update all multiples of p
			for i := p * 2; i <= n; i += p {
				bools.ClearBit(i)
			}
		}
	}

	if isReverse {
		count := uint64(0)
		for p := uint64(n); p > 1; p-- {
			b, e := bools.GetBit(p)
			if b == true && e == nil {
				count = count + 1
				ending := ","
				if count%numsPerLine == 0 {
					ending = "\n"
				}
				fmt.Printf("%d%s", p, ending)
			}
		}
	} else {
		// print all prime numbers <= n
		count := uint64(0)
		for p := uint64(2); p <= n; p++ {
			b, e := bools.GetBit(p)
			if b == true && e == nil {
				count = count + 1
				ending := ","
				if count%numsPerLine == 0 {
					ending = "\n"
				}
				fmt.Printf("%d%s", p, ending)
			}
		}
	}

}
