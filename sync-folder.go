package maint

import (
	"flag"
)

const DEFAULT = "##DEFAULT##"

var f1 string
var f2 string

func main() {
	flag.StringVar(&f1, "f1", DEFAULT, "Folder 1.")
	flag.StringVar(&f2, "f2", DEFAULT, "Folder 2.")
	flag.Parse()

	if f1 == DEFAULT || f2 == DEFAULT {
		panic("f1 and f2 are both required.")
	}

}
