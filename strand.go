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

func (s Strand) patternCount(pattern string) int {
	count := 0
	for i := 0; i <= len(s)-len(pattern); i++ {
		section := s[i : len(pattern)+i]
		if pattern == string(section) {
			count++
		}
	}
	return count
}

func (s Strand) frequentWords(k int) []string {
	frequentPatterns := make([]string, 0)
	DnaFrequencyDict := make(DnaFrequencyDict, 0)
	maxCount := 0

	for i := 0; i <= len(s)-k; i++ {
		pattern := string(s)[i : k+i]
		DnaFrequencyDict[pattern]++
		if DnaFrequencyDict[pattern] > maxCount {
			maxCount = DnaFrequencyDict[pattern]
			frequentPatterns = nil
			frequentPatterns = append(frequentPatterns, pattern)
		} else if DnaFrequencyDict[pattern] == maxCount && !stringInSlice(pattern, frequentPatterns) {
			frequentPatterns = append(frequentPatterns, pattern)
		}

	}
	return frequentPatterns
}
