package exp

import "net"

// Conatins IP
type expContainsIp struct {
	key, str string
	cidrnet *net.IPNet
}

func (e expContainsIp) Eval(p Params) bool {
	testIp := net.ParseIP(p.Get(e.key))
	return e.cidrnet.Contains(testIp)
}


func (e expContainsIp) String() string {
	return sprintf("[%s∋%s]", e.key, e.str)
}

// Contains is an expression that evaluates to true if substr falls within the cidr range
// given example:
//
// 192.168.1.0/24 will match all IPs that fall between
// 192.168.1.1 and 	192.168.1.254
//
// 192.168.1.0/32 will only match 192.168.1.0
func ContainsIp(key, substr string) Exp {
	_, cidrnet, err := net.ParseCIDR(substr)
	if err != nil {
		return nil
	}
	return expContainsIp{key, substr, cidrnet}
}
