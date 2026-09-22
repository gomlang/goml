package native

import (
	"path"
	"regexp"
)

func PathRegexp(pattern, name string) (bool, bool) {
	matched, err := regexp.MatchString("\\A(?:"+pattern+")\\z", name)
	return matched, err == nil
}

func PathMatch(pattern, name string) (bool, bool) {
	matched, err := path.Match(pattern, name)
	return matched, err == nil
}
