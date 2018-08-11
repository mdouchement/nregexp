package nregexp_test

import (
	"reflect"
	"regexp"
	"testing"

	"github.com/mdouchement/nregexp"
)

const pattern = `.*\s.*\s(?P<my_capture>[^\s]*)\s.*`

func TestNamedCapture(t *testing.T) {
	re := regexp.MustCompile(pattern)
	matches := re.FindSubmatch([]byte("Example of nregexp package"))

	actual := nregexp.NamedCapture(re, matches)
	expected := map[string][]byte{"my_capture": []byte("nregexp")}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestNamedCaptureString(t *testing.T) {
	re := regexp.MustCompile(pattern)
	matches := re.FindStringSubmatch("Example of nregexp package")

	actual := nregexp.NamedCaptureString(re, matches)
	expected := map[string]string{"my_capture": "nregexp"}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestNRegexpNamedCapture(t *testing.T) {
	re := nregexp.MustCompile(pattern)

	actual := re.NamedCaptureString("Example of nregexp package")
	expected := map[string]string{"my_capture": "nregexp"}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestNRegexpNamedCaptureString(t *testing.T) {
	re := nregexp.MustCompile(pattern)

	actual := re.NamedCapture([]byte("Example of nregexp package"))
	expected := map[string][]byte{"my_capture": []byte("nregexp")}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Expected %v but got %v", expected, actual)
	}
}

func TestNRegexpMatchString(t *testing.T) {
	// Check if regexp.Regexp methods are exported
	re := nregexp.MustCompile(pattern)
	if !re.MatchString("Example of nregexp package") {
		t.Fail()
	}
}
