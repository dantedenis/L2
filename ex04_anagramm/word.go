package ex04_anagramm

import (
	"golang.org/x/exp/slices"
	"sort"
	"strings"
)

// wordPair low - lowercase wordPair; hash - pseudoHash of wordPair
type wordPair struct {
	low  string
	hash string
}

type Set struct {
	set []wordPair
}

func NewSet() *Set {
	return &Set{set: make([]wordPair, 0)}
}

func (s *Set) Add(words *[]string) {
	for _, str := range *words {
		str = strings.ToLower(str)
		runes := sliceRune(StringToRune(str))
		sort.Sort(runes)
		s.set = append(s.set, wordPair{low: str, hash: string(runes)})
	}
}

func (s *Set) ToMap() *map[string]*[]string {
	preResult, result := make(map[string][]string), make(map[string]*[]string)

	for _, s := range s.set {
		if !slices.Contains(preResult[s.hash], s.low) {
			preResult[s.hash] = append(preResult[s.hash], s.low)
		}
	}

	for _, v := range preResult {
		if len(v) != 1 {
			sort.Strings(v[1:])
			temp := v[1:]
			result[v[0]] = &temp
		}
	}
	return &result
}
