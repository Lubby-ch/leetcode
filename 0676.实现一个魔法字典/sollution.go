package main

import "fmt"

type node struct {
	isend    bool
	children [26]*node
}

func NewNode() *node {
	return &node{}
}

type MagicDictionary struct {
	root *node
}

func Constructor() MagicDictionary {
	return MagicDictionary{root: NewNode()}
}

func (this *MagicDictionary) BuildDict(dictionary []string) {
	for _, word := range dictionary {
		this.root.insertWord(word)
	}
}

func (this *MagicDictionary) Search(searchWord string) bool {
	var isfind =  dnf(this.root, searchWord, false)
	this.root.insertWord(searchWord)
	return isfind
}

func (this *node) insertWord(word string) {
	if len(word) == 0 {
		this.isend = true
		return
	}
	head := word[0]
	if this.children[head-'a'] == nil {
		this.children[head-'a'] = NewNode()
	}
	this.children[head-'a'].insertWord(word[1:])
}

func dnf(this *node, word string, modified bool) bool {
	if len(word) == 0 {
		return this.isend && modified
	}
	head := word[0] - 'a'
	if this.children[head] != nil {
		if dnf(this.children[head], word[1:], modified) {
			return true
		}
	}
	if !modified {
		for i := range this.children {
			if this.children[i] == nil || i == int(head) {
				continue
			}
			if dnf(this.children[i], word[1:], true) {
				return true
			}
		}
	}
	return false
}

func main() {
	testCases := []struct {
		Dictionary []string
		SearchWord []string
	}{
		{
			Dictionary: []string{"magic", "dictionary", "builddict", "search"},
			SearchWord: []string{"hello", "hhllo"},
		},
	}
	dictionary := Constructor()
	for _, testCase := range testCases {
		dictionary.BuildDict(testCase.Dictionary)
		for _, word := range testCase.SearchWord {
			fmt.Println(dictionary.Search(word))
		}
	}
}
