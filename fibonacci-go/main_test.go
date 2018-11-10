package main

import "testing"

func TestFibonacciSequenceLength(t *testing.T) {
	var termsList = fibonacciSequence(6)
	var numTerms = len(termsList)

	if numTerms != 6 {
		t.Error("Expected 6, got ", numTerms)
	}
}

func TestFibonacciSequenceCorrectFirstTerm(t *testing.T) {
	var termsList = fibonacciSequence(9)
	var firstTerm = termsList[0]

	if firstTerm != 0 {
		t.Error("Expected 0, got ", firstTerm)
	}
}

func TestFibonacciSequenceCorrectLastTerm(t *testing.T) {
	var termsList = fibonacciSequence(9)
	var lastTerm = termsList[8]

	if lastTerm != 21 {
		t.Error("Expected 21, got ", lastTerm)
	}
}

func TestFibonnacciSequenceAPINegativeNumber(t *testing.T) {
	// didnt have time to get this working the way i wanted.
}

func TestFibonnacciSequenceAPITooHighNumber(t *testing.T) {
	// didnt have time to get this working the way i wanted.
}

func TestFibonnacciSequenceAPIString(t *testing.T) {
	// didnt have time to get this working the way i wanted.
}
