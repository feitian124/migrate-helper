package migrate

import (
	"regexp"
	"strings"
)

type FunctionResult struct {
	match    bool
	function string
	caller   string
	name     string
	rest     string
}

func Function(line string) (*FunctionResult, error) {
	p := `(?P<function>\s*func\s*)(?P<caller>\(.*\)\s*)(?P<name>Test.*)(?P<rest>\(.*\)\s*\{\s*)`
	r := regexp.MustCompile(p)
	match := r.FindStringSubmatch(line)
	groupNames := r.SubexpNames()
	//fmt.Printf("%v, %v, %d, %d\n", match, groupNames, len(match), len(groupNames))
	result := &FunctionResult{}
	if len(match) == len(groupNames) {
		// 转换为map
		for i, name := range groupNames {
			if i != 0 && name != "" { // 第一个分组为空（也就是整个匹配）
				if name == "function" {
					result.function = strings.TrimSpace(match[i])
				}
				if name == "caller" {
					result.caller = strings.TrimSpace(match[i])
				}
				if name == "name" {
					result.name = strings.TrimSpace(match[i])
				}
				if name == "rest" {
					result.rest = strings.TrimSpace(match[i])
				}
				result.match = true
			}
		}
	}

	return result, nil
}
