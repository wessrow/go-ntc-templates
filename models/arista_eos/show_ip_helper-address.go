package arista_eos

type ShowIpHelperAddress struct {
	Interface string   `json:"INTERFACE"`
	Ip_helper []string `json:"IP_HELPER"`
}

var ShowIpHelperAddress_Template string = `Value Required INTERFACE (\S+)
Value List IP_HELPER (\d+\.\d+\.\d+\.\d+|\S+)

Start
 ^DHCP
 ^Interface -> Continue.Record
 ^Interface:\s+${INTERFACE}$$
 ^\s+DHCP\s+Smart
 ^\s+DHCP\s+servers:\s+${IP_HELPER}$$
 ^\s+${IP_HELPER}$$
 ^$$
 ^. -> Error
`
