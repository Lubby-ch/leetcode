package main

import "fmt"

// 字典树
type Node struct {
	prefix    bool
	childrens [26]*Node
}

func NewNode() *Node {
	return &Node{}
}

func (node *Node) InsertChild(str []byte) {
	if len(str) == 0 {
		node.prefix = true
		return
	}
	if node.childrens[str[0]-'a'] != nil {
		node.childrens[str[0]-'a'].InsertChild(str[1:])
		return
	}
	child := NewNode()
	node.childrens[str[0]-'a'] = child
	child.InsertChild(str[1:])
}

func PrefixMatch(word []byte, root *Node) []byte {
	var (
		i     int
		value byte
	)
	for i, value = range word {
		if root.prefix {
			if i > 0 && root.prefix {
				word = word[:i]
			}
			break
		}
		if root.childrens[value-'a'] != nil {
			root = root.childrens[value-'a']
		} else {
			if i > 0 && root.prefix {
				word = word[:i]
			}
			break
		}
	}
	return word
}

func replaceWords(dictionary []string, sentence string) string {
	root := NewNode()
	for i := range dictionary {
		root.InsertChild([]byte(dictionary[i]))
	}
	var (
		start  = int(0)
		end    = int(0)
		length = len(sentence)
		ret    []byte
	)

	for {
		if sentence[end] != ' ' {
			end++
		} else {
			ret = append(ret, PrefixMatch([]byte(sentence[start:end]), root)...)
			ret = append(ret, ' ')
			end++
			start = end
		}
		if end >= length {
			word := sentence[start:]
			ret = append(ret, PrefixMatch([]byte(word), root)...)
			return string(ret)
		}
	}
}

func replaceWords1(dictionary []string, sentence string) string {
	var (
		word    string
		dictMap = make(map[string]struct{})
	)
	for _, word = range dictionary {
		dictMap[word] = struct{}{}
	}
	var (
		start  = int(0)
		end    = int(0)
		length = len(sentence)
		ret    []byte
	)
	for {
		if end < length && sentence[end] != ' ' {
			end++
		} else {
			var (
				isfind bool
				i      int
			)
			for i = start; i < end; i++ {
				if _, ok := dictMap[sentence[start:i]]; ok {
					isfind = true
					break
				}
			}
			if isfind {
				ret = append(ret, sentence[start:i]...)
			} else {
				ret = append(ret, sentence[start:end]...)
			}
			if end == length {
				return string(ret)
			}
			ret = append(ret, ' ')
			end++
			start = end
		}
	}
}

func main() {
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
		fmt.Println(replaceWords1(testCase.dictionary, testCase.sentence))
	}
}
