package main

import (
	"testing"
)

func BenchmarkFib(b *testing.B) {
	testCases := []struct {
		dictionary []string
		sentence   string
	}{
		{
			dictionary: []string{"cat", "bat", "rat"},
			sentence:   "by the battery",
		},
		{
			dictionary: []string{"a", "aa", "aaa", "aaaa"},
			sentence:   "baba ababa",
		},
		{
			dictionary: []string{"a", "b", "c"},
			sentence:   "aadsfasf absbs bbab cadsfafs",
		},
	}
	for i := range testCases {
		testCase := testCases[i]
		replaceWords(testCase.dictionary, testCase.sentence)
	}
}
