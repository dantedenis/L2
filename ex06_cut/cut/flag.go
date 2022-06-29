package cut

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type flags struct {
	f, d, input string
	s           bool
	targetRow   []int
}

func ParseFlags() *flags {
	var f flags
	flag.StringVar(&f.f, "f", "", "space (column's)")
	flag.StringVar(&f.d, "d", "\t", "delim")
	flag.BoolVar(&f.s, "s", false, "with delim")

	flag.Parse()
	f.input = flag.Arg(0)

	if f.f != "" {
		f.targetRow = parseColumn(f.f)
	}

	return &f
}

func (f flags) String() string {
	return fmt.Sprintf("Flags: \n -f: %s\n -d: %s\n -s: %t\n -input:%s\n -range:%v", f.f, f.d, f.s, f.input, f.targetRow)
}

func parseColumn(str string) []int {
	rows := strings.Split(str, ",")
	result := make([]int, 0, len(rows))

	for _, r := range rows {
		pair := strings.Split(r, "-")

		if len(pair) == 2 {
			num1, _ := strconv.Atoi(pair[0])
			num2, _ := strconv.Atoi(pair[1])
			fmt.Println(num1, num2)
			for num1 <= num2 {
				result = append(result, num1)
				num1++
			}
		} else {
			num, _ := strconv.Atoi(pair[0])
			result = append(result, num)
		}
	}

	sort.Ints(result)
	return result
}
