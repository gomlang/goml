package native

import (
	"encoding/hex"
	"net/netip"
)

func ParseAddressBytes(input string) (bool, string, string) {
	ip, err := netip.ParseAddr(input)
	if err != nil {
		return false, "", ""
	}
	return true, hex.EncodeToString(ip.AsSlice()), ip.Zone()
}

func ParsePrefixBytes(input string) (bool, string, int) {
	prefix, err := netip.ParsePrefix(input)
	if err != nil {
		return false, "", 0
	}
	return true, hex.EncodeToString(prefix.Addr().AsSlice()), prefix.Bits()
}
