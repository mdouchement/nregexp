package nregexp

import "regexp"

// NamedCapture converts all matches to named matches.
func NamedCapture(re *regexp.Regexp, matches [][]byte) map[string][]byte {
	result := make(map[string][]byte)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = matches[i]
		}
	}
	return result
}

// NamedCaptureString converts all matches to named matches.
func NamedCaptureString(re *regexp.Regexp, matches []string) map[string]string {
	result := make(map[string]string)
	for i, name := range re.SubexpNames() {
		if i != 0 {
			result[name] = matches[i]
		}
	}
	return result
}

// A NRegexp is a wrapper around regexp.Regexp that allows to use named capture.
// All the methods of regexp.Regexp are exposed.
type NRegexp struct {
	*regexp.Regexp
}

// Compile parses a regular expression and returns, if successful,
// a NRegexp object that can be used to match against text.
//
// When matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses the one that a backtracking search would have found first.
// This so-called leftmost-first matching is the same semantics
// that Perl, Python, and other implementations use, although this
// package implements it without the expense of backtracking.
// For POSIX leftmost-longest matching, see CompilePOSIX.
func Compile(expr string) (*NRegexp, error) {
	re, err := regexp.Compile(expr)
	return &NRegexp{Regexp: re}, err
}

// CompilePOSIX is like Compile but restricts the regular expression
// to POSIX ERE (egrep) syntax and changes the match semantics to
// leftmost-longest.
//
// That is, when matching against text, the regexp returns a match that
// begins as early as possible in the input (leftmost), and among those
// it chooses a match that is as long as possible.
// This so-called leftmost-longest matching is the same semantics
// that early regular expression implementations used and that POSIX
// specifies.
//
// However, there can be multiple leftmost-longest matches, with different
// submatch choices, and here this package diverges from POSIX.
// Among the possible leftmost-longest matches, this package chooses
// the one that a backtracking search would have found first, while POSIX
// specifies that the match be chosen to maximize the length of the first
// subexpression, then the second, and so on from left to right.
// The POSIX rule is computationally prohibitive and not even well-defined.
// See http://swtch.com/~rsc/regexp/regexp2.html#posix for details.
func CompilePOSIX(expr string) (*NRegexp, error) {
	re, err := regexp.CompilePOSIX(expr)
	return &NRegexp{Regexp: re}, err
}

// MustCompile is like Compile but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func MustCompile(str string) *NRegexp {
	return &NRegexp{Regexp: regexp.MustCompile(str)}
}

// MustCompilePOSIX is like CompilePOSIX but panics if the expression cannot be parsed.
// It simplifies safe initialization of global variables holding compiled regular
// expressions.
func MustCompilePOSIX(str string) *NRegexp {
	return &NRegexp{Regexp: regexp.MustCompilePOSIX(str)}
}

// NamedCapture converts all matches to named matches.
func (re *NRegexp) NamedCapture(b []byte) map[string][]byte {
	return NamedCapture(re.InnerRegexp(), re.FindSubmatch(b))
}

// NamedCaptureString converts all matches to named matches.
func (re *NRegexp) NamedCaptureString(s string) map[string]string {
	return NamedCaptureString(re.InnerRegexp(), re.FindStringSubmatch(s))
}

// InnerRegexp returns the underlying regexp.
func (re *NRegexp) InnerRegexp() *regexp.Regexp {
	return re.Regexp
}
