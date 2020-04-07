package main

import (
	"flag"
	"fmt"

	euler "github.com/tlsrc/sandbox-go/euler-project"
)

type problem func() int

func main() {
	var problems = make(map[int]problem)
	problems[1] = euler.Problem0001
	problems[2] = euler.Problem0002

	var pb = flag.Int("pb", 1, "Please enter the -pb number")
	flag.Parse()

	if fn, ok := problems[*pb]; ok {
		fmt.Printf("Problem %d : %d\n", *pb, fn())
	} else {
		fmt.Printf("Problem %d not solved yet\n", *pb)
	}

}
