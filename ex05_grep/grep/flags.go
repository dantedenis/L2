package grep

import (
	"flag"
)

type flags struct {
	aft, bf, ctx  int
	count, ignore, invert, fix, num bool
	pattern string
}

func NewFlags() *flags {
	var f flags
	
	flag.IntVar(&f.aft, "A", 0, "\"After\" печать +N строк после совпадения")
	flag.IntVar(&f.bf, "B", 0, "\"Before\" печать +N строк до совпадения")
	flag.IntVar(&f.ctx, "C", 0, "\"Context\" печать N строк вокруг совпадения")
	flag.BoolVar(&f.count, "c", false, "\"Count\" количество строк")
	flag.BoolVar(&f.ignore, "i", false, "\"Ignore-case\" игнорировать регистр")
	flag.BoolVar(&f.invert, "v", false, "\"Invert\" исключать совпадения")
	flag.BoolVar(&f.fix, "F", false, "\"Fixed\" точное совпадение со строкой")
	flag.BoolVar(&f.num, "n", false, "\"Line num\" печатать номера строк")
	
	flag.Parse()
	
	f.pattern = flag.Arg(0)
	
	return &f
}

