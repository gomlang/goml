package native

import (
	"net/netip"
	"net/url"
)

func URLIPv6(input string) bool {
	value, err := netip.ParseAddr(input)
	return err == nil && value.Is6() && value.Zone() == ""
}

func URLAuthority(input string) (string, string, string, bool) {
	value, err := url.Parse("http://" + input + "/")
	if err != nil {
		return "", "", "", false
	}
	return value.Hostname(), value.Port(), value.Redacted(), true
}

func URLEscape(input []byte, query bool) string {
	if query {
		return url.QueryEscape(string(input))
	}
	return url.PathEscape(string(input))
}

func URLUnescape(input string, query bool) ([]byte, bool) {
	var value string
	var err error
	if query {
		value, err = url.QueryUnescape(input)
	} else {
		value, err = url.PathUnescape(input)
	}
	return []byte(value), err == nil
}

func QueryCanonical(input string) (string, bool) {
	values, err := url.ParseQuery(input)
	if err != nil {
		return "", false
	}
	return values.Encode(), true
}

func ResolveURL(base, reference string) (string, bool) {
	b, err := url.Parse(base)
	if err != nil {
		return "", false
	}
	r, err := url.Parse(reference)
	if err != nil {
		return "", false
	}
	return b.ResolveReference(r).String(), true
}
