package native

import "net/netip"

func AddressFacts(input string) (uint32, string) {
	ip := netip.MustParseAddr(input)
	flags := []bool{ip.Is4(), ip.Is6(), ip.Is4In6(), ip.IsUnspecified(), ip.IsLoopback(), ip.IsPrivate(), ip.IsMulticast(), ip.IsLinkLocalUnicast(), ip.IsLinkLocalMulticast(), ip.IsInterfaceLocalMulticast(), ip.IsGlobalUnicast()}
	var result uint32
	for i, flag := range flags {
		if flag {
			result |= 1 << i
		}
	}
	return result, ip.Unmap().String()
}
