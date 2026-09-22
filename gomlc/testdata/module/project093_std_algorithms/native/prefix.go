package native

import "net/netip"

func PrefixFacts(prefix, address, other string) (string, bool, bool) {
	p := netip.MustParsePrefix(prefix)
	return p.Masked().String(), p.Contains(netip.MustParseAddr(address)), p.Overlaps(netip.MustParsePrefix(other))
}

func ValidPrefix(input string) bool {
	_, err := netip.ParsePrefix(input)
	return err == nil
}
