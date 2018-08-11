package nregexp_test

import (
	"fmt"
	"regexp"

	"github.com/mdouchement/nregexp"
)

func ExampleNamedCaptureString() {
	re := regexp.MustCompile(`.*\s.*\s(?P<my_capture>[^\s]*)\s.*`)
	matches := re.FindStringSubmatch("Example of nregexp package")
	if len(matches) <= 1 {
		fmt.Println("No match")
	}
	fmt.Printf("%v\n", nregexp.NamedCaptureString(re, matches))
	// Output: map[my_capture:nregexp]
}

func ExampleNRegexp_NamedCaptureString() {
	re := nregexp.MustCompile(`.*\s.*\s(?P<my_capture>[^\s]*)\s.*`)
	fmt.Printf("%v\n", re.NamedCaptureString("Example of nregexp package"))
	// Output: map[my_capture:nregexp]
}
