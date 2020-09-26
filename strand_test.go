package main

import (
	"io/ioutil"
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
