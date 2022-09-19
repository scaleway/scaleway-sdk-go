package strcase

import (
	"testing"
)

func TestAllStrCases(t *testing.T) {
	tests := []struct {
		name, publicGoName, privateGoName, bashArgName string
	}{
		{"foo_bar", "FooBar", "fooBar", "foo-bar"},
		{"foo_bar_baz", "FooBarBaz", "fooBarBaz", "foo-bar-baz"},
		{"Foo_bar", "FooBar", "fooBar", "foo-bar"},
		{"foo_WiFi", "FooWiFi", "fooWiFi", "foo-wi-fi"},
		{"id", "ID", "id", "id"},
		{"Id", "ID", "id", "id"},
		{"foo_id", "FooID", "fooID", "foo-id"},
		{"fooId", "FooID", "fooID", "foo-id"},
		{"fooUid", "FooUID", "fooUID", "foo-uid"},
		{"idFoo", "IDFoo", "idFoo", "id-foo"},
		{"uidFoo", "UIDFoo", "uidFoo", "uid-foo"},
		{"midIdDle", "MidIDDle", "midIDDle", "mid-id-dle"},
		{"APIProxy", "APIProxy", "apiProxy", "api-proxy"},
		{"ApiProxy", "APIProxy", "apiProxy", "api-proxy"},
		{"apiProxy", "APIProxy", "apiProxy", "api-proxy"},
		{"_Leading", "_Leading", "_Leading", "-leading"},
		{"___Leading", "_Leading", "_Leading", "-leading"},
		{"trailing_", "Trailing", "trailing", "trailing"},
		{"trailing___", "Trailing", "trailing", "trailing"},
		{"a_b", "AB", "aB", "ab"},
		{"a__b", "AB", "aB", "ab"},
		{"a___b", "AB", "aB", "ab"},
		{"Rpc1150", "RPC1150", "rpc1150", "rpc1150"},
		{"case3_1", "Case3_1", "case3_1", "case3-1"},
		{"case3__1", "Case3_1", "case3_1", "case3-1"},
		{"IEEE802_16bit", "IEEE802_16bit", "iEEE802_16bit", "ieee802-16bit"},
		{"IEEE802_16Bit", "IEEE802_16Bit", "iEEE802_16Bit", "ieee802-16-bit"},
		{"IPv4", "IPv4", "ipv4", "ipv4"},
		{"Ipv4", "IPv4", "ipv4", "ipv4"},
		{"iPV4", "IPV4", "iPV4", "ipv4"},
		{"RepeatedIpv4", "RepeatedIPv4", "repeatedIPv4", "repeated-ipv4"},
		{"eSport", "ESport", "eSport", "e-sport"},
		{"stopped in place", "StoppedInPlace", "stoppedInPlace", "stopped-in-place"},
		{"l_ssd", "LSSD", "lSSD", "lssd"},
		{"ids", "IDs", "ids", "ids"},
		{"my_resource_ids", "MyResourceIDs", "myResourceIDs", "my-resource-ids"},
		{"acids", "Acids", "acids", "acids"},
		{"secret-key", "SecretKey", "secretKey", "secret-key"},
		{"ip-id", "IPID", "ipID", "ip-id"},
		{"lb-id", "LBID", "lbID", "lb-id"},
		{"acl-id", "ACLID", "aclID", "acl-id"},
		{"dhcp-id", "DHCPID", "dhcpID", "dhcp-id"},
	}
	for _, test := range tests {
		got := ToPublicGoName(test.name)
		if got != test.publicGoName {
			t.Errorf("ToPublicGoName(%q) == %q, want %q", test.name, got, test.publicGoName)
		}
		got = ToPrivateGoName(test.name)
		if got != test.privateGoName {
			t.Errorf("ToPrivateGoName(%q) == %q, want %q", test.name, got, test.privateGoName)
		}
		got = ToBashArg(test.name)
		if got != test.bashArgName {
			t.Errorf("ToBashArg(%q) == %q, want %q", test.name, got, test.bashArgName)
		}
	}
}
