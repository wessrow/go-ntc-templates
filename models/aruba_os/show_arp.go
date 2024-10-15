package aruba_os

type ShowArp struct {
	Protocol    string `json:"PROTOCOL"`
	Ip_address  string `json:"IP_ADDRESS"`
	Mac_address string `json:"MAC_ADDRESS"`
	Interface   string `json:"INTERFACE"`
	Age         string `json:"AGE"`
}

var ShowArp_Template string = `Value PROTOCOL (\w+)
Value IP_ADDRESS (\d+.\d+.\d+.\d+)
Value MAC_ADDRESS (\S+)
Value INTERFACE (\S+)
Value AGE (\d+)

Start
 ^Codes:\s+\*\s+-\s+Local\s+Addresses,\s+S\s+-\s+Static,\s+A\s+-\s+Auth
 ^Total\s+ARP\s+entries:\s+\d+
 ^IPV4\s+ARP\s+Table                                                                                
 ^-+                                                                                                
 ^\s+Protocol\s+IP\s+Address\s+Hardware\s+Address\s+Interface\s+Age                                 
 ^\s+${PROTOCOL}\s+${IP_ADDRESS}\s+${MAC_ADDRESS}\s+${INTERFACE}\s+${AGE} -> Record                            
 ^\s*$$                                                                                             
 ^. -> Error
`
