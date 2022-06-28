package pkg

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Sort struct {
	rule *flags
	text [][]string
}

type SliceWords [][]string

func NewSort() (*Sort, error) {
	f, err := ParserFlags()
	if err != nil {
		return nil, errors.New(fmt.Sprintf("sort: %s", err.Error()))
	}

	content, err := ioutil.ReadFile(f.input)
	if err != nil {
		return nil, err
	}

	tmp := string(content)
	tmp = strings.Trim(tmp, "\n")
	ss := make([][]string, 0)

	splitText := strings.Split(tmp, "\n")
	if f.GetU() {
		splitText = unique(splitText)
	}

	for _, t := range splitText {
		ss = append(ss, strings.Split(t, " "))
	}
	result := &Sort{
		rule: f,
		text: ss,
	}
	return result, nil
}

func (s *Sort) Run() {
	var sortFunc func(i, j int) bool

	switch true {
	case s.rule.GetN():
		sortFunc = func(i, j int) bool {
			a, _ := strconv.ParseFloat(getData(s.text, i, s.rule.GetK()), 64)
			b, _ := strconv.ParseFloat(getData(s.text, j, s.rule.GetK()), 64)
			if s.rule.GetR() {
				return a > b
			}
			return a < b
		}
	default:
		sortFunc = func(i, j int) bool {
			if s.rule.GetR() {
				return getData(s.text, i, s.rule.GetK()) > getData(s.text, j, s.rule.GetK())
			}
			return getData(s.text, i, s.rule.GetK()) < getData(s.text, j, s.rule.GetK())
		}
	}

	sort.Slice(s.text, sortFunc)
}

func getData(data [][]string, i, k int) string {
	if k < len(data[i]) {
		return data[i][k]
	}
	return ""
}

func unique(text []string) []string {
	hash := make(map[string]bool)
	result := make([]string, 0)
	for _, line := range text {
		if _, ok := hash[line]; !ok {
			result = append(result, line)
			hash[line] = true
		}
	}
	return result
}

func (s *Sort) Write() (n int, err error) {
	f, err := os.Create(s.rule.output)
	if err != nil {
		return
	}
	defer func(f *os.File) {
		err = f.Close()
	}(f)
	for _, line := range s.text {
		var temp int
		temp, err = fmt.Fprintf(f, strings.Join(line, " "))
		if err != nil {
			return
		}
		n += temp
	}
	return
}
