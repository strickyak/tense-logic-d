// +build main

/*
	Usage:   go run main.go -n=4

	Right now the proposition is hardwired
	in the first line of main().
	TODO: Write a parser.
*/
package main

import D "github.com/strickyak/tense-logic-d"
import "flag"
import "fmt"

var N = flag.Uint("n", 4, "number of days")

func printStory(s D.Story) {
	for _, b := range s {
		ch := '-'
		if b {
			ch = 'R'
		}
		fmt.Printf(" %c", ch)
	}
	fmt.Printf("*\n")
}

func main() {
	p := D.Eventually{D.And{D.Rains{}, D.Tomorrow{D.Rains{}}}}
	flag.Parse()
	stories := D.Generate(*N)
	for i, s := range stories {
		if p.Eval(s) {
			fmt.Printf("[%5d]  ", i)
			printStory(s)
		}
	}
}
