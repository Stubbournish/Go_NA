package main

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestStrandGeneration(t *testing.T) {
	s := Strand("AAAACCCGGT")
	if s != "AAAACCCGGT" {
		t.Errorf("Strand is NOK")
	}
}

func TestStrandComplement1(t *testing.T) {
	s := Strand("AAAACCCGGT")
	if s.reverseComplement() != "ACCGGGTTTT" {
		t.Errorf("Strand is NOK. Expected %v, but returned %v", "ACCGGGTTTT", s.reverseComplement())
	}
}

func TestStrandComplement2(t *testing.T) {
	testIN, _ := ioutil.ReadFile("./TestFiles/testInStrandComplement.txt")
	testOUT, _ := ioutil.ReadFile("./TestFiles/testOutStrandComplement.txt")
	s := Strand(string(testIN))
	if s.reverseComplement() != string(testOUT) {
		t.Errorf("Strand is NOK. Expected %v, but returned %v", string(testOUT), s.reverseComplement())
	}
}

func TestSubStrand1(t *testing.T) {
	s := Strand("CTA")
	if s.subStrand(0, len(s)) != "CTA" {
		t.Errorf("subStrand is NOK. Expected %v, but returned %v", "CTA", s.subStrand(0, len(s)))
	}

}

func TestSubStrand2(t *testing.T) {
	s := Strand("CTA")
	if s.subStrand(0, len(s)-1) != "CT" {
		t.Errorf("subStrand is NOK. Expected %v, but returned %v", "CT", s.subStrand(0, len(s)-1))
	}

}

func TestSubStrand3(t *testing.T) {
	s := Strand("CTA")
	if s.subStrand(1, len(s)) != "TA" {
		t.Errorf("subStrand is NOK. Expected %v, but returned %v", "TA", s.subStrand(0, len(s)-1))
	}

}

func TestStrandPatternCount(t *testing.T) {
	s := Strand("aasdasd").patternCount("as")
	if s != 2 {
		t.Errorf("Strand method 'PatternCount is NOK.' Expected %v, but returned %v", 2, s)
	}
}

func TestStrandFrequentWords(t *testing.T) {
	s := strings.Join(Strand("ACGTTGCATGTCGCATGATGCATGAGAGCT").frequentWords(4), " ")
	if s != "GCAT CATG" {
		t.Errorf("FrequentWords is NOK. Expected %v, but returned %v", s, "CATG GCAT")
	}

}

func TestStrandFrequentWords2(t *testing.T) {
	s := strings.Join(Strand("GCGGCGAAGCATTGAAGAACAATTGGAACAATTGAGCATTGAAAGCATTGAAACAGAGTACAGAGTGAACAATTGAGCATTGAAGCGGCGAAGCATTGAAACAGAGTACAGAGTACAGAGTAGCATTGAAAGCATTGAAGCGGCGACGCGTAGGCGCGTAGGGAACAATTGAGCATTGAAAGCATTGAACGCGTAGGCGCGTAGGGCGGCGACGCGTAGGAGCATTGAAGCGGCGAAGCATTGAAACAGAGTACAGAGTAGCATTGAACGCGTAGGGAACAATTGACAGAGTGCGGCGAGCGGCGAGAACAATTGAGCATTGAAACAGAGTAGCATTGAAGCGGCGAACAGAGTAGCATTGAACGCGTAGGGAACAATTGAGCATTGAACGCGTAGGCGCGTAGGGAACAATTGAGCATTGAAGAACAATTGCGCGTAGGGCGGCGAGAACAATTGGCGGCGAACAGAGTAGCATTGAAGAACAATTGGCGGCGAACAGAGTCGCGTAGGCGCGTAGGAGCATTGAAAGCATTGAAAGCATTGAAACAGAGTAGCATTGAAGCGGCGACGCGTAGGGAACAATTGGAACAATTGGAACAATTGGAACAATTGCGCGTAGGAGCATTGAACGCGTAGGGCGGCGACGCGTAGGCGCGTAGGGAACAATTGCGCGTAGGGCGGCGAAGCATTGAAAGCATTGAACGCGTAGGGAACAATTGGAACAATTGGAACAATTGACAGAGTACAGAGTGCGGCGAACAGAGTGAACAATTGGAACAATTGGCGGCGAACAGAGTACAGAGTCGCGTAGGCGCGTAGGGAACAATTG").frequentWords(12), " ")
	o := "CGCGTAGGGAAC GCGTAGGGAACA CGTAGGGAACAA GTAGGGAACAAT TAGGGAACAATT AGGGAACAATTG"
	if s != o {
		t.Errorf("FrequentWords is NOK. Expected %v, but returned %v", o, s)
	}
}

// TODO - Complete Test function PatternIndexes
func TestStrandPatternIndexes(t *testing.T) {
	//if  {
	//		t.Errorf("")
	//	}
}
