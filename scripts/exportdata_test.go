package main

import (
	"github.com/thoas/go-funk"
	"testing"
)

func TestNewDataIsImported(t *testing.T) {
	wildcards, free, disposable := fetchData()

	pWildcards := readFile("../resources/wildcard-disposable-email-providers.txt")
	pFree := readFile("../resources/free-email-providers.txt")
	pDisposable := readFile("../resources/disposable-email-providers.txt")

	left, _ := funk.DifferenceString(wildcards, pWildcards)
	if len(left) > 0 {
		t.Fatal("Not all wildcards domains were added")
	}

	left, _ = funk.DifferenceString(free, pFree)
	if len(left) > 0 {
		t.Fatal("Not all free domains were added")
	}

	left, _ = funk.DifferenceString(disposable, pDisposable)
	if len(left) > 0 {
		t.Fatal("Not all wildcards were added")
	}
}
