package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/satori/go.uuid"
)

var (
	noDash = flag.Bool("n", false, "remove dash from result")
	upper  = flag.Bool("u", false, "uppercase result")
)

func main() {
	flag.Parse()

	u := uuid.Must(uuid.NewV4())
	r := u.String()
	if *noDash {
		r = strings.Replace(r, "-", "", -1)
	}
	if *upper {
		r = strings.ToUpper(r)
	}
	fmt.Println(r)
}
