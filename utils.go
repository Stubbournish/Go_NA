package main

// keys will be k-mers, values will be count of frequency
type DnaFrequencyDict map[string]int

func (d DnaFrequencyDict) GenerateFreqDict(k_mer string) DnaFrequencyDict {
	return DnaFrequencyDict{}
}

func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// TODO - Develop function pattern to number to help making que frequecy array

// TODO - Develop function number to pattern to help making que frequecy array
