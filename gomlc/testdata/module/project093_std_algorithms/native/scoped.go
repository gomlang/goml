package native

import "net/netip"

func ScopedAddressFacts(input string) (bool, string, string, string) {
	ip, err := netip.ParseAddr(input)
	if err != nil {
		return false, "", "", ""
	}
	return true, ip.Zone(), ip.WithZone("").String(), ip.Unmap().String()
}
