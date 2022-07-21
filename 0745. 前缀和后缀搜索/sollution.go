package main

import (
	"fmt"
)

type Node struct {
	index    int
	children [26]*Node
}

func NewNode() *Node {
	return &Node{
		index: -1,
	}
}

func (node *Node) insert(words []string) {
	for i := range words {
		word := words[i]
		n := node
		for {
			if len(word) == 0 {
				n.index = i
				break
			}
			head := word[0] - 'a'
			word = word[1:]
			if n.children[head] == nil {
				n.children[head] = NewNode()
			}
			n = n.children[head]
		}
	}
}

func (node *Node) insertWord(word string, index int) {
	if len(word) == 0 {
		node.index = index
		return
	}
	head := word[0] - 'a'
	if node.children[head] == nil {
		node.children[head] = NewNode()
	}
	node.children[head].insertWord(word[1:], index)
}

func (node *Node) searchPref(pref string) *Node {
	if len(pref) == 0 {
		return node
	}
	head := pref[0] - 'a'
	if node.children[head] == nil {
		return nil
	}
	return node.children[head].searchPref(pref[1:])
}

func dnf(node *Node, word, suff string) int {
	var ret = -1
	if node.index != -1 && len(word)-len(suff) >= 0 && word[len(word)-len(suff):] == suff {
		ret = node.index
	}
	for i, child := range node.children {
		if child == nil {
			continue
		}
		if tmp := dnf(child, word+string([]byte{'a' + byte(i)}), suff); ret < tmp {
			ret = tmp
		}
	}
	return ret
}

type WordFilter struct {
	root *Node
}

func Constructor(words []string) WordFilter {
	wordFilter := WordFilter{
		root: NewNode(),
	}
	wordFilter.root.insert(words)
	return wordFilter
}

func (this *WordFilter) F(pref string, suff string) int {
	node := this.root.searchPref(pref)
	if node == nil {
		return -1
	}
	return dnf(node, pref, suff)
}

func main() {
	testCases := []struct {
		words []string
		pref  string
		suff  string
	}{
		{
			words: []string{"aaaaa", "apple"},
			pref:  "a",
			suff:  "e",
		},
	}
	for _, testCase := range testCases {
		wordFilter := Constructor(testCase.words)
		fmt.Println(wordFilter.F(testCase.pref, testCase.suff))
	}
}
