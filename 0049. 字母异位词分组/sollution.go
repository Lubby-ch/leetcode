package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	var assistMap = make(map[string][]string)
	for _, str := range strs {
		bs := []byte(str)
		sort.Slice(bs, func(i, j int) bool { return bs[i] < bs[j] })
		sstr := string(bs)
		assistMap[sstr] = append(assistMap[sstr], str)
	}
	var ans = make([][]string, 0, len(assistMap))
	for _, v := range assistMap {
		ans = append(ans, v)
	}
	return ans
}
