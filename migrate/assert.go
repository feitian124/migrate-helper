package migrate

import (
	"regexp"
	"strings"
)

type AssertResult struct {
	match   bool
	caller  string
	actual  string
	checker string
	expect  string
}

func Equals(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*Equals),(?P<expect>.+)\)`
	return Assert(line, p)
}

func DeepEquals(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*DeepEquals),(?P<expect>.+)\)`
	return Assert(line, p)
}

func IsNil(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*IsNil)\)`
	return Assert(line, p)
}

func NotNil(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*NotNil)\)`
	return Assert(line, p)
}

func IsTrue(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*IsTrue)\)`
	return Assert(line, p)
}

func IsFalse(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*IsFalse)\)`
	return Assert(line, p)
}

func Greater(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*Greater),(?P<expect>.+)\)`
	return Assert(line, p)
}

func LessEqual(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*LessEqual),(?P<expect>.+)\)`
	return Assert(line, p)
}

func GreaterEqual(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*GreaterEqual),(?P<expect>.+)\)`
	return Assert(line, p)
}

func HasLen(line string) (*AssertResult, error) {
	p := `(?P<caller>\w+)\.Assert\((?P<actual>.+),(?P<checker>\s*HasLen),(?P<expect>.+)\)`
	return Assert(line, p)
}

func Assert(line string, p string) (*AssertResult, error) {
	r := regexp.MustCompile(p)
	match := r.FindStringSubmatch(line)
	groupNames := r.SubexpNames()
	//fmt.Printf("%v, %v, %d, %d\n", match, groupNames, len(match), len(groupNames))
	result := &AssertResult{}
	if len(match) == len(groupNames) {
		// ?????????map
		for i, name := range groupNames {
			if i != 0 && name != "" { // ????????????????????????????????????????????????
				if name == "caller" {
					result.caller = strings.TrimSpace(match[i])
				}
				if name == "actual" {
					result.actual = strings.TrimSpace(match[i])
				}
				if name == "checker" {
					result.checker = strings.TrimSpace(match[i])
				}
				if name == "expect" {
					result.expect = strings.TrimSpace(match[i])
				}
				result.match = true
			}
		}
	}

	return result, nil
}
