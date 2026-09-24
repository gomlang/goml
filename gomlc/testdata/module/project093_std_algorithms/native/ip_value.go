package native

import (
	"encoding/hex"
	"net/netip"
)

func AddressValues(input string) (string, string, string, string) {
	ip := netip.MustParseAddr(input)
	next, prev := "", ""
	if value := ip.Next(); value.IsValid() {
		next = value.String()
	}
	if value := ip.Prev(); value.IsValid() {
		prev = value.String()
	}
	bytes := ip.As16()
	return next, prev, hex.EncodeToString(bytes[:]), hex.EncodeToString(ip.AsSlice())
}

func CompareAddresses(left, right string) int {
	return netip.MustParseAddr(left).Compare(netip.MustParseAddr(right))
}
