package main

import (
	"fmt"
	"strings"
)

// Strand representing DNA
type Strand string

func getNucleotideComplement(nucleotide string) string {
	// If something is wrong, then is expected to Return a field "E" for traceability purposes
	switch nucleotide {
	case "A":
		return "T"
	case "T":
		return "A"
	case "G":
		return "C"
	case "C":
		return "G"
	}
	return "E"

}

// Strand
func (s Strand) print() {
	fmt.Println(s)
}

func (s Strand) reverseComplement() string {
	r := strings.Split(reverse(string(s)), "")
	for i, nucl := range r {
		r[i] = getNucleotideComplement(nucl)
	}

	return strings.Join(r, "")
}

func (s Strand) subStrand(start int, finish int) Strand {
	return s[start:finish]
}
